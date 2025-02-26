package main

import "fmt"

func main() {
	fmt.Print("Enter your name: ")

	var name string

	_, err := fmt.Scan(&name)
	if err != nil {
		return
	}

	fmt.Println("Hello, " + name)

}
