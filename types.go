package main

import(
	"fmt"
)

func main()  {

	var i int
	i = 42
	fmt.Println(i)

	var f float64 = 32.14
	fmt.Println(f);

	firstname := "ogaga"
	fmt.Println(firstname)

	b := true
	fmt.Println(b)

	c := complex(3, 4)
	fmt.Println(c)

	r, im := real(c), imag(c)
	fmt.Println(r, im)
}
