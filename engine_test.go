package nql

import "testing"

func Test_SetGetValue(t *testing.T) {
	e := NewEngine()
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
