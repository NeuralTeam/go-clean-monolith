package rotator

import (
	"fmt"
	"log"
	"os"
	"sync"
)

func NewFileWriter() FileWriter {
	size, err := ParseStringSize("256MB")
	if err != nil {
		log.Fatalf("rotator: cannot parsing 'defaultMaxSize': %s", err)
	}

	return FileWriter{
		Filename:   "",
		Prefix:     "",
		MaxSize:    size,
		MaxAge:     60 * 60 * 24 * 30,
		MaxBackups: 32,
		IsCompress: false,
		size:       0,
		file:       nil,
		mu:         sync.Mutex{},
	}
}

type FileWriter struct {
	// Filename without extension
	Filename  string
	Extension string
	Prefix    string

	// MaxSize in B
	MaxSize uint64

	// MaxAge in seconds
	MaxAge     int
	MaxBackups uint
	IsCompress bool

	size uint64
	file *os.File
	mu   sync.Mutex
}

func (fw *FileWriter) Write(p []byte) (n int, err error) {
	fw.mu.Lock()
	defer fw.mu.Unlock()

	writeLen := uint64(len(p))
	if writeLen > fw.MaxSize {
		return 0, fmt.Errorf(
			"write 'length=%d' exceeds maximum 'file size=%d'", writeLen, fw.MaxSize,
		)
	}

	if fw.file == nil {
		if err = fw.openExistingOrNew(len(p)); err != nil {
			return 0, err
		}
	}

	if fw.size+writeLen > fw.MaxSize {
		if err = fw.rotate(); err != nil {
			return 0, err
		}
	}

	n, err = fw.file.Write(p)
	fw.size += uint64(n)

	return n, err
}

func (fw *FileWriter) openNew() error {
	name := fw.Filename
	mode := os.FileMode(0600)

	// we use truncate here because this should only get called when we've moved
	// the file ourselves. if someone else creates the file in the meantime,
	// just wipe out the contents.
	f, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, mode)
	if err != nil {
		return fmt.Errorf("can't open new logfile: %s", err)
	}
	fw.file = f
	fw.size = 0
	return nil
}

func (fw *FileWriter) openExistingOrNew(writeLen int) error {
	filename := fw.Filename
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return fw.openNew()
	}
	if err != nil {
		return fmt.Errorf("error getting log file info: %s", err)
	}

	if uint64(info.Size())+uint64(writeLen) >= fw.MaxSize {
		return fw.rotate()
	}

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		// if we fail to open the old log file for some reason, just ignore
		// it and open a new log file.
		return fw.openNew()
	}
	fw.file = file
	fw.size = uint64(info.Size())
	return nil
}

func (fw *FileWriter) compress() error {
	return nil
}

func (fw *FileWriter) rotate() error {
	return nil
}
