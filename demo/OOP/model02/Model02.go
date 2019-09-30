package main

type Person struct {
	Name string
}

type emp struct {
	*Person
	Age int
}

type lea struct {
	Person
	Age int
}

func main()  {
	e := emp{Age: 10}
	e.Person=&Person{Name:"emp"}

	lea:=lea{Age:11}
	lea.Person=Person{Name:"lea"}

}

