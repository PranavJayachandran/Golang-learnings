package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)
	done := make(chan string)
	go func(channel chan string) {
		for {
			select {
			case s := <-channel:
				fmt.Println(s)
			case <-done:
				return
			}
		}
	}(channel)

	go func(channel chan string) {
		channel <- ("done1")
		done <- ("complete")
	}(channel)
	go func(channel chan string) {
		channel <- ("done2")
	}(channel)
	go func(channel chan string) {
		channel <- ("done3")
	}(channel)

	time.Sleep(time.Second * 3)
}
