package main

import "fmt"

func main() {
	var dd error = P{}
	switch dd.(type) {
	case P:
		fmt.Println("P Type")
	}
}

type P struct {
}

func (p P) Error() string {
	return "no error"
}
