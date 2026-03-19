package main

func main() {
	println("Hello, World!")
	sum := add(3, 5)
	println("The sum of 3 and 5 is:", sum)
}

func add(a, b int) int {
	return a + b
}
