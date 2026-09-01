package main

import "fmt"

func main() {
	d := "harish"
	dd := []rune(d)
	revString(dd, 0, len(dd)-1)
	d = string(dd)
	fmt.Println(d)
}

func revString(str []rune, left, right int) {
	if left >= right {
		return
	}
	str[left], str[right] = str[right], str[left]
	revString(str, left+1, right-1)
}
