package main

import "fmt"

type dog struct {
	first string
}

func (d dog) walk() {
	fmt.Println("My name is", d.first, "and I'm walking.")
}

func (d *dog) run() {
	d.first = "Rover"
	fmt.Println("My name is", d.first, "and I'm running.")
}
func (d dog) run2() {
	d.first = "Rover"
	fmt.Println("My name is", d.first, "and I'm running.")
}

type youngin interface {
	walk()
	run()
}

func youngRun(y youngin) {
	y.run()
}

type kid struct {
	first string
}

func john(k kid) kid {
	k.first = "john"
	return k
}
func johnP(k *kid) {
	k.first = "johnP"
}

func main() {
	a := 42
	fmt.Println(&a)
	d1 := dog{"Henry"}
	fmt.Println(&d1)
	fmt.Printf("%p\t%T\n", &d1, &d1)
	d1.walk()
	fmt.Println(d1)
	d1.run2()
	fmt.Println(d1)
	d1.run()
	fmt.Println(d1)
	// youngRun(d1) // doesn't run

	d2 := &dog{"Padget"}
	fmt.Println(d2)
	d2.walk()
	fmt.Println(d2)
	d2.run()
	fmt.Println(d2)
	youngRun(d2)
	fmt.Println(d2)
	// pointers vs values changing 'root' variable
	fmt.Println("\nKids")
	k1 := kid{"tim"}
	fmt.Println(k1.first)
	k1 = john(k1)
	fmt.Println(k1.first)
	johnP(&k1)
	fmt.Println(k1.first)

}

/*

The idea of the method set is integral to how interfaces are implemented and used in Go.

The method set of a type T consists of all methods with receiver type T.
These methods can be called using variables of type T.

The method set of a type *T consists of all methods with receiver *T or T
These methods can be called using variables of type *T.
it can call methods of the corresponding non-pointer type as well

*/
