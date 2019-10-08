package main

import "fmt"
/**
 *Go中只有组合没有继承
 */

type Person struct {
	Name string
}

func (p Person) getName() string {
	return p.Name
}

type A struct {
	Person
	Name string
}

func (a A) getName() string {
	return a.Name
}

func main()  {
	a := A{ Name: "c"}
	person := Person{Name: "Person"}
	a.Person=person
	fmt.Println(a.getName())
}
