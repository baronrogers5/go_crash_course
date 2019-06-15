package main

import (
	"fmt"
	"math"

	"github.com/baronrogers5/go_crash_course/03_packages/strutil"
)

//anotherReverse
func anotherReverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func main() {
	fmt.Println("Hello World !")
	fmt.Println(math.Floor(100.71))
	fmt.Println(strutil.Reverse("okay"))

	fmt.Println(anotherReverse("ola") == strutil.Reverse("ola"))
}
