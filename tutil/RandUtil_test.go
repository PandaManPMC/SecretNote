package tutil

import (
	"strings"
	"testing"
)

func TestRandNumber(t *testing.T) {
	for i := 0; i < 100; i++ {
		t.Log(RandNumber(100))
	}
	s := strings.ToUpper(RandCharacterString(320))
	t.Log(s)
	t.Log(0)
	t.Log("o")
	t.Log("O")
	t.Log("1")
	t.Log("l")
	t.Log("L")

	var sum uint64 = 36
	for i := 0; i < 16; i++ {
		sum *= 36
	}
	t.Log(sum)
}
