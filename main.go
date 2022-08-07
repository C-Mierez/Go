package main

import "fmt"

func main() {
	variables()
}

func variables() {
	var a int
	var b bool = true
	var c = "inferredStr"
	d := "inferredAndDeclaredStr"
	a = 10

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
