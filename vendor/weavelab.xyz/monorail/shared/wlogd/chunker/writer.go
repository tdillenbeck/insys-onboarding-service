package chunker

import (
	"encoding/binary"
	"io"
)

type ChunkWriter struct {
	w       io.Writer
	sizeBuf []byte
}

func NewChunkWriter(w io.Writer) ChunkWriter {
	return ChunkWriter{w, make([]byte, 4)}
}

// WriteChunk writes a complete []byte chunk with a little-endian uint32 length prefix
func (w *ChunkWriter) WriteChunk(buf []byte) error {
	var err error

	binary.LittleEndian.PutUint32(w.sizeBuf, uint32(len(buf)))

	// Write must return a non-nil error if it returns n < len(p).

	_, err = w.w.Write(w.sizeBuf)
	if err != nil {
		return err
	}

	_, err = w.w.Write(buf)
	if err != nil {
		return err
	}

	return nil
}
