package hangmanpackage

import (
	"bufio"
	"log"
	"os"
)

func Openfile() (mots []string) {

	f, err := os.Open("words.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		mots = append(mots, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return mots

}
