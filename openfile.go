package hangmanpackage

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
)

func Openfile() (mot string) {

	f, err := os.Open("words.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var tabmots []string

	for scanner.Scan() {
		tabmots = append(tabmots, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	index := rand.Intn(len(tabmots))
	mot = tabmots[index]

	fmt.Println(mot)
	return mot

}
