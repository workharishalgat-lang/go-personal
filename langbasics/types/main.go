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

}
