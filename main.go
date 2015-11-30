package main

import (
	"fmt"
	"strconv"
)


func main() {
	// accept number from user
	// tell how many spaces it is (powers of 10).
	// breakdown into expanded form
	
	var userNumber int
	var rawuserNumber string
	fmt.Printf("Please enter a non negative number: ")
	fmt.Scanf("%s", &rawuserNumber)
	for _ , value := range rawuserNumber  {
		switch {
			case value >= '0' && value <= '9':
			default :
				fmt.Println("Not an Int")
				break
		}
	}
	userNumber, err := strconv.Atoi(rawuserNumber)
	fmt.Println(err)
	
	
	
	fmt.Println(userNumber)
	
}