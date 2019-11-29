package languages

func GetDictionary() Dictionary {
	return Dictionary{
		Words: []Word{
			{
				Translations: []Translation{
					Translation{
						Language: "english",
						Word:     "cat",
					},
					Translation{
						Language: "german",
						Word:     "katze",
					},
				},
			},
			{
				Translations: []Translation{
					Translation{
						Language: "english",
						Word:     "house",
					},
					Translation{
						Language: "german",
						Word:     "hause",
					},
				},
			},
			{
				Translations: []Translation{
					Translation{
						Language: "english",
						Word:     "man",
					},
					Translation{
						Language: "german",
						Word:     "mann",
					},
				},
			},
			{
				Translations: []Translation{
					Translation{
						Language: "english",
						Word:     "ugly",
					},
					Translation{
						Language: "german",
						Word:     "hasslich",
					},
				},
			},
		},
	}
}
