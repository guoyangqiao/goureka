package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"os"
	"reflect"
	"strings"
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
	var meHint string
	Args := os.Args
	if len(Args) > 2 {
		meHint = Args[1]
	}
	println("using ", meHint)
	var me reflect.Value
	switch strings.ToUpper(meHint) {
	case "SHA384":
		me = reflect.ValueOf(sha512.Sum384)
	case "SHA512":
		me = reflect.ValueOf(sha512.Sum512)
	default:
		me = reflect.ValueOf(sha256.Sum256)
	}
	i := callMe(me, []byte("A"))
	j := callMe(me, []byte("X"))
	count := BitCount(i[:], j[:])
	println(count)
}

func callMe(me reflect.Value, x []byte) [32]uint8 {
	return me.Call([]reflect.Value{reflect.ValueOf(x)})[0].Interface().([32]uint8)
}
