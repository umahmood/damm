# Damm Algorithm
 
Damm is a small Go library which implements the [Damm check digit algorithm](https://en.wikipedia.org/wiki/Damm_algorithm).

# Installation

> go get github.com/umahmood/damm

> cd $GOPATH/src/github.com/umahmood/damm

> go test ./...

# Usage

    package main

    import (
        "fmt"

        "github.com/umahmood/damm"
    )

    func main() {
        c := damm.Calc(572) // c = 5724

        fmt.Println(damm.Validate(c)) // output: true
    }

# Documentation

> http://godoc.org/github.com/umahmood/damm

# License

See the [LICENSE](LICENSE.md) file for license rights and limitations (MIT).

