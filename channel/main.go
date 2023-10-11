package main

import (
	"fmt"
	"time"
)

func main() {
	var ch chan struct{}
	ch = make(chan struct{})
	go func() {
		time.Sleep(time.Second)
		<-ch
	}()
	fmt.Println("print before")
	ch <- struct{}{}
	fmt.Println("print after")
}
