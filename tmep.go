package main

import "fmt"

func main() {
	var tmp float64
	var funit, tounit string
	fmt.Printf("Enter The Temperature: ")
	fmt.Scanln(&tmp)
	fmt.Printf("Enter The Current Unit Of Temperature: ")
	fmt.Scanln(&funit)
	fmt.Printf("Enter The Unit Where You Want To Move: ")
	fmt.Scanln(&tounit)
	var ans float64
	if funit == string('C') {
		if tounit == string('F') {
			ans = ((tmp * 9) / 5) + 32
		} else if tounit == string('K') {
			ans = tmp + 273
		}
	} else if funit == string('F') {
		if tounit == string('C') {
			ans = ((tmp * 9) / 5) + 32
		} else if tounit == string('K') {
			ans = (((tmp - 32) / 9) * 5) + 273
		}
	} else if funit == string('K') {
		if tounit == string('F') {
			ans = (((tmp - 273) / 5) * 9) + 32
		} else if tounit == string('C') {
			ans = tmp - 273
		}
	}
	fmt.Println("The Final Temparature is: ", ans)
}
