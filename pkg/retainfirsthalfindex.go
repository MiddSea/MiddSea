package piscine 

import "strings"

func RetainFirstHalfIndex(str string) string {
	var lenStr int = len(str)
	var midStr int = lenStr/2 // will round down
    if lenStr == 1 {
		midStr = 1
	}
	println("Str:", str, "lenStr:", lenStr, "midStr", midStr)
	// return str[0:midStr] 
	return str[0:midStr] 
}

func RetainFirstHalfRange(str string) string {
	// try using a range []rune str
	var runeStr []rune = []rune(str)
	var lenStr int = len(runeStr)
	var midStr int = lenStr/2 // will round down
	var firstHalfRuneRange []rune = runeStr[:midStr]
	if lenStr == 1 {
		midStr = 1
	}
	println("Str:", str, "lenStr:", lenStr, "midStr", midStr)
	// return str[0:midStr]
	return string(firstHalfRuneRange) // TO DO
}


func RetainFirstHalf(str string) string {
	if len(str) == 0 {
		return ""
	} else if len(str) == 1 {
		return (str)
	} else {
		var res strings.Builder
		i := 0
		for i = 0; i < int(len(str)/2); i++ {
			res.WriteRune(rune(str[i]))
		}
		return res.String()
	}
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