package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a interface{}

	var b int
	var c = 30
	a = 10
	b = a.(int)
	fmt.Println("Addition of two numbers are b+c=", b+c)
	fmt.Printf("a is % T\n", a)

	//--------------------
	//Constants
	const i int = 10
	const j = 1000
	const k = (1000 - 100) / i
	fmt.Println("constant value", k)
	if k%2 == 0 {
		fmt.Println("k is even number")
	} else {
		fmt.Println("k is odd number")
	}
	fmt.Printf("values of %v %v %v ", i, j, k)
	//------------------------
	//complex numbers
	{
		cmplx1 := complex(1000.343, 321.12312)
		fmt.Println("Complex number ", cmplx1, "type of complex number", reflect.TypeOf(cmplx1))

		var r1, i1 float32 = 1000.343, 321.12312
		//var i2 float64 = 321.12312
		cmplx2 := complex(r1, i1)

		fmt.Println("Complex number ", cmplx2, "type of complex number", reflect.TypeOf(cmplx2))
	}

}
