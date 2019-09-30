package main

import "fmt"

type Person interface {
	eat()
	SayHello()
}

type emp struct {
	Name string
}

func (e emp) eat()  {
	fmt.Println(e.Name+" eat")
}

func (e emp) SayHello()  {
	fmt.Println(e.Name+" say hello")
}

func main()  {
	e1 := emp{Name: "小小"}
	var p1 Person=e1
	p1.eat()
	p1.SayHello()

}