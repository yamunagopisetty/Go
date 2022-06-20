package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Print("Enter the three numbers to find largest =")
	fmt.Scanln(&a, &b, &c)
	switch {
	case a > b && a > c:
		fmt.Println(a, "is Greater than ", b, "and", c)
	case b > a && b > c:
		fmt.Println(b, "is Greater than ", a, "and", c)
	case c > a && c > b:
		fmt.Println(c, "is Greater than ", a, "and", b)
	default:
		fmt.Println("all are equal")
	}
}
