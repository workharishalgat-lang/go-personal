package main

import "fmt"

func main() {
	d := "string"
	dr := []rune(d) // strings are convertibe to slice of runes
	fmt.Println(dr)
	d = string(dr) // runes are convertibe to string

	funcVar := a
	var pp p = funcVar // this is automatcally convetibble as funcvar is of non named type
	pp()

	// but lets define a varibale  of type c and check if they are assignale to each other

	var cc c

	// below line will produce compiler error as both are named types
	// pp = cc
	// to make it assignable we can use type conversion
	pp = p(cc) // this is an convertible as cc shares the same structure
}
func a() int {
	return 0
}

type p func() int

type c func() int
