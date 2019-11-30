package languages

func GetDictionary() Dictionary {
	return Dictionary{
		Words: []Word{
			{
				Translations: []Translation{
					Translation{
						Language: "English",
						Word:     "cat",
					},
					Translation{
						Language: "German",
						Word:     "katze",
					},
					Translation{
						Language: "Spanish",
						Word:     "gato",
					},
				},
			},
			{
				Translations: []Translation{
					Translation{
						Language: "English",
						Word:     "house",
					},
					Translation{
						Language: "German",
						Word:     "hause",
					},
					Translation{
						Language: "Spanish",
						Word:     "casa",
					},
				},
			},
			{
				Translations: []Translation{
					Translation{
						Language: "English",
						Word:     "man",
					},
					Translation{
						Language: "German",
						Word:     "mann",
					},
					Translation{
						Language: "Spanish",
						Word:     "hombre",
					},
				},
			},
			{
				Translations: []Translation{
					Translation{
						Language: "English",
						Word:     "ugly",
					},
					Translation{
						Language: "German",
						Word:     "hasslich",
					},
					Translation{
						Language: "Spanish",
						Word:     "feo",
					},
				},
			},
		},
	}
}
