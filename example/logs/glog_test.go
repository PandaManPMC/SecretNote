package logs

import (
	"strconv"
	"testing"
)

func Test1(t *testing.T) {
	s := "21.00"
	f, _ := strconv.ParseFloat(s, 10)
	i := int64(f)
	t.Log(i)

	var a uint64 = 0
	t.Log(a - 1)
}
