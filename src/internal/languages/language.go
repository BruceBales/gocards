package languages

import (
	"strings"
)

type Dictionary struct {
	Words []Word
}

type Word struct {
	Translations []Translation
}

type Translation struct {
	Language string
	Word     string
}

//GetTranslation is an abomination from the pits of hell
func (d *Dictionary) GetTranslation(givenWord, from, to string) string {
	for _, word := range d.Words {
		for _, translation := range word.Translations {
			if strings.Compare(translation.Language, from) == 0 && strings.Compare(translation.Word, givenWord) == 0 {
				for _, translateTo := range word.Translations {
					if strings.Compare(translateTo.Language, to) == 0 {
						return translateTo.Word
					}
				}
			}
		}
	}
	return "Fuck"
}
