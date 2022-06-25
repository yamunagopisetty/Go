package main

import (
	"fmt"
)

func main() {

	var b *int
	var s *string

	c := "hello"
	var d *string = &c
	fmt.Println(d)

	e := "yamuna"
	var f = &e
	fmt.Println(f)

	fmt.Println(b)
	fmt.Println(s)

	var a int = 4565
	i := 10

	var p = &a

	fmt.Println("value stored in p=", p)
	println("address of a=", &a)

	fmt.Println("the address of i is", &i)
	fmt.Printf("%T %v\n", *(&i), *(&i))

}
