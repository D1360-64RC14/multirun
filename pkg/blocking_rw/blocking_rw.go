package blockingrw

import (
	"io"
	"sync"
)

var _ io.ReadWriter = (*BlockingReadWriter)(nil)

type BlockingReadWriter struct {
	reader io.Reader
	writer io.Writer
	mutex  *sync.RWMutex
}

func New(r io.Reader, w io.Writer) *BlockingReadWriter {
	return &BlockingReadWriter{
		mutex:  new(sync.RWMutex),
		reader: r,
		writer: w,
	}
}

func (b *BlockingReadWriter) Read(p []byte) (int, error) {
	b.mutex.RLock()
	defer b.mutex.RUnlock()

	return b.reader.Read(p)
}

func (b *BlockingReadWriter) Write(p []byte) (int, error) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	return b.writer.Write(p)
}
