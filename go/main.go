package main

import "net/http"

func AddHttpHandler(w http.ResponseWriter, r *http.Request) {
	println("Received a request")
	sum := add(3, 5)
	println("The sum of 3 and 5 is:", sum)
	w.Write([]byte("Hello, World!"))
}

func main() {
	http.HandleFunc("/hello", AddHttpHandler)
	http.ListenAndServe(":8080", nil)
}

func add(a, b int) int {
	return a + b
}
