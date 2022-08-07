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