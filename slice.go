package main

import (
	"fmt"
)

func main() {
	slice := []int{10, 20, 30, 70, 80}
	fmt.Println(slice)
	slice_2 := slice[0:3]
	fmt.Println(slice_2)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	fmt.Println(len(slice_2))
	fmt.Println(cap(slice_2))

	slice_1 := make([]int, 5, 8)
	fmt.Println(slice_1)

	arr := []int{3, 5, 6, 7, 8}
	slice_3 := arr[1:4]
	fmt.Println(arr)
	fmt.Println(slice_3)
	fmt.Println(cap(slice_3))
	fmt.Println(cap(arr))
	slice_3[2] = 9
	fmt.Println(arr)
	fmt.Println(slice_3)
	//appending
	slice_3 = append(slice_3, 10, 60)
	fmt.Println(slice_3)

	new_slice := append(slice, slice_2...)
	fmt.Println(new_slice)

	// deleting
	i := 2
	slice[i] = slice[len(slice)-1]
	slice = slice[:len(slice)-1]
	fmt.Println(slice)

	src_slice := []int{1, 2, 3, 4}
	dest_slice := []int{5, 6, 7}
	num := copy(src_slice, dest_slice)
	fmt.Println(dest_slice)
	fmt.Println("number of elements copied", num)
	//looping through a slice (Range)
	for index, value := range src_slice {
		fmt.Println(index, "=>", value)
	}
	for _, value := range src_slice {
		fmt.Println(value)

	}

	for index, _ := range src_slice {
		fmt.Print(index)

	}
	fmt.Println()

	slice_4 := []string{"Hi", "welcome", "to", "Go world"}
	fmt.Println(slice_4)

	slice_4 = append(slice_4, "yamuna")
	fmt.Println(slice_4)

}
