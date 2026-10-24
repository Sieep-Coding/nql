package main

import (
	"fmt"
	"nql"
)

func main() {
	e, err := nql.NewEngine()
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
