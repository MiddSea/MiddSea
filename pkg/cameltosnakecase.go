package piscine

import (
	"unicode"
	"fmt"
	//
)
func isUpperRune(r rune) bool {
	if r >= 'A' && r <= 'Z' {
		return true
	}
	return false
}

func isLowerRune(r rune) bool {
	if r >= 'a' && r <= 'z' {
		return true
	}
	return false
}

func isLetterString(s string) bool {
	for _, r := range s {
		if isUpperRune(r) || isLowerRune(r) {
			return true
		}
	}
	return false
}
func IsCamelCase(s string) bool {
	lastIndex := len(s)-1
	if !isLetterString(s) {
		return false
	} 
	for i, r := range s {
		if i > 0 && i <= lastIndex {
		// TO DO 
		//if isUpperRune(prevRune) && isUpperRune(r) {

		//}
		}
	}
}

func PrintIndexAndRune(i int, r rune) {
	fmt.Printf("i:%2d r: %c | ", i, r)
}

func CheckCamelCase(s string) bool {
	// lowerCamelCase
	// UpperCamelCase
	// Rules for writing in camelCase:
	var lenS = len(s)
    var currentLenRune int
	for i, r := range s {
		if i == 0 && (unicode.IsLower(r) || unicode.IsUpper(r)) {
			if i == 0 { // skip first rune if is lower or upper letter
				fmt.Print("FST")
				PrintIndexAndRune(i, r)
				continue 
			}

		} 
		PrintIndexAndRune(i, r)
        currentLenRune = lenS - i // in byte
	    if i + currentLenRune == lenS {
			fmt.Printf ("*** last i: %d currentLenRune: %d  lenS: %d ", i, currentLenRune, lenS)
		}
	}
	fmt.Print("len(s) = ", len(s))

	return true
	// The word does not end on a capitalized letter (CamelCasE).
	// No two capitalized letters shall follow directly each other (CamelCAse).
	// Numbers or punctuation are not allowed in the word anywhere (camelCase1).
	
}

/* func SnakeCaseStr(s string) string {

} */

func CamelToSnakeCase(s string) string {
	// lenS = len(s)

	if s == "" {
		return "" // return empty string as per requirement.
	} else if true /* TO DO */ {
       // do camel case


	} else {
		return s  // string is not camel case, return unchanged
	}
}




/*cameltosnakecase


Instructions
Write a function that converts a string from camelCase to snake_case.

If the string is empty, return an empty string.
If the string is not camelCase, return the string unchanged.
If the string is camelCase, return the snake_case version of the string.
For this exercise you need to know that camelCase has two different writing alternatives that will be accepted:

lowerCamelCase
UpperCamelCase
Rules for writing in camelCase:

The word does not end on a capitalized letter (CamelCasE).
No two capitalized letters shall follow directly each other (CamelCAse).
Numbers or punctuation are not allowed in the word anywhere (camelCase1).
Expected function
func CamelToSnakeCase(s string) string{

}
Usage
Here is a possible program to test your function:

package main

import (
	"fmt"
)

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}
And its output:

$ go run .
Hello_World
hello_World
camel_Case
CAMELtoSnackCASE
camel_To_Snake_Case
hey2
*/