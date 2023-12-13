package hangmanpackage

import (
	"bufio"
	"log"
	"math/rand"
	"os"
)

func Openfile() (mot string) {

	f, err := os.Open(os.Args[1])

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
	return mot

}
