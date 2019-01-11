package main

type none uint32

const (
	fooo none = 1<<iota + 1
	bar
	qux
)

func main() {
	println(fooo)
	println(bar)
	println(qux)
}
