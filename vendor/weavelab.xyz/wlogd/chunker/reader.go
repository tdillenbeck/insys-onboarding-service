package chunker

import (
	"encoding/binary"
	"io"
)

type ChunkReader struct {
	r            io.Reader
	buf          []byte
	afterMessage []byte
}

func NewChunkReader(r io.Reader) ChunkReader {
	return ChunkReader{r, make([]byte, 4096), nil}
}

// ReadChunk returns a complete []byte chunk. It is a slice not a copy, so
// if you need a copy you must make it yourself.
func (r *ChunkReader) ReadChunk() ([]byte, error) {
	copy(r.buf, r.afterMessage)
	readStart := len(r.afterMessage)

	for readStart < 4 {
		length, err := r.r.Read(r.buf[readStart:])
		readStart += length

		if err != nil {
			return r.buf[:readStart], err
		}
	}

	chunkSize := int(binary.LittleEndian.Uint32(r.buf[:4]))

	if len(r.buf) < 4+chunkSize {
		newLen := len(r.buf) * 2
		if 4+chunkSize > newLen {
			newLen = 4 + chunkSize
		}

		// buffer is full but message is still longer, so double the buffer
		nbytes := make([]byte, newLen)
		copy(nbytes, r.buf)
		r.buf = nbytes
	}

	for readStart < int(4+chunkSize) {
		length, err := r.r.Read(r.buf[readStart:])
		readStart += length

		if err != nil {
			return r.buf[:readStart], err
		}
	}

	chunk := r.buf[4 : 4+chunkSize]
	// if there is anything left, save it and start there in the next ReadChunk
	r.afterMessage = r.buf[4+chunkSize : readStart]

	return chunk, nil
}
