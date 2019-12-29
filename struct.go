package main

import(
	"fmt"
)

type User struct {
	ID int
	FirstName string
	LastName string
}

func main()  {

	var u User
	u.FirstName = "Ogaga"
	u.ID = 123
	u.LastName = "Agi"
	fmt.Println(u)

	u2 := User{
		ID:        0,
		FirstName: "onos",
		LastName:  "agi",
	}

	fmt.Println(u2)
}