package main

import "crypto/sha256"

var ru [256]byte

func init() {
	for k := range ru {
		ru[k] = ru[k>>1] + byte(k&0x01)
		println(ru[k])
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

func BitCount(x, y [32]byte) int {

}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	BitCount(c1, c2)
}
