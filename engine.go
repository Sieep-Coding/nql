package nql

import (
	"fmt"
	"io"
	"os"
	"sync"
)

type Engine struct {
	data map[string]string
	file *os.File
	mu   sync.Mutex
}

var keyValueSeparator = " "

func NewEngine() (*Engine, error) {
	file, err := os.OpenFile("data.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	return &Engine{
		data: make(map[string]string),
		file: file,
		mu:   sync.Mutex{},
	}, nil
}

func (e *Engine) Get(key string) (string, error) {
	e.mu.Lock()
	defer e.mu.Unlock()

	if _, ok := e.data[key]; !ok {
		return "", fmt.Errorf("key not found")
	}

	_, err := e.file.Seek(e.data[key]+int64(len(key))+1,0)
	if err != nil {
		return "", fmt.Errorf("error seeking in file: %w", err)
	}
	return e.data[key], nil
}

func (e *Engine) Set(key, value string) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	_, err := e.file.Seek(0, io.SeekEnd)
	if err != nil {
		fmt.Println("Error seeking to end of file:", err)
		return err
	}

	_, err = e.file.WriteString(key + keyValueSeparator + value + "\n")
	if err != nil {
		fmt.Println("Error appending text to file:", err)
		return err
	}
	e.data[key] = value
	return nil
}
