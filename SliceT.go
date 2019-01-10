package main

import (
	"fmt"
	"math/rand"
)

func main() {
	i := rand.Int()
	println(i)
	vv := make([]byte, i)
	fmt.Printf("%b", vv)
}
