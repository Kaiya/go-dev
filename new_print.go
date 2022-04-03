package main

import (
	"fmt"
	"time"
)

func kaiya() {
	fmt.Println("new fmt.Println, should have message above this line")
	go testNewProc()
}

func testNewProc() {
	time.Sleep(3)
}
