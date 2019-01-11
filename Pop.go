package main

var pc [256]byte

func init() {
	for e := range pc {
		pc[e] = pc[e>>1] + byte(e)&0X01
	}
	println()
}

func main() {
	s := "E"
	c := countBit1(s)
	println(c)
}
func countBit1(s string) int {
	c := []byte(s)
	var sum int
	for _, v := range c {
		sum += int(pc[v])
	}
	return sum
}
