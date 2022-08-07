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

## Packages

Go programs start running in the main package. It is a special package that is used with programs that are meant to be executable.

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