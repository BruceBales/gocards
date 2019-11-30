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

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare(text, "1") == 0 {
			german(reader)
		}
	}
}

func german(reader *bufio.Reader) {

	index := rand.Intn(len(dictionary.Words))

	foreignWord := dictionary.Words[index]

	var foreignVersion string

	for _, w := range foreignWord.Translations {
		if w.Language == "german" {
			foreignVersion = w.Word
			fmt.Println("Word in German: ", w.Word)
		}
	}
	fmt.Print("-> ")

	translation := dictionary.GetTranslation(foreignVersion, "german", "english")

	englishWord, _ := reader.ReadString('\n')
	englishWord = strings.Replace(englishWord, "\n", "", -1)

	if strings.Compare(englishWord, translation) == 0 {
		fmt.Print("Correct!\n\n")
		time.Sleep(1 * time.Second)
	} else {
		fmt.Print("Big oof\n\n")
		time.Sleep(1 * time.Second)
	}

	german(reader)
}
