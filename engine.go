package nql

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"
)

type Engine struct {
	data map[string]int64
	file *os.File
	mu   sync.Mutex
}

type Item struct {
	Key    string
	Value  string
	Offset int64
}

var keyValueSeparator = " "

const Seconds = 5

func NewEngine() (*Engine, error) {
	file, err := os.OpenFile("data.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	return &Engine{
		data: make(map[string]int64),
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

	_, err := e.file.Seek(e.data[key]+int64(len(key))+1, 0)
	if err != nil {
		return "", fmt.Errorf("error seeking in file: %w", err)
	}

	buffer := make([]byte, 1)

	var content []byte

	for {
		n, err := e.file.Read(buffer)

		if err != nil {
			fmt.Println("Error reading from file:", err)
			break
		}

		if n == 0 {
			fmt.Println("End of file reached")
			break
		}

		if buffer[0] == '\n' {
			break
		}
		content = append(content, buffer[0])
	}
	return string(content), nil
}

func (e *Engine) Set(key, value string) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	offset, err := e.file.Seek(0, io.SeekEnd)
	if err != nil {
		fmt.Println("Error seeking to end of file:", err)
		return err
	}

	_, err = e.file.WriteString(key + keyValueSeparator + value + "\n")
	if err != nil {
		fmt.Println("Error appending text to file:", err)
		return err
	}
	e.data[key] = offset
	return nil
}

func (e *Engine) CompactFile() error {
	for {
		time.Sleep(time.Duration(Seconds) * time.Second)
		fmt.Println("Compacting file...")
		e.mu.Lock()
		_, err := e.GetMapFromFile()
		e.mu.Unlock()
		if err != nil {
			fmt.Println("Error getting map from file:", err)
		}
	}
}

func (e *Engine) GetMapFromFile() ([]Item, map[string]string) {
	m := make(map[string]string)
	i := []Item{}

	_, err := e.file.Seek(0, 0)
	if err != nil {
		fmt.Println("Error seeking to start of file:", err)
		return i, m
	}
	var totalBytesRead int64
	scanner := bufio.NewScanner(e.file)
	for scanner.Scan() {
		line := scanner.Text()
		offset := totalBytesRead
		parts := strings.Split(line, keyValueSeparator)
		if len(parts) >= 2 {
			m[parts[0]] = parts[1]
			i = append(i, Item{Key: parts[0], Value: parts[1], Offset: offset})
		}
	}
	return i, m
}
