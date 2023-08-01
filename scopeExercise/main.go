package main

import "fmt"

var i = 42

const (
	wife = "syd"
	son  = "baby fern"
)

func main() {
	myName := "fern"
	fmt.Printf("My name is %v\nMy wifes name is %v\nmy son's name is %v\nI have %v cousins\n", myName, wife, son, i)
}
