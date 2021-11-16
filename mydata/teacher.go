package mydata

import "fmt"

type Teacher struct {
	Class int
	Age   int
	Name  string
}

func (t *Teacher) SayHello() {
	//panic("not implemented") // TODO: Implement
	fmt.Printf("%s SayHello!", t.Name)
}

func (t *Teacher) ShowAge() {
	//panic("not implemented") // TODO: Implement
	fmt.Printf("%s ShowAge:%d", t.Name, t.Age)

}

func (t *Teacher) ShowClass() {
	fmt.Printf("%s ShowClass:%d", t.Name, t.Class)

}
