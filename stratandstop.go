package hangmanpackage

import (
	"encoding/json"
	"log"
	"os"
)

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
