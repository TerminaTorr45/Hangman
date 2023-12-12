package main

import (
	"fmt"
	"hangman"
	"os"
)

func main() {
	var fileName string

	if len(os.Args) < 2 {
		fmt.Println("Select a words file")
		os.Exit(1)
	}

	if len(os.Args) == 3 && os.Args[1] == "--startWith" {
		fileName = os.Args[2]
	} else {
		hangman.Openfile(os.Args[1])
		return
	}

	hangman.ResumeGame(fileName)
}
