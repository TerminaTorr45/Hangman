package hangman

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func Userenter(mots *Hangman, resume bool, mot1, motCache1 string, vie1 Hangman) {
	var vie Hangman
	if resume {
		vie = vie1
	} else {
		vie.Cont = 10
	}

	scanner := bufio.NewScanner(os.Stdin)
	var mot, motCache string

	if resume {
		mot = mot1
		motCache = motCache1
	} else {
		// Choisissez un mot au hasard
		index := rand.Intn(len(mots.Words))
		mot = mots.Words[index]
		motCache = CacherMot(mot)
	}

	fmt.Println("Good Luck, you have", vie.Cont, "attempts.")
	fmt.Println(motCache)

	for vie.Cont > 0 {
		fmt.Println()
		fmt.Print("Choose: ")
		scanner.Scan()
		fmt.Println()
		input := strings.ToLower(scanner.Text())

		if len(input) > 1 {
			if strings.ToUpper(input) == "STOP" {
				SaveGame(mot, motCache, vie)
				fmt.Println("Game Saved in save.txt.")
				return
			} else if input == mot {
				fmt.Println("CONGRATS !")
				fmt.Println()
				break
			} else if vie.Cont == 1 {
				vie.Cont--
				fmt.Println("You have run out of lives. The word was:", mot)
				break
			} else {
				vie.Cont -= 2
				fmt.Println("The word is incorrect. You have", vie.Cont, "attempts remaining.")

			}
		} else if len(input) == 1 {
			if strings.Contains(mot, input) {
				motCache = RevelerLettres(mot, motCache, input)
				fmt.Println(motCache)
			} else {
				vie.Cont--
				fmt.Println("The letter is not in the word. You have", vie.Cont, "attempts remaining.")
				fmt.Println(motCache)
			}
		} else {
			fmt.Println("Please enter at least one letter or a complete word.")
		}

		if mot == motCache {
			fmt.Println("CONGRATS !")
			fmt.Println()
			break
		}

		Displaypendue(&vie)

		if vie.Cont == 0 {
			fmt.Println("You have run out of lives. The word was:", mot)
		}
	}
}
