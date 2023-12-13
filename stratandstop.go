package hangmanpackage

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func Start() {

	var fileName string

	if len(os.Args) < 2 {
		fmt.Println("Select a words file")
		os.Exit(1)
	}

	if len(os.Args) == 3 && os.Args[1] == "--startWith" {
		fileName = os.Args[2]
	} else {
		Openfile(os.Args[1])
		return
	}
	ResumeGame(fileName)
}

func SaveGame(word, hiddenWord string, hang Hangman) {
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

func ResumeGame(fileName string) {
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

	Userenter(nil, true, saveData.Word, saveData.HiddenWord, saveData.Hangman)
}
