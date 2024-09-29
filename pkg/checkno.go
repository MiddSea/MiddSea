package piscine

func CheckNumber(arg string) bool {
    for _, r := range arg {
		if r >= '0' && r <= '9' {
			return true
		} /* else {	
			return false
 		} */
	}
	return false
}
/* checknumber
Instructions
Write a function that takes a string as an argument and returns true if the string contains any number, otherwise return false.

Expected function
func CheckNumber(arg string) bool {
}
Usage
Here is a possible program to test your function:

package main

import (
	"fmt"
)

func main() {
	fmt.Println(CheckNumber("Hello"))
	fmt.Println(CheckNumber("Hello1"))
}
And its output:

$ go run .
false
true
$
 */