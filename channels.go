package main

import (
	"fmt"
	"math/rand"
	"time"
)

func processNum(numchan chan int) {

	for num := range numchan {
		fmt.Println("Process Num:", num)
		time.Sleep(time.Millisecond * 250)
	}
}

func sum(result chan int, num1 int, num2 int) {
	numResult := num1 + num2
	result <- numResult
}

func task(done chan bool) {

	defer func() {
		done <- true
		fmt.Println("Task Done:")
	}()
	fmt.Println("Proccessing....")
}

func emailsender(emailchan chan string, done chan bool) {

	defer func() {
		done <- true
	}()
	for email := range emailchan {
		fmt.Println("Email sending At:", email)
		time.Sleep(time.Millisecond * 250)
	}
}
func main() {

	chan1 := make(chan int)
	chan2 := make(chan string)
	chan3 := make(chan bool)
	chan4 := make(chan float64)
	go func() {
		chan1 <- 10
	}()
	go func() {
		chan2 <- "Ping"
	}()
	go func() {
		chan3 <- true
	}()
	go func() {
		chan4 <- 20.34
	}()

	for i := 0; i < 4; i++ {
		select {
		case chan1val := <-chan1:
			fmt.Println("Received data from chan1", chan1val)
		case chan2val := <-chan2:
			fmt.Println("Received data from chan2", chan2val)
		case chan3val := <-chan3:
			fmt.Println("Received data from chan2", chan3val)
		case chan4val := <-chan4:
			fmt.Println("Received data from chan2", chan4val)

		}
		time.Sleep(1 * time.Second)
	}
	emailchan := make(chan string, 50)
	done := make(chan bool)

	go emailsender(emailchan, done)

	for i := 0; i < 50; i++ {
		emailchan <- fmt.Sprintf("%d@gmail.com", i)
	}

	fmt.Println("Done Sending")

	close(emailchan)
	<-done

	go task(done)

	<-done
	result := make(chan int)

	go sum(result, 4, 5)

	res := <-result

	fmt.Println(res)

	numChan := make(chan int) //create channel

	go processNum(numChan)

	for {
		numChan <- rand.Intn(100)
	}
	time.Sleep(time.Microsecond) //wait for millisecond
	messageChan := make(chan string)

	messageChan <- "ping"

	msg := <-messageChan

	fmt.Println(msg)
}
