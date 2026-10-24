package nql

import (
	"bufio"
	"fmt"
	"os"
	"testing"
	"time"
)

func Test_SetGetValue(t *testing.T) {
	e, err := NewEngine()
	if err != nil {
		t.Error("Engine setup failed")
	}
	e.Set("foo", "bar")
	value, err := e.Get("foo")
	if err != nil {
		t.Error(err)
	}
	if value != "bar" {
		t.Error("Value not 'Bar'")
	}
	_, err = e.Get("notfound")
	if err == nil {
		t.Error("should return error")
	}
}

func Test_SetGetKeyValue(t *testing.T) {
	e, err := NewEngine()
	if err != nil {
		t.Error("Engine setup failed")
	}
	e.Set("foo", "bar")
	value, err := e.Get("foo")
	if err != nil {
		t.Error(err)
	}
	if value != "bar" {
		t.Error("Value not 'Bar'")
	}
	_, err = e.Get("notfound")
	if err == nil {
		t.Error("should return error")
	}
}

// GetFileContent is a helper function.
func (e *Engine) GetFileContent(f *os.File) []string {
	e.mu.Lock()
	defer e.mu.Unlock()

	_, err := f.Seek(0, 0)
	if err != nil {
		fmt.Println(err)
		return []string{}
	}

	scanner := bufio.NewScanner(f)

	var content []string
	for scanner.Scan() {
		line := scanner.Text()
		content = append(content, line)
	}
	return content
}

func Test_GetFileContent(t *testing.T) {
	e, err := NewEngine()
	if err != nil {
		t.Error("Engine setup failed")
	}
	e.Set("foo", "bar")
	content := e.GetFileContent(e.file)
	if len(content) == 0 {
		t.Error("File content is empty")
	}
	if content[0] != "foo bar" {
		t.Error("File content is not 'foo bar'")
	}
}

func TestEngine_Compact(t *testing.T) {
	os.Remove("data.txt")
	v1 := "latestvalue1"
	v2 := "latestvalue2"
	e, _ := NewEngine()
	e.Set("key1", "value1")
	e.Set("key2", "value2")
	e.Set("key1", v1)
	e.Set("key2", v2)
	e.Set("key3", "value3")
	go e.CompactFile()
	time.Sleep((Seconds + 3) * time.Second)
	if len(e.GetFileContent(e.file)) != 3 {
		t.Errorf("Expected %d, but got %d", 3, len(e.GetFileContent(e.file)))
	}
}
