package main

import (
	"math/rand"
)

func main() {
	//Even or odd number
	y := 20
	if y%2 == 0 {
		println(y, "is an even number")
	} else {
		println(y, "is an odd number")
	}

	if num1, num2 := rand.Intn(1000), rand.Intn(2000); (num1+num2)%2 == 0 {
		println((num1 + num2), "is an even number")
	} else {
		println((num1 + num2), "is an odd number")
	}
	var age uint8 = 18

	if age >= 18 {
		println("Major")
	} else {
		println("Minor")
	}

	// ---------------
	var gender string = "female"

	if age >= 18 && gender == "female" {
		println("She is eligible for marriage according to the law of land")
	} else if age >= 21 && gender == "male" {
		println("He is eligible for marriage according to the law of land")
	} else if gender != "male" && gender != "female" {
		println("no idea")
	} else {
		println("not eligible for marriage")
	}
	//pattern program0
	for a := 0; a < 5; a++ {
		for b := 0; b <= a; b++ {
			print("*")
		}
		println()
	}

	//fibonacci series
	n := 10
	t1 := 0
	t2 := 1

	print("First ", n, " terms: ")
	for i := 0; i < 10; i++ {
		print(t1, " , ")
		sum := t1 + t2
		t1 = t2
		t2 = sum
	}

}
