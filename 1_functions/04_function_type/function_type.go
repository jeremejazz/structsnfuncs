package main

import "fmt"

func main() {

	parrot := greet
	//from here we have assigned the function itself greet
	// to the variable parrot.
	// another way of writing this without using the shorthand declaration (:=) is by using the signature

	var tommy func(string) = greet

	// execute the variable as if it were a function
	parrot("Matey!")
	tommy("Mark")

}

func greet(message string) {
	fmt.Println("Hello", message)
}
