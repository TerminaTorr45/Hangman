package hangman

import (
	"bufio"
	"log"
	"os"
)

func Openfile(filename string) {
	var mots Hangman
	f, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		mots.Words = append(mots.Words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	Userenter(&mots, false, "", "", Hangman{})
}
