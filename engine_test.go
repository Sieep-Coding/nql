package nql

import "testing"

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

func Test_GetMapFromFile(t *testing.T) {
	e, err := NewEngine()
	if err != nil {
		t.Error("Engine setup failed")
	}
	e.mu.Lock()
	_, _ = e.GetMapFromFile()
	e.mu.Unlock()
	if err != nil {
		t.Error("Failed to get map from file")
	}
}
