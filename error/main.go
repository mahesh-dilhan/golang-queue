package main

type MyError struct {
}

func (er *MyError) Error() []rune {
	merror := []rune{100, 100, 100}
	return merror
}

func main() {
	e := MyError{}
}
