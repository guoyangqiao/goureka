package main

import (
	"crypto/sha256"
)

var ru [256]byte

func init() {
	for k := range ru {
		ru[k] = ru[k>>1] + byte(k&0x01)
	}
}

func PopCount(x uint64) int {
	return int(ru[byte(x>>(0*8))] +
		ru[byte(x>>(1*8))] +
		ru[byte(x>>(2*8))] +
		ru[byte(x>>(3*8))] +
		ru[byte(x>>(4*8))] +
		ru[byte(x>>(5*8))] +
		ru[byte(x>>(6*8))] +
		ru[byte(x>>(7*8))])
}

func BitCount1(x, y []byte) int {
	var sum byte
	for i := 0; i < len(x); i++ {
		sum += ru[byte(x[i]^y[i])]
	}
	return int(sum)
}
func main() {
	c1 := sha256.Sum256([]byte("A"))
	c2 := sha256.Sum256([]byte("B"))
	count := BitCount1(c1[:], c2[:])
	println(count)
}
