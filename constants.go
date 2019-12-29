package main

import(
	"fmt"
)
const (
	first = iota + 6
	second = 2 << iota
)
func main()  {
	const pi = 3.142

	fmt.Println(pi)

	const c = 3

	fmt.Println(c + 3)

	fmt.Println(c + 1.2)
	fmt.Println("iota values ")
	fmt.Println(first, second)


}