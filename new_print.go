package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("new fmt.Println, should have message above this line")
	go testNewProc()
	file, err := os.Create("newfile.csv")
	if err != nil {
		fmt.Println("create file error: ", err.Error())
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	writer.Write([]string{"hello", "world"})
	writer.Flush()

}

func testNewProc() {
	time.Sleep(3)
}
