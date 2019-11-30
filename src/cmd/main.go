package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/brucebales/gocards/src/internal/languages"
)

var dictionary = languages.GetDictionary()

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to GoCards! Select a language:")
	fmt.Println("1. German")
	fmt.Println("2. Spanish")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare(text, "1") == 0 {
			cardGame("German", reader)
		}
		if strings.Compare(text, "2") == 0 {
			cardGame("Spanish", reader)
		}
	}
}

func cardGame(language string, reader *bufio.Reader) {

	//Randomly select a word from the dictionary
	index := rand.Intn(len(dictionary.Words))
	foreignWord := dictionary.Words[index]
	var foreignVersion string

	//Loop over translations, match to the chosen language
	var found bool
	for _, w := range foreignWord.Translations {
		if w.Language == language {
			found = true
			foreignVersion = w.Word
			fmt.Printf("Word in %s: %s", language, w.Word)
		}
	}
	//If translation for the randomly selected word in the chosen language isn't found,
	//choose a new word by restarting the cardGame function
	if !found {
		cardGame(language, reader)
	}
	fmt.Print("-> ")

	translation := dictionary.GetTranslation(foreignVersion, language, "English")

	englishWord, _ := reader.ReadString('\n')
	englishWord = strings.Replace(englishWord, "\n", "", -1)

	if strings.Compare(englishWord, translation) == 0 {
		fmt.Print("Correct!\n\n")
		time.Sleep(1 * time.Second)
	} else {
		fmt.Print("Big oof\n\n")
		time.Sleep(1 * time.Second)
	}

	cardGame(language, reader)
}
