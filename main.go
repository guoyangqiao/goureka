package main

func main() {
	s := do("wefi2p34kp2k")
	println(s)
}

func do(kk string) string {
	n := len(kk)
	if n <= 3 {
		return kk
	}
	return do(kk[:n-3]) + "," + kk[n-3:]
}
