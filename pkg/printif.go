package piscine

func PrintIf(str string) string {
    if len(str) >= 3 || len(str) == 0 {
		return "G\n"
	}
	return "Invalid Input\n"
}

/*printif
Instructions
Write a function that takes a string as an argument and returns the letter G followed by a newline \n if the argument length is more or equal than 3, otherwise returns Invalid Input followed by a newline \n.

If it's an empty string return G followed by a newline \n.
Expected function
func PrintIf(str string) string {

}
Usage
Here is a possible program to test your function:

package main

import (
	"fmt"
)

func main() {
	fmt.Print(PrintIf("abcdefz"))
	fmt.Print(PrintIf("abc"))
	fmt.Print(PrintIf(""))
	fmt.Print(PrintIf("14"))
}
And its output:

$ go run . | cat -e
G$
G$
G$
Invalid Input$*/