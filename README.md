# Go<sub>lang</sub> 🐿

> First time ever learning Go.
> Writing down all my finding about the language in here for future reference.

- [Go<sub>lang</sub> 🐿](#gosublangsub-)
  - [Packages](#packages)
  - [Modules](#modules)
    - [Naming Convention](#naming-convention)
  - [Variables must be used](#variables-must-be-used)
  - [Factoring statements](#factoring-statements)
  - [Primitive Types](#primitive-types)
  - [Other Types](#other-types)
    - [Arrays and Slices](#arrays-and-slices)
  - [Maps](#maps)
  - [Structs](#structs)
  - [Pointers](#pointers)
  - [Operators](#operators)
  - [Constants and Enums](#constants-and-enums)
  - [Control Flow](#control-flow)
    - [Defer](#defer)
    - [Panic](#panic)
    - [Recover](#recover)
  - [Functions!](#functions)
    - [Interfaces](#interfaces)
    - [Go Routines!](#go-routines)
    - [Channels](#channels)

## Packages

Go programs start running in the `main` package. It is a special package that is used with programs that are meant to be executable.

By convention, Executable programs (the ones with the main package) are called Commands. Others are called simply Packages.

Packages only really make sense in the context of a separate program which uses them.

## Modules

A module is a collection of Go packages stored in a directory with a `go.mod` file at its root. The go.mod file defines the module's path, which is also the import path used while importing packages that are part of this module.


### Naming Convention

Variables, Types and Functions declared beginning with uppercase are globally visible and can be accessed outside of the package.


Anything beginning with lowercase are scoped only to the local package, and is not exported.

Variables declared inside functions are scoped to that block only.

```golang 
var A int // Global 
var a int // Locally global

func f() {
    var b int // Block scoped
}
```

## Variables must be used

Programs will **not** compile if there are variables that are not used. This is an opinionated design decision from the language itself.

## Factoring statements

It is possible to factor in different kinds of statements in Go by using grouping them in `()`.

```golang
import (
    "fmt"
    "time"
)

// or

var (
    x int
    y bool
    z int
)
```

## Primitive Types

Integers can be declared of multiple sizes (`int8, int16...int64`, and unsigned equivalents). Bigger numbers should be handled with other libraries.  

- `int` / `uint`
- `bool`
- `byte` 
- `float32` / `float64`
- `complex64` /  `complex128` (Can use `real()` and `imag()` to access components)
## Other Types

- `string` (UTF-8 Byte arrays)
- `rune` (int32)

### Arrays and Slices

Arrays
- Fixed Size `[2]int{...elements...}`
- Dynamic `[...]int{elements}`
> Initializing with elements is optional

- Access length with `len(arr)`

- Default to be a value type. Copies are real copies.

Slice
- Declared with `[]int{...elements...}`

- Access length with `len(slice)` (Current length)
- Access capacity with `cap(slice)` (Current capacity before a new copy and memory reallocation is needed in order to store more values)

- Default to be a reference type. Copies take the address.
  
- Can be sliced with the `:` operator to specify an interval. 
  - ex. `[:], [3:], [:9], [3:9]`

- Can be created with the `make([]type, length, capacity?)` function.
  - ex. `make([]int,3), make([]int, 10, 50)` 

- Elements can be spread with the operator `...`
  - ex. `append(slice, elementA, elementB)` can be done using an already existing slice by doing `append(slice, []int{elementA, elementB}...)` 

## Maps

- Can use the `delete()` function to remove elements. 
  - ex. `delete(myMap, "John")`

## Structs

- Can be declared anonymously `foo := struct{name string}{name: "John"}`

**Embedding** is a tool similar to OOP inheritance. Inheritance is not a thing in Go.

Also, it is possible to use **Tags** on fields to add restrictions to it. These are just strings. Any logic and actual restrictions need to be parsed and figured out by some library that knows how to interpret these.
```golang
type Animal struct {
  name string
  age int `required max:"100"` // <-- Tags!
}

type Bird struct {
  Animal // <-- Embedding!
  canFly bool 
}
```

## Pointers

Yes, there are pointers.
- `&` Address
- `*` Pointer

It is possible to use the `new(myStruct)` to create a new empty object and receive a pointer to it.

Dereferencing can be done in two ways:
- `(*pointer).field` to access `field` from the dereferenced pointer.
-  `pointer.field` is syntactic sugar. Compiler understand we are not accessing `field` from the pointer but from the underlying object it points to.

Slices and Maps are some data structures that when "copied", instead a pointer to the underlying data is given.
## Operators

- Boolean Byte operators: `&` And, `|` Or, `^` XOR, `&^` AndNot
- Byte Operators: `<<` Left Shift, `>>` Right Shift

## Constants and Enums

Constants can't be computed at compile time. Must be explicit values.

Blocks can be used in conjunction with `iota` type to create enumerated constants.

```golang
const (
    _ = iota    // Ignore val 0
    a           // val 1
    b           // val 2
    c           // val 3
)

// Can modify the pattern

const (
    _ = iota                // Ignore first value
    KB = 1 << (10 * iota)   // 1 ^ 10
    MB                      // 1 ^ 100
    GB                      // 1 ^ 1000
    TB                      // 1 ^ 10000
)
```

## Control Flow

### Defer
The keyword `defer` executes any function passed into it AFTER every statement has been run by the function, but BEFORE it returns. Deferred functions get executed in LIFO order.

Additionally, parameters in deferred functions are eagerly assigned; i.e. values used will be those declared before the deferred functions statement, ignoring what happens afterwards even-though the function is executing at the end.

### Panic
Go does not have exceptions. Many "exceptional" cases are instead considered normal. To avoid common connotational meaning from exceptions, Go instead refers to these situations as **Panic**.

In order to throw a Panic it is possible to use `panic(errMessage)`

The pattern is that things are normally not opinionated about whether something is a panic or not. Instead, if something goes wrong, it is returned as an error, and it is then up to the developer to decide whether to panic or not.

> **Panics** happen AFTER deferred statements. And deferred statements always execute even if the function panics.

### Recover

`recover()` is a built-in function that regains control of a panicking goroutine. Recover is only useful inside deferred functions.

This is especially useful on nested functions. `recover()` recovers from a panic, in which the function stops its execution, but functions higher up the call stack can still continue running, since the panicked function has "handled" the error.

However, if `recover()` is used, and the error can still not be handled, another `panic()` would need to be thrown in order to make sure the application stops running (because it is in an irrecoverable state).

## Functions!
They can return pointers to locally declared variables. Go automatically upgrades these variables from the stack to the heap.

They can be treated as variables. Thus creating the possibility of having high-order functions that take other functions as parameters.

They can return multiple values. Normally this is used to return `(expectedType, error)` tuples, since errors are preferred instead of panicking.

Deep scoped functions should use variables from outside their scope by receiving them as parameters instead of accessing them directly. This avoids future conflicts when async execution is introduced.
- ```golang
  func main() {
    x := 10
    for 1 < 3 {
      shouldBreak := func myDeepFunc(x int) bool {
        return x == 10
      }(x) // <-- Pass as parameter
      if shouldBreak {
        break
      }
    }
  }
  ```
### Methods? 👀

These are basically functions that are executing in a known context (Any type in Go). So we are adding *methods* to a certain type in the context of its value.

```golang
type myStruct struct {
	name string
	age  int
}

func (ms myStruct) toString() string {
	return fmt.Sprintf("Name: %v, Age: %v \n", ms.name, ms.age)
}

func main() {
  someStruct := myStruct{
    name: "John",
    age: 50,
  }

  someStruct.toString()
}
```

Methods can receive the address of the type instead, and handle values from reference instead of creating copies.

Methods can only be defined on types you have access to. That is, no primitives.

## Interfaces

We can create interfaces as a type, in which we declare functions without implementation.

Later on, **any type** could potentially implement this function as a **method** and be used as an instance of this interface.

The way Go is thought out is that, interfaces are not something you need to worry about at design time, but instead consumers can create an interface themselves and shape it into the way they need in their application.  

Interfaces can also be embedded in a similar way as Structs. 

```golang
type IntA interface {}

type IntB interface {}

type IntAB interface {
  IntA
  IntB
}

```

There is also a special kind of interface named the "empty interface" which is just as the name implies, an interface with no methods defined, and it is declared on the fly.

To do anything useful with it, it is likely gonna require reflection steps in which to try and figure out what methods are available on it.

```golang
var myObj interface{} = MyActualInstance{}

if ai, ok := myObj.(AnotherInterface); ok {
  ai.Write("Hi")
}

yai, ok := myObj.(YetAnotherInterface) // Type Conversion
if ok {
  // ...
}
```

The best practice is to export the explicit objects and don't define interfaces unless strictly necessary. Go is backwards compared to other languages, since interfaces are implicit.

- Your components don't explicitly *implement* interfaces 
  
  This avoids forcing consumers to have to implement every single thing and instead allow them to focus on just the things they need.

- but do try to *use* other interfaces, hence allowing a more free future compose-ability.

  Export interfaces of the types that you are consuming. Then whoever is using the package knows what methods need to be implemented in order for your component to work - but this saves you from having to worry about the implementation of those consumed objects.

- also try to receive interfaces whenever possible.

## Go Routines!

We use the `go` keyword to launch a function as a routine. These are "green" routines that are managed and scheduled by the Go scheduler and make smart and efficient use of actual system threads in order to avoid overhead from the creation/destruction of these. 

It is common practice to use anonymous functions with go routines, though it is important to make sure coupling is handled properly in order to avoid race condition between routines. 

Some already existing libraries allow for sync management between routines, like the `RWMutex` and `WaitGroup`.

### Channels 

They are cool.
