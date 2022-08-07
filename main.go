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

	// Declaring arrays and slices
	usingArrays()
	usingSlices()

	// Declaring maps
	usingMaps()

	// Declaring structs
	usingStructs()

	// Looooooooooops
	usingLoops()
}

func usingLoops() {

	// For loop
	for i := 0; i < 3; i++ {
		fmt.Printf("For loop: %v \n", i)
	}
	// While loop
	i := 0
	for i < 3 {
		fmt.Printf("While loop: %v \n", i)
		i++
	}
	// Infinite loop
	i = 0
	for {
		fmt.Printf("Infinite loop: %v \n", i)
		i++
		if i == 3 {
			break
		}
	}

	// Multiple variables in loop
	for i, j := 0, 0; i < 3; i, j = i+1, j+1 {
		fmt.Printf("Multiple variables in loop: i:%v, j:%v \n", i, j)
	}

	// Range loop
	for k, v := range []int{1, 2, 3} {
		fmt.Printf("Range loop from Slice: Key:%v, Value:%v \n", k, v)
	}

	// Range loop from map
	for k, v := range map[string]int{"Charles": 1, "John": 2, "Mary": 7} {
		fmt.Printf("Range loop from Map: Key:%v, Value:%v \n", k, v)
	}

	// Range loop from channel
	// TODO

}

type LivingThing struct {
	isAlive bool
}

type Doctor struct {
	LivingThing // Embedded struct
	id          int
	age         int
	name        string
	patientIds  []int
}

func usingStructs() {

	// These can also be declared without specifying the names of the fields
	// and using positional order instead
	aDoctor := Doctor{
		LivingThing: LivingThing{true},
		id:          1,
		age:         60,
		name:        "Charles",
		patientIds:  []int{10, 11, 12},
	}

	fmt.Printf("Doctor: %v, %T \n", aDoctor, aDoctor)
	fmt.Printf("Doctor's id: %v, %T \n", aDoctor.id, aDoctor.id)
	fmt.Printf("Doctor's first patient id: %v, %T \n", aDoctor.patientIds[0], aDoctor.patientIds[0])
	fmt.Printf("Doctor's embedded struct: %v, %T \n", aDoctor.LivingThing, aDoctor.LivingThing)
}

func usingMaps() {

	userCities := map[string]int{
		"Charles": 1,
		"John":    2,
		"Mary":    7,
		"Jane":    3,
		"Tom":     15,
	}

	fmt.Printf("User cities: %v, %T \n", userCities, userCities)
	fmt.Printf("Charles' city: %v, %T \n", userCities["Charles"], userCities["Charles"])

	// Can also declare an empty map using the make() function

	delete(userCities, "Mary")
	fmt.Printf("User cities after deleting Mary: %v, %T \n", userCities, userCities)

	// Using ok syntax to check for errors
	_, ok := userCities["Mary"]
	fmt.Printf("Mary's city: %v, OK: %v \n", userCities["Mary"], ok)
	_, ok2 := userCities["Charles"]
	fmt.Printf("Charles's city: %v, OK: %v \n", userCities["Charles"], ok2)
}

func usingSlices() {

	sliceA := []int{1, 2, 3}
	sliceB := []int{4, 5, 6}

	fmt.Printf("Slice A: %v, %T \n", sliceA, sliceA)
	fmt.Printf("Slice B: %v, %T \n", sliceB, sliceB)

	// Appending two slices together using the ... operator
	sliceC := append(sliceA, sliceB...)

	fmt.Printf("Slice Appended: %v, %T \n", sliceC, sliceC)
	fmt.Printf("Slice A: %v, %T \n", sliceA, sliceA)
	fmt.Printf("Slice B: %v, %T \n", sliceB, sliceB)

	// Slicing with intervals can affect the original slice
	sliceD := append(sliceA[:1], sliceA[2:]...)

	fmt.Printf("Slice Appended intervals: %v, %T \n", sliceD, sliceD)
	fmt.Printf("Slice A: %v, %T \n", sliceA, sliceA)

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
