package main

import "fmt"

func main() {
	var a, b, c = 10, 20, 40
	a, b, c = c, a, b

	fmt.Println(a, b, c)
}
