package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Print("Enter the three numbers to find largest =")
	fmt.Scanln(&a, &b, &c)
	if a > b && a > c {
		fmt.Println(a, "is Greater than ", b, "and", c)
	} else if b > a && b > c {
		fmt.Println(b, "is Greater than ", a, "and", c)
	} else if c > a && c > b {
		fmt.Println(c, "is Greater than ", a, "and", b)
	} else {
		fmt.Println("all are equal")
	}

}
