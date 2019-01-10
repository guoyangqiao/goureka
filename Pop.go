package main

var bc = [256]byte{}

func init() {
	for x := range bc {
		bc[x] = bc[x>>1] + byte(x&0x01)
	}
}

func popCnt(x uint32) int {
	cc := int(
		bc[byte(x>>(0*8))] +
			bc[byte(x>>(1*8))] +
			bc[byte(x>>(2*8))] +
			bc[byte(x>>(3*8))])
	return cc
}

func diffBit(x, y uint32) int {
	return popCnt(x ^ y)
}

func main() {
	cnt := popCnt(255)
	println(cnt)
	bit := diffBit(165, 187)
	println(bit)
}
