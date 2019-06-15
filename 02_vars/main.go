package main

import "fmt"

func main() {
	// var name = "Daga"
	// var age int32 = 23

	name, age := "Daga", 23

	fmt.Println(name, age)
	fmt.Printf("%T  %T\n", name, age)
}
