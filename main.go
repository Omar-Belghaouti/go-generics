package main

type Arg interface {
	string | int
}

func print[T Arg](x T) {
	println(x)
}

func main() {
	print("Hello")
	print(3)
}
