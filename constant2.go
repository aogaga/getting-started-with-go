package main

import(
	"fmt"
)

const(
	first1 = iota
	second1
	third
)

func main()  {
	fmt.Println(first1, second1, third)
}