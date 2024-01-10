package hangmanpackage

import (
	"bufio"
	"math/rand"
	"os"
)

func Openfile(filePath string) (mot string, err error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var tabmots []string

	for scanner.Scan() {
		tabmots = append(tabmots, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	index := rand.Intn(len(tabmots))
	mot = tabmots[index]
	return mot, nil
}
