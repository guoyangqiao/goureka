package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"os"
	"reflect"
)

var ru [256]byte

func init() {
	for k := range ru {
		ru[k] = ru[k>>1] + byte(k&0x01)
	}
}

func BitCount(x, y []byte) int {
	var sum byte
	for i := 0; i < len(x); i++ {
		sum += ru[byte(x[i]^y[i])]
	}
	return int(sum)
}
func main() {
	method := os.Args[1]
	println("using ", method)
	x := []byte("A")
	y := []byte("A")
	var me
	switch method {
	case "SHA384":
		me := reflect.ValueOf(sha512.Sum384)
	case "SHA512":
		me := reflect.ValueOf(sha512.Sum512)
	default:
		me := reflect.ValueOf(sha256.Sum256)
	}
	count := BitCount(x[:], y[:])
	println(count)
}
