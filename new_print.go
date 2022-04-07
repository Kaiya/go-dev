package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("new fmt.Println, should have message above this line")
	go testNewProc()
}

func testNewProc() {
	time.Sleep(3)
}
