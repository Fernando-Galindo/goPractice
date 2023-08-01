package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("Program starts here")
}

func main() {
	// for i := 0; i < 4; i++ {
	// 	x := rand.Intn(250)
	// 	fmt.Printf("x is an %T equal to %v\n", x, x)
	// 	switch {
	// 	case x <= 100:
	// 		fmt.Println("Between 0-100")
	// 	case x <= 200:
	// 		fmt.Println("Between 101-200")
	// 	case x <= 250:
	// 		fmt.Println("Between 201-250")
	// 	}
	// 	if x >= 0 && x <= 100 {
	// 		fmt.Println("Between 0-100")
	// 	} else if x >= 101 && x <= 200 {
	// 		fmt.Println("Between 101-200")
	// 	} else if x >= 201 && x <= 250 {
	// 		fmt.Println("Between 201-250")
	// 	}
	// }
	// x, y := rand.Intn(10), rand.Intn(10)
	// fmt.Printf("x : %v\ny: %v\n", x, y)
	// if x < 4 && y < 4 {
	// 	fmt.Println("Less than 4")
	// } else if x > 6 && y > 6 {
	// 	fmt.Println("greater than 6")
	// } else if x >= 4 && x <= 6 {
	// 	fmt.Println("x>=4 && x<=6")
	// } else if y != 5 {
	// 	fmt.Println("y not 5")
	// } else {
	// 	fmt.Println("none met")
	// }
	// for {
	// 	x := rand.Intn(5)
	// 	fmt.Printf("x = %v\n", x)
	// 	if x != 3 {
	// 		continue
	// 	}
	// 	break
	// }
	// x := []int{45, 34, 23, 12}
	// for i, v := range x {
	// 	fmt.Printf("index %v; value %v; same as %v\n", i, v, x[i])
	// }
	// x := map[string]int{
	// 	"james": 42,
	// 	"matt":  10,
	// }
	// m1 := x["james"]
	// fmt.Println(m1)
	// if m2, ok := x["q"]; ok {
	// 	fmt.Println(m2)
	// } else {
	// 	fmt.Println("q does not exist")
	// }
	for i := 0; i < 100; i++ {
		if x := rand.Intn(5); x == 3 {
			fmt.Println("x equals 3!!!")
		}
	}
}
