package main

import (
	"fmt"
	"math"
	"math/rand"
)

func foo() {
	fmt.Println("The square root of 4 is", math.Sqrt(4))
}

func main() {
	foo()
	fmt.Println("A nuber between 1-100", rand.Intn(100))
}
