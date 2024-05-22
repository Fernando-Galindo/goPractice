package main

import "fmt"

type Span interface {
    SetName(name string)
    GetName() string
}

type MySpan struct {
    name string
}

func (ms *MySpan) SetName(name string) {
    ms.name = name
}

func (ms MySpan) GetName() string {
    return ms.name
}

func errHelper(span Span) {
    span.SetName("newName")
}

func main() {
	second := MySpan{name: "originalName"}
	fmt.Println("Before:", second.GetName()) // Output: originalName
	second.SetName("not original name")
	fmt.Println("After:", second.GetName()) // Output: ???
	
	mySpan := MySpan{name: "originalName"}
	var span Span = &mySpan
	
	fmt.Println("Before:", span.GetName()) // Output: originalName
	errHelper(span)
	fmt.Println("After:", span.GetName()) // Output: newName
}
