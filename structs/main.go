package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type student struct {
	personalInfo person
	studentId    string
	gpa          float64
}

func (s student) honors() {
	if s.gpa >= 3.5 {
		fmt.Printf("%v will graduate with a gpa of %v and will have honors \n", s.personalInfo.first, s.gpa)
	}
}

func (p person) sayName() {
	fmt.Println(p.first)
}
func (s student) sayName() {
	fmt.Println(s.personalInfo.first)
}
func (p person) legalDrinker() bool {
	return p.age >= 21
}

type generalInfo interface {
	sayName()
	honors()
	//legalDrinker()
}

func spillTea(gI generalInfo) {
	gI.sayName()
	gI.honors()
}

func main() {
	s1 := student{
		studentId: "12345",
		gpa:       3.89,
		personalInfo: person{
			first: "tim",
			last:  "jones",
			age:   22,
		},
	}
	s1.honors()
	spillTea(s1)
	fmt.Println("Can drink ", s1.personalInfo.legalDrinker())

}
