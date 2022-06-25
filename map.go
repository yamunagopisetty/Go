package main

import "fmt"

func main() {
	codes := map[int]string{1: "yamuna", 2: "sai", 3: "teja"}

	println(codes)
	println(codes[1])
	fmt.Println(len(codes))

	value, found := codes[1]
	fmt.Println(value, found)

}
