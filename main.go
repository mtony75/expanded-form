package main

import (
	"fmt"
	"strconv"
)

func prove_int(uinput string) bool {
	notInt := false
	for _, value := range uinput {
		switch {
		case value >= '0' && value <= '9':
		default:
			//				fmt.Println("Not an Int")
			//				break
			notInt = true
		}
	}
	return notInt
}

func main() {
	// accept number from user
	// tell how many spaces it is (powers of 10).
	// breakdown into expanded form

	var userNumber, powerOf int
	var rawuserNumber string
	fmt.Printf("Please enter a non negative number: ")
	fmt.Scanf("%s", &rawuserNumber)
	notInt := false
	notInt = prove_int(rawuserNumber)
	//	userNumber, err := strconv.Atoi(rawuserNumber)
	// 	fmt.Println(err)
	if notInt {
		fmt.Println("You did not enter an Integer")

	} else {
		fmt.Println("That is an Integer")
		userNumber, _ = strconv.Atoi(rawuserNumber)
		fmt.Println(userNumber)
	}

	// use algorithm to find how how big a number is and
	// use that to display the expanded form.
	fmt.Println("Start of For Loop")

	for powerOf = 10; userNumber != (userNumber % powerOf); powerOf = powerOf * 10 {
		//		fmt.Println(powerOf)

	}
}
