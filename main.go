package main

import (
	"fmt"

	"github.com/C-Mierez/go/packer"
	// Could also declare a package alias as:
	// alias "github.com/C-Mierez/go/packer"
)

func main() {

	// Testing out variable declarations
	variableDeclaration()

	// Testing my own package
	ownPackage()

	// Declaring and using constants
	usingConstants()

	// Declaring arrays
	usingArrays()
}

func usingArrays() {

	fixedArr := [3]int{1, 2, 3}
	dynamicArr := [...]int{9, 8, 7}

	fmt.Printf("Fixed size array: %v, %T \n", fixedArr, fixedArr)
	fmt.Printf("Dynamic size array: %v, %T \n", dynamicArr, dynamicArr)

}

const (
	constA = iota
	constB
	constC
)

const (
	_   = iota             // Ignore first value
	KiB = 1 << (10 * iota) // 2 ^ 10
	MiB                    // 2 ^ 100
	GiB                    // 2 ^ 1000
	TiB                    // 2 ^ 10000
)

// Encoding boolean "roles" in one single byte of data
const (
	isAdmin = 1 << iota
	isActive
	isHeadquarters
	canSeeFinance
	canSeeHR
)

func usingConstants() {
	// Watch out for casing when declaring constants
	const (
		ExportedPI = 3.14
		localPI    = 3.14
		// No computed constants
		// computed = math.Sin(1.57)
	)

	fmt.Printf("ExportedPI: %v, localPI: %v \n", ExportedPI, localPI)

	// Using constants declared as iota
	// Differently scoped constants don't share the same iota value
	fmt.Printf("constA: %v, constB: %v, constC: %v \n", constA, constB, constC)

	// Using enumerated constants with a different pattern
	fmt.Printf("KiB: %v \n", KiB)
	fmt.Printf("MiB: %v \n", MiB)
	fmt.Printf("TiB: %v \n", TiB)

	var roles byte = isAdmin | isActive | canSeeHR
	fmt.Printf("Encoded roles as a byte: %b \n", roles)
	fmt.Printf("isAdmin: %v \n", isAdmin&roles == isAdmin)

}

func ownPackage() {

	// Using a function from another package
	i := 5
	fmt.Printf("Is %v prime? %v \n", i, packer.IsPrime(i))
}

func variableDeclaration() {
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
