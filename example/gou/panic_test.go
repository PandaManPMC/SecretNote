package main

import (
	"SecretNote/tutil"
	"fmt"
	"testing"
	"time"
)

func catch() {
	e := recover()
	fmt.Println(e)
	err, isOk := e.(error)
	if isOk {
		tutil.GetInstanceByReflectUtil().ErrToMap(err)
	}
}

func Test1(t *testing.T) {
	go func() {
		time.Sleep(1 * time.Second)
		//panic("test panic2")
	}()
	go func() {
		defer catch()
		panic("test panic")
	}()
	time.Sleep(10 * time.Second)
}
