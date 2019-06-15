package main

import (
	"fmt"
	"strconv"
)

func arrays() {
	fmt.Println("Executing arrays")
	// Arrays
	// var fruitArr [2]string
	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Orange"

	// Declare and assign at the same time
	fruitArr := [2]string{"Apple", "Orange"}
	fruitSlice := []string{"Apple", "Orange", "Grapes"}

	fruitSlice = append(fruitSlice, "Mango")

	fmt.Println(fruitArr, fruitSlice)
	fmt.Println("Slicing in Arrays", fruitSlice[1:3])

}

func conditionals() {
	// If-else is the same so not writing it here
	// Writing a switch case
	color := "red"

	switch color {
	case "red":
		fmt.Println("color is red")
	case "blue":
		fmt.Println("color is blue")
	default:
		fmt.Println("color is NOT red or blue")
	}
}

func loops() {
	// Long method for loop -- Like While loop in other languages
	i := 1
	for i < 10 {
		fmt.Print(i, " ")
		i++
	}
	fmt.Println()

	// Normal for loop
	for i := 0; i < 10; i++ {
		fmt.Print(i, " ")
	}

	fmt.Println()

	// For range loop
	aString := "hello"
	for i, v := range aString {
		fmt.Print(i, string(v), " ")
	}

	// FizzBuzz
	for i := 1; i < 100; i++ {
		if i%21 == 0 {
			fmt.Println("fizzbuzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%7 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}

}

func maps() {
	// Maps are key-value pairs in go
	// emails := make(map[string]string)

	// emails["Bob"] = "bobdabadob@personal.com"
	// emails["Sharon"] = "sharoncool@bob.com"

	// Declaration of a map directly in go
	emails := map[string]string{"Bob": "Bobadoob@gmail.com"}

	fmt.Println(emails["Bob"])

	// Delete from maps
	delete(emails, "Bob")
	// fmt.Println(emails)

	// This does not throw an error
	userName := "Bob"
	userEmail, userExists := emails[userName]
	if userExists {
		fmt.Println(userEmail)
	} else {
		fmt.Printf("The given user %s does not exist", userName)
	}

}

func rangeLoops() {
	ids := []int{10, 20, 30, 40, 50}

	// Loop trough ids using range
	// If we don't want to use the ndex then put _ in its place
	for index, value := range ids {
		fmt.Println(index, " ", value)
	}

	nbaBestPlayers := map[int]string{1: "Steph", 2: "Klay", 3: "Kevin", 4: "Dray", 5: "Kevon"}

	for pos, player := range nbaBestPlayers {
		fmt.Printf("%d - %s\n", pos, player)
	}
}

func pointers() {
	// same as C, for the basic level
}

func adderClosure() func(int) int {
	sum := 0
	// This is a closure, as soon as the var to adderClosure fn is made,
	// the sum var is still available to his func, this ensures that we can still
	// access it inside the closure and adderClosure returns a func that returns an int, so
	// return type is int
	return func(val int) int {
		sum += val
		return sum
	}
}

func stringToPrint(str string) func(int) []string {

	return func(no int) []string {
		var arrayOfStrings []string

		for i := 0; i < no; i++ {
			arrayOfStrings = append(arrayOfStrings, str)
		}
		return arrayOfStrings
	}
}

// structs are an important part of go, serving as the class alternative
func structs() {
	type Person struct {
		name string
		age  int
	}

	person1 := Person{name: "Piyush Daga", age: 23}
	person2 := Person{"Ola Boi", 22}

	fmt.Println(person1)
	fmt.Println(person2)

}

// Person is struct that contains the first and last name
type Person struct {
	firstName string
	lastName  string
	age       int
}

// Greeting mthod (value receivers)
func (p Person) greeting() {
	fmt.Println("Hello my name is " + p.firstName + " " + p.lastName +
		" and my age is " + strconv.Itoa(p.age))
}

// hasBirthday method pointer receiver
func (p *Person) hasBirthday() {
	p.age++
}

func main() {
	// arrays()
	// conditionals()
	// loops()
	// maps()
	// rangeLoops()
	// pointers()

	// sum := adderClosure()
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(sum(i))
	// }
	// fmt.Println(stringToPrint("Johhny Boi")(10))
	// structs()
	per1 := Person{"Nathon", "Fallon", 26}
	per1.greeting()
	per1.hasBirthday()
	per1.greeting()

}
