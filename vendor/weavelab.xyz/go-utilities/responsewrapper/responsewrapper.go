package responsewrapper

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"weavelab.xyz/wlib/werror"
	"weavelab.xyz/wlib/wlog"
)

type ResponseWrapper struct {
	start time.Time

	ResponseTime int64                    `json:"response_time"`
	Timers       map[string]time.Duration `json:"timers,omitempty"`
	Data         interface{}              `json:"data,omitempty"`
	ErrorMsg     string                   `json:"error,omitempty"`

	InsertID        string      `json:"insert_id,omitempty"`
	RecordsInserted int64       `json:"records_inserted,omitempty"`
	RecordsUpdated  int64       `json:"records_updated,omitempty"`
	BadRecords      []BadRecord `json:"bad_records,omitempty"`

	writer  io.Writer
	request *http.Request

	status int
	sent   bool
}

func New(writer io.Writer, request *http.Request) *ResponseWrapper {

	w := ResponseWrapper{start: time.Now()}

	w.writer = writer
	w.request = request

	return &w
}

func (w *ResponseWrapper) Error(err error) {

	w.ErrorMsg = err.Error()

	msg, _ := json.Marshal(w.ErrorMsg)

	w.sent = true
	if hw, ok := w.writer.(http.ResponseWriter); ok {

		status := http.StatusBadRequest

		if w.status != 0 {
			status = w.status
		}

		http.Error(hw, `{"error": `+string(msg)+`}`, status)
	}
}

func (w *ResponseWrapper) StatusMessage(status int, message string) {
	w.StatusError(status, message, nil)
}

func (w *ResponseWrapper) StatusError(status int, message string, err error) {
	w.Status(status)

	msgErr := werror.New(message)

	if status >= 500 {
		if err == nil {
			err = msgErr
		}

		wlog.WError(werror.Wrap(err, message))
	}

	w.Error(msgErr)
}

func (w *ResponseWrapper) Send() {
	if w.sent {
		return
	}

	if hw, ok := w.writer.(http.ResponseWriter); ok {
		hw.Header().Add("Content-Type", "application/json")
		if w.status != 0 {
			hw.WriteHeader(w.status)
		}
	}

	w.ResponseTime = int64(int64(time.Since(w.start)/time.Microsecond)) / 1000

	enc := json.NewEncoder(w.writer)

	err := enc.Encode(w)

	if err != nil {
		w.Error(err)
		log.Println("Unable to Marshal to JSON: ", err)
		return
	}
}

func (w *ResponseWrapper) Status(status int) {
	w.status = status
}
