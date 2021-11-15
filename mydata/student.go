package mydata

import (
	"fmt"
)

type Student struct {
	Age  int
	Name string
}

func (s *Student) SetAge(age int) {
	//panic("not implemented") // TODO: Implement
	s.Age = age
}

func (s *Student) SetName(name string) {
	//panic("not implemented") // TODO: Implement
	s.Name = name
}

func (s *Student) SayHello() {
	//panic("not implemented") // TODO: Implement
	fmt.Printf("%s say Hello!", s.Name)
}

func (s *Student) ShowAge() {
	//panic("not implemented") // TODO: Implement
	fmt.Printf("%s show age:%d", s.Name, s.Age)
}

func ShowName2() {
	fmt.Println("hello")
	// nums := []float64{4.62, 90.31, 18.4, 70, 498}
	// avg := Average(nums)
}
