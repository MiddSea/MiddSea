package main

import (
	"fmt"
	"os"
)

var WhtSpace string = "\n\t "

func isWhtSpace(r rune) bool {
	for _, whtSp := range WhtSpace {
		if r == whtSp {
			return true
		}
	}
	return false
}

func main() {
	args := os.Args[1:]
	lenArgs := len(args)
	fmt.Println("args:", args, "lenArgs", lenArgs)
	cleanStr := ""
	gotWord := false
	wasSpace := false
	for _, r := range args[0] {
		// if is a non space after some space add " "
		if !isWhtSpace(r) {
			if gotWord && wasSpace {
				cleanStr += string(" ")
				wasSpace = false
			}
			cleanStr += string(r)
			gotWord = true
		}
		if gotWord && isWhtSpace(r) {
			wasSpace = true
		}

	}
	fmt.Println("end:", cleanStr)
}
