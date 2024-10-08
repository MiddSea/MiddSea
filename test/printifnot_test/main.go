package main

import (
	"fmt"
	piscine "piscine/pkg"
)

func main() {
	fmt.Print(piscine.PrintIfNot("abcdefz"))
	fmt.Print(piscine.PrintIfNot("abc"))
	fmt.Print(piscine.PrintIfNot(""))
	fmt.Print(piscine.PrintIfNot("14"))
}

/* printifnot
Instructions
Write a function that takes a string as an argument and returns the letter G if the argument length is less than 3, otherwise returns Invalid Input followed by a newline \n.

If it's an empty string return G followed by a newline \n.
Expected function
func PrintIfNot(str string) string {

}
Usage
Here is a possible program to test your function:

package main

import (
	"fmt"
)

func main() {
	fmt.Print(PrintIfNot("abcdefz"))
	fmt.Print(PrintIfNot("abc"))
	fmt.Print(PrintIfNot(""))
	fmt.Print(PrintIfNot("14"))
}
And its output:

$ go run . | cat -e
Invalid Input$
Invalid Input$
G$
G$ 
*/