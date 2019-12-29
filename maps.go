package main

import(
	"fmt"
)

func main()  {

	mymaps := map[string] int{ "foo": 42}
	fmt.Println(mymaps)

	delete(mymaps, "foo")

	fmt.Println(mymaps)
}
