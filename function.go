package main

import(
	"fmt"
)

func main()  {
	fmt.Println("Hello, Playground")
	port := 3000
	port2, err := startWebServer(port, 2)
	fmt.Println(port2, err)
}

func startWebServer(port int, numberOfRetries int) (int, error)  {

	fmt.Println("Starting Web Server....")
	fmt.Println("Server Started")
	return port, nil;
}