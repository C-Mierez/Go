# Go<sub>lang</sub> 🐿

> First time ever learning Go.
> Writing down all my finding about the language in here for future reference.

- [Go<sub>lang</sub> 🐿](#gosublangsub-)
  - [Packages](#packages)
    - [Naming Convention](#naming-convention)
  - [Variables must be used](#variables-must-be-used)
  - [Factoring statements](#factoring-statements)

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
- `array`: `[]{type}()`


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