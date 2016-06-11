/*
Package damm implements the Damm check digit algorithm.

Example:

    package main

    import (
        "fmt"

        "github.com/umahmood/damm"
    )

    func main() {
        c := damm.Calc(572)

        fmt.Println(damm.Validate(c)) // output: true
    }
*/
package damm
