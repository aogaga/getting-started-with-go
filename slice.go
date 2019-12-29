package main

import(
	"fmt"
)

func main()  {
	arr := [3]int{1, 2, 3}

	slice := arr[:]

	arr[1] = 42
	slice[2] = 27

	fmt.Println(arr, slice)


	slice2 := []int{1, 2, 3}

	fmt.Println(slice2)

	slice2 = append(slice2, 4, 5, 6, 7, 8)

	fmt.Println(slice2)

	s2 := slice2[3:]

	fmt.Println(s2)
}
