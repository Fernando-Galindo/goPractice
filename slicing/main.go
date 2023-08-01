package main

import "fmt"

func main() {
	ogSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	fg := ogSlice[:4]
	sg := ogSlice[4:8]
	tg := ogSlice[8:12]
	fmt.Printf("fg : %#v\nsg : %#v\ntg : %#v\n", fg, sg, tg)
	tg = append(tg, 13, 14)
	fmt.Printf("fg : %#v\nsg : %#v\ntg : %#v\n", fg, sg, tg)
	tg = append(tg[:3], tg[4:]...)
	fmt.Printf("fg : %#v\nsg : %#v\ntg : %#v\n", fg, sg, tg)

	jb := []string{"James", "Bond", "007"}
	es := []string{"ethan", "Stone", "MI6"}
	compSlice := [][]string{jb, es}
	fmt.Printf("comp: %#v\n", compSlice)
	fmt.Printf("comp: %v\n", compSlice)
}
