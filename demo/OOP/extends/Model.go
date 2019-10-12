package main

import "fmt"
/**
 *Go 中通过组合的方式来实现继承的，又称为匿名组合
 */

type Person struct {
	Id   int
	Name string
	Age  int
}

type Student struct {
	Person
	Class string
	SNo   string
}

func (p Person) SayName()  {
	fmt.Println(p.Name)
}

func (p Person) SayAge()  {
	fmt.Println(p.Age)
}

func (s Student) SayName()  {
	fmt.Println(s.Name," ",s.Class)
}

func main()  {
	person := Person{Name: "ss",Age:10}
	student := Student{Class: "class"}
	student.Person=person
	fmt.Println(student.Person.Name)
	student.SayName()
	student.SayAge()
}
