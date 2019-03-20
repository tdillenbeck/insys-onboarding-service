package aggregator

import (
	"io"
	"sync"
	"time"

	"weavelab.xyz/monorail/shared/wlib/werror"
	"weavelab.xyz/monorail/shared/wlib/wlog/tag"
	"weavelab.xyz/monorail/shared/wlib/wmetrics/wmetricslog"
)

type message []byte

//AggregateWriter implements io.WriteCloser and is used to wrap another io.WriteCloser. It aggregates bytes before sending through to the destination writer.
type AggregateWriter struct {
	dst          func() (io.WriteCloser, error)
	separator    []byte
	sendInterval time.Duration
	capacity     int

	msgQueue  chan message
	quit      chan bool
	closed    bool
	closedMtx *sync.RWMutex

	writeMtx *sync.Mutex
}

//NewAggregateWriter returns a new io.WriteCloser that aggregate bytes until either the given duration or capacity, at which point the dst function will be used to create the destination writer and write the bytes.
func NewAggregateWriter(separator []byte, sendInterval time.Duration, capacity int, dst func() (io.WriteCloser, error)) *AggregateWriter {
	w := &AggregateWriter{
		dst:          dst,
		separator:    separator,
		sendInterval: sendInterval,
		capacity:     capacity,
		msgQueue:     make(chan message, 30),
		quit:         make(chan bool),
		closedMtx:    &sync.RWMutex{},
		writeMtx:     &sync.Mutex{},
	}

	go w.sendWorker()
	return w
}

//Write queues bytes to be aggregated and sent eventually.
func (w *AggregateWriter) Write(p []byte) (n int, err error) {
	w.writeMtx.Lock()
	defer w.writeMtx.Unlock()

	w.closedMtx.RLock()
	defer w.closedMtx.RUnlock()

	if w.closed {
		return 0, werror.New("AggregateWriter is closed! Cannot write to closed writer.")
	}

	//Send this message to the queue where it will be aggregated
	select {
	case w.msgQueue <- p:
		return len(p), nil
	case <-time.After(time.Second):
		return 0, werror.New("Timed out while queueing new message")
	}
}

func (w *AggregateWriter) sendWorker() {
	currentState := &pendingMessageState{
		capacity: w.capacity,
	}

	for {
		//Wait for either a quit signal or a new message to be written
		select {
		case <-w.quit:
			//The sendWorker should stop working
			return
		case p := <-w.msgQueue:
			//Lock the current state so we can change it
			currentState.lock.Lock()

			//If we already have some pending messages we either need to send the data or use the separator
			if currentState.pendingCount > 0 {
				//We need to either send them or prepend a separator to this new message
				if len(currentState.pendingMessages)+len(p)+len(w.separator) > w.capacity {
					currentState.send(w.dst)
				} else {
					p = append(w.separator, p...)
				}
			}
			//Queue the newest message
			currentState.queueMsg(p)

			if w.sendInterval == 0 {
				//If the sendInterval is 0 then we send it immediately
				currentState.send(w.dst)
			} else if !currentState.sendStarted {
				//If a send has not already been started then we will start one now
				currentState.sendStarted = true
				time.AfterFunc(w.sendInterval, func() {
					//The state must be locked before we send the aggregated bytes
					currentState.lock.Lock()
					defer currentState.lock.Unlock()

					currentState.send(w.dst)
				})
			}
			currentState.lock.Unlock()
		}
	}
}

//Close attempts to stop the sender
func (w *AggregateWriter) Close() error {
	w.closedMtx.Lock()
	defer w.closedMtx.Unlock()

	select {
	case w.quit <- true:
		w.closed = true

		return nil
	case <-time.After(1 * time.Second):
		return werror.New("Timed out while attempting to close aggregate writer")
	}
}

type pendingMessageState struct {
	capacity        int
	pendingMessages message
	pendingCount    int
	sendStarted     bool
	lock            sync.Mutex
}

func (s *pendingMessageState) send(dstFunc func() (io.WriteCloser, error)) {
	if len(s.pendingMessages) == 0 {
		return
	}

	defer func() {
		s.pendingMessages = make([]byte, 0, s.capacity)
		s.pendingCount = 0
		s.sendStarted = false
	}()

	//Get the destination writer and send bytes to it
	dst, err := dstFunc()
	if err != nil {
		wmetricslog.Logger.WError(werror.Wrap(err, "Error getting destination writer, could not write aggregated message"))
		return
	}
	if dst == nil {
		return
	}
	defer dst.Close()

	_, err = dst.Write(s.pendingMessages)
	if err != nil {
		wmetricslog.Logger.WError(werror.Wrap(err, "Error writing to destination, could not write aggregated message"))
		return
	}
	wmetricslog.Logger.Debug("Wrote aggregated message", tag.String("message", string(s.pendingMessages)))

}

func (s *pendingMessageState) queueMsg(msg []byte) {
	s.pendingMessages = append(s.pendingMessages, msg...)
	s.pendingCount++
}
