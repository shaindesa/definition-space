package cmd

type WordData struct{
	Word string `json:"value"`
	POS string `json:"partofspeech"`
	Definition string `json:"definition"`
	Example string `json:"example"`
}

type WordInfo struct{
	Word string `json:"word"`
	WordGroup []WordGroup `json:"meanings"`
}

type WordGroup struct{
	PartOfSpeech string `json:"partOfSpeech"`
	Definitions []Definition `json:"definitions"`
}

type Definition struct{
	Val string `json:"definition"`
	Example string `json:"example"`
}

