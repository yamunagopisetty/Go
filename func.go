package main

func value(i interface{}) {
	println(i)
}

func main() {
	a := "yamuna"
	value(a)

	Greet()
}

func Greet() {
	println("Hello world")
}
