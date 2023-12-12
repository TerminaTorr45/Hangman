package hangman

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"

	"github.com/common-nighthawk/go-figure"
)

func Openfile1(filename string) {
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

	Userenter1(&mots, false, "", "", Hangman{})
}

func Userenter1(mots *Hangman, resume bool, mot1, motCache1 string, vie1 Hangman) {
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
	myFigure := figure.NewFigure(motCache, "", true)
	myFigure.Print()

	for vie.Cont > 0 {
		fmt.Println()
		fmt.Print("Choose: ")
		scanner.Scan()
		fmt.Println()
		input := strings.ToLower(scanner.Text())

		if len(input) > 1 {
			if strings.ToUpper(input) == "STOP" {
				SaveGame1(mot, motCache, vie)
				fmt.Println("Game Saved in save.txt.")
				return
			} else if input == mot {
				fmt.Println()
				fmt.Println()
				fmt.Println()
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
				myFigure := figure.NewFigure(motCache, "", true)
				myFigure.Print()
			} else {
				vie.Cont--
				fmt.Println("The letter is not in the word. You have", vie.Cont, "attempts remaining.")
				myFigure := figure.NewFigure(motCache, "", true)
				myFigure.Print()
			}
		} else {
			fmt.Println("Please enter at least one letter or a complete word.")
		}

		if mot == motCache {
			fmt.Println()
			fmt.Println()
			fmt.Println()
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

func SaveGame1(word, hiddenWord string, hang Hangman) {
	saveData := struct {
		Word       string
		HiddenWord string
		Hangman    Hangman
	}{word, hiddenWord, hang}

	file, err := os.Create("save.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(saveData)
	if err != nil {
		log.Fatal(err)
	}
}

func ResumeGame1(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var saveData struct {
		Word       string
		HiddenWord string
		Hangman    Hangman
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&saveData)
	if err != nil {
		log.Fatal(err)
	}

	Userenter1(nil, true, saveData.Word, saveData.HiddenWord, saveData.Hangman)
}
