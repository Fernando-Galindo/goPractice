package main

import "fmt"

func addInt(a, b int) int {
	return a + b
}
func addFloat(a, b float64) float64 {
	return a + b
}

// for a generic the possible types (T but could be any letter) are identified in the squar brackets
// then we are saying the parameters are of this generic type T and we are returning this generic type T
func addGeneric[T int | float64](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(addInt(2, 3))
	fmt.Println(addFloat(2.7, 3.2))
	// Generics allow us to create possible types that we will have in a function and consolidate the code
	// Currently to add both floats and ints need a separate function. With generics could have one func that
	// accepts both types and thus can reuse the same code
	fmt.Println(addGeneric(2, 3))
	fmt.Println(addGeneric(2.7, 3.2))
}
