package main

import (
	"fmt"
	"strings"
)

type DerivedType string

func (d *DerivedType) toUppercase(str string) string {
	return strings.ToUpper(str)
}

func main() {
	d := DerivedType("sg")
	fmt.Println(d.toUppercase("sg"))
}
