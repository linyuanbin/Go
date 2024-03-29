package main

import (
	"errors"
	"fmt"
	"sync"
)
var lock sync.WaitGroup
func sayHello()  {
	fmt.Println("sayHello start")
	panic(errors.New("show panic"))
	fmt.Println("hello")
	lock.Done()
}

func main()  {
	lock.Add(1)
	defer func() {
		fmt.Println("defer-----------")
		if r := recover(); r != nil {
			fmt.Printf("Runtime error caught: %v", r)
		}
	}()
	go sayHello()
	lock.Wait()
	fmt.Println("main")
}

