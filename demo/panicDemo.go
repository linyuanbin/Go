package main

import (
	"errors"
	"fmt"
)

func hello()  {
	fmt.Println("sayHello start")
	panic(errors.New("show panic"))
	fmt.Println("hello")
}

func main()  {
	defer func() { //需要先声明defer，否则捕获不到错误
		fmt.Println("defer-----------")
		if r := recover(); r != nil {
			fmt.Printf("Runtime error caught: %v", r)
		}
	}()
	hello()
	//错误捕获也不会继续执行下面的代码
	fmt.Println("main")
}
