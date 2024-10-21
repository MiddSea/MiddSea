package main

import (
	"fmt"
	piscine "piscine/pkg"
)

func main() {

	fmt.Println(piscine.RetainFirstHalfIndex("This is the 1st halfThis is the 2nd half"))
	fmt.Println(piscine.RetainFirstHalfIndex("A"))
	fmt.Println(piscine.RetainFirstHalfIndex(""))
	fmt.Println(piscine.RetainFirstHalfIndex("Hello World"))
	//more tests by Seán
	fmt.Println(piscine.RetainFirstHalfIndex("AB"))
	fmt.Println(piscine.RetainFirstHalfIndex("áb")) // fails using RetainFirstHalf


	fmt.Println(piscine.RetainFirstHalf("This is the 1st halfThis is the 2nd half"))
	fmt.Println(piscine.RetainFirstHalf("A"))
	fmt.Println(piscine.RetainFirstHalf(""))
	fmt.Println(piscine.RetainFirstHalf("Hello World"))
	//more tests by Seán
	fmt.Println(piscine.RetainFirstHalf("AB"))
	fmt.Println(piscine.RetainFirstHalf("áb")) // fails using RetainFirstHalf

	fmt.Println(piscine.RetainFirstHalf("This is the 1st halfThis is the 2nd half"))
	fmt.Println(piscine.RetainFirstHalf("A"))
	fmt.Println(piscine.RetainFirstHalf(""))
	fmt.Println(piscine.RetainFirstHalf("Hello World"))
	//more tests by Seán
	fmt.Println(piscine.RetainFirstHalf("AB"))
	fmt.Println(piscine.RetainFirstHalf("áb")) // fails using RetainFirstHalf

}

/*retainfirsthalf
Instructions
Write a function called RetainFirstHalf() that takes a string as an argument and returns the first half of this string.

If the length of the string is odd, round it down.
If the string is empty, return an empty string.
If the string length is equal to one, return the string.
Expected function
func RetainFirstHalf(str string) string {

}
Usage
Here is a possible program to test your function:

package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.RetainFirstHalf("This is the 1st halfThis is the 2nd half"))
	fmt.Println(piscine.RetainFirstHalf("A"))
	fmt.Println(piscine.RetainFirstHalf(""))
	fmt.Println(piscine.RetainFirstHalf("Hello World"))
}
And its output:

$ go run . | cat -e
This is the 1st half$
A$
$
Hello$
*/