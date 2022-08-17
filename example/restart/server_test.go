package main

import (
	"net/http"
	"testing"
)

func TestGet(t *testing.T) {
	for {
		r, e := http.Get("http://162.14.100.243:7080/test")
		t.Log(e)
		t.Log(r)
	}
}
