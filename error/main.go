package main

import "fmt"

type MyError struct {
}

func (er *MyError) Error() string {

	return "something bad happen"
}

func main() {
	e := &MyError{}
	fmt.Println(e)
}
