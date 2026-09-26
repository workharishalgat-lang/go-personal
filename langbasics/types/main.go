package main

func main() {
	type iff interface{}
	type pp func(int, int) int
	type p struct {
		a   string
		b   map[string]string
		c   chan int
		d   []int
		e   float64
		ifs iff
		fu  pp
	}

	var c CustomFunc = dataFunc
	c.AssociatedMethod()
}

type CustomFunc func(int) int

func (c CustomFunc) AssociatedMethod() {
	c(10)
}

func dataFunc(a int) int {
	return 10
}
