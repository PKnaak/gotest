package main

import (
	"fmt"
	"math"
)

func test (a, b string) (string, string) {
 return b, a
}

func main () {
    fmt.Println("Hello Nia!")    
    fmt.Println(test("Nia!","Sup"))
    fmt.Println(math.Pi)
}