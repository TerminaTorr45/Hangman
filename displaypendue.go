package hangmanpackage

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Displaypendue(vie *Hangman) {

	pendue, err := os.Open("hangman.txt")
	var tab []string

	if err != nil {
		log.Fatal(err)
	}
	defer pendue.Close()

	scanner := bufio.NewScanner(pendue)

	for scanner.Scan() {
		tab = append(tab, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if vie.Cont == 9 {

		index := 6
		pendue1 := tab[index]
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println(pendue1)

	}

	if vie.Cont == 8 {

		fmt.Println()
		for i := 9; i <= 14; i++ {
			pendue1 := tab[i]
			fmt.Println(pendue1)
		}

	}

	if vie.Cont == 7 {

		fmt.Println()
		for i := 16; i <= 22; i++ {
			pendue1 := tab[i]
			fmt.Println(pendue1)
		}

	}
	if vie.Cont == 6 {

		fmt.Println()
		for i := 24; i <= 30; i++ {
			pendue1 := tab[i]
			fmt.Println(pendue1)
		}

	}
	if vie.Cont == 5 {

		fmt.Println()
		for i := 32; i <= 38; i++ {
			pendue1 := tab[i]
			fmt.Println(pendue1)
		}

	}
	if vie.Cont == 4 {

		fmt.Println()
		for i := 40; i <= 46; i++ {
			pendue1 := tab[i]
			fmt.Println(pendue1)
		}

	}
	if vie.Cont == 3 {

		fmt.Println()
		for i := 48; i <= 54; i++ {
			pendue1 := tab[i]
			fmt.Println(pendue1)
		}

	}
	if vie.Cont == 2 {

		fmt.Println()
		for i := 56; i <= 62; i++ {
			pendue1 := tab[i]
			fmt.Println(pendue1)
		}

	}
	if vie.Cont == 1 {

		fmt.Println()
		for i := 64; i <= 70; i++ {
			pendue1 := tab[i]
			fmt.Println(pendue1)
		}

	}
	if vie.Cont == 0 {

		fmt.Println()
		for i := 72; i <= 78; i++ {
			pendue1 := tab[i]
			fmt.Println(pendue1)
		}

	}

}
