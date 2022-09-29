package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type WordInfo struct{
	Word string `json:"word"`
	Meanings []Meaning `json:"meanings"`
}

type Meaning struct{
	PartOfSpeech string `json:"partOfSpeech"`
	Definitions []Definition `json:"definitions"`
}

type Definition struct{
	Val string `json:"definition"`
	Example string `json:"example"`
}

func main() {
	word := "hello"
	url := "https://api.dictionaryapi.dev/api/v2/entries/en/" + word
	
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var w []WordInfo

	json.Unmarshal(body, &w)
	count := 0

	for _, wordtype := range w[0].Meanings{
		for _, definition := range wordtype.Definitions{
			count++
			fmt.Printf("Definition %v\n", count)
			fmt.Println(definition.Val)
			fmt.Printf("%v\n\n", definition.Example)
		}
	}
}