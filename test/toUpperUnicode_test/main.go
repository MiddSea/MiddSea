package main

import (
	"fmt"
	"unicode"
)

func main() {
	const ucG = 'ú'
	const lcG = 'Ú'
	
	fmt.Printf("%#U\n", unicode.ToUpper(ucG))

	fmt.Printf("%#U\n", unicode.ToLower(lcG))


}
