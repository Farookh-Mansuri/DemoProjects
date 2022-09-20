package main

import (
	"fmt"
	"time"
)

func printnumbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}
func printalphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}
func main() {
	go printnumbers()
	go printalphabets()
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("main exit")
}
