package main

import "fmt"

func main() {
	friends := map[string]string{"Syd": "galindo", "Fern": "Galindo"}
	check := "Fern"
	if f, ok := friends[check]; !ok {
		fmt.Printf("%v does not exist\n%#v\n", check, ok)
	} else {
		fmt.Printf("%v does exist\n%v\n%#v\n", check, f, ok)
	}
}
