package main

import "fmt"

func main() {
	// x := 10
	// y := 12
	if age := 12; age >= 5 && age < 10 {
		fmt.Println("adult", age)
	} else if age == 0 {
		fmt.Println("baby")
	} else {
		fmt.Println("child")
	}
	// fmt.Println(string_b)
	// fmt.Println(reflect.TypeOf(string_b))
}
