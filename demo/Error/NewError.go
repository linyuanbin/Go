package main

import (
	"fmt"
)

type myError struct {
	msg string
}

func (err myError) Error() string {
	return err.msg
}

func Less(a int, b int) error {
	if a<b{
		return myError{"a<b"}
	}
	return nil
}

func main() {
	err := Less(1, 2)
	fmt.Println(err.Error())
}
