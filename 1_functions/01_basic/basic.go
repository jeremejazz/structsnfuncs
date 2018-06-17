package main

import "fmt"

func main() {

	greet("Bob")

	s := greetStr("Mark")
	fmt.Println(s)

}

func greet(name string) {
	fmt.Println("Hello " + name)
}

func greetStr(name string) string {

	return "Oh Hai " + name + "!"
}
