package main

import "fmt"

func fibo() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
func main() {
	fib := fibo()
	var n int
	fmt.Println("Enter n:")
	fmt.Scan(&n)
	var result int
	for i := 0; i < n; i++ {
		result = fib()
	}
	fmt.Println("The Nth Fibonacci value is: ", result)
}
