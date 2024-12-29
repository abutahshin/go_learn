package main

import (
	"fmt"
	"time"
)

func task1(id int) {
	fmt.Println("Doing task", id)
}
func main() {
	for i := 0; i <= 10; i++ {
		go task1(i)
	}
	time.Sleep(time.Second * 1)
}
