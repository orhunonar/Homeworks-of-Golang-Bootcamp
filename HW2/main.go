package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	args := os.Args[1:] //Take string from terminal as an input
	realstring := strings.Join(args, " ")
	newrealstring := strings.Split(realstring, " ") //Make array

	booklist := []string{ //Book List
		"Brain",
		"1984",
		"Silmarillion",
		"Drive",
		"Nutuk",
	}
	if newrealstring[0] == "list" { // Show the Book List
		for i := 0; i < len(booklist); i++ {
			fmt.Println(booklist[i])
		}

	} else if newrealstring[0] == "search" { //Search Algorithm

		for i := 0; i < len(booklist); i++ {
			if newrealstring[1] == booklist[i] {
				fmt.Println(booklist[i], " is valid on the list")
				break
			}
			if len(booklist) == i+1 {
				fmt.Println("The book is not valid")
			}

		}
	}

}
