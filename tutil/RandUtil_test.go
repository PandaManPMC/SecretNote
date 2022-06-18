package tutil

import "testing"

func TestRandNumber(t *testing.T){
	for i:=0;i<100;i++{
		t.Log(RandNumber(100))
	}
}


