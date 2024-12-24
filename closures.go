package main

import "fmt"

func sequence() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
func main() {
	nextval := sequence()

	fmt.Println(nextval())
	fmt.Println(nextval())
	fmt.Println(nextval())
	fmt.Println(nextval())

	nextint := sequence()
	fmt.Println(nextint())
	fmt.Println(nextval())
	fmt.Println(nextval())
}
