package nql

import (
	"errors"
	"fmt"
)

func (e *Engine) Get(key string) (string, error) {
	value, ok := e.data[key]
	if !ok {
		return "", errors.New("key not found")
	}
	return value, nil
}

func main() {
	e, err := NewEngine()
	if err != nil {
		panic(err)
	}
	e.Set("foo", "bar")
	value, err := e.Get("foo")
	if err != nil {
		panic(err)
	}
	fmt.Println(value)
}
