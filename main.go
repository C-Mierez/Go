package main

import "fmt"

func main() {
	variables()
}

func variables() {
	// Normal way of declaring variables
	var a int
	var b bool = true
	a = 10

	// Inferring the type of a variable from its first value
	var c = "inferredStr"

	// Declaring and inferring the type of a variable at the same time from the first value
	d := "inferredAndDeclaredStr"
	f := 9.

	// Declaring multiple variables at once
	// Usually for variables that are related to each other but independent to everything else
	var (
		name    string = "Charles"
		age     int    = 60
		isAlive bool   = true
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	// Printing out the value and type of a variable
	fmt.Printf("%v, %T \n", d, d)
	fmt.Printf("%v, %T \n", f, f) // Value of `f` declared and initialized as Float

	fmt.Printf("Name: %v, Age: %v, Alive: %v, \n", name, age, isAlive)

}
