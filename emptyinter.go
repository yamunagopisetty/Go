package main

func main() {
	var title interface{} = "yamuna"
	var date interface{} = 20

	s := title.(string)
	println(s)

	stringTitle, ok := title.(string)
	println(stringTitle, ok)

	//println()

	dateInt, ok := date.(int)
	println(dateInt)
}
