/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
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


// lookupCmd represents the lookup command
var lookupCmd = &cobra.Command{
	Use:   "lookup",
	Short: "Look up definitions. Does not commit anything to memory.",
	Long: `Use this command alongside a single real English word to lookup the word's definition(s)
	in the dictionary.
	For example:
	definition-space lookup apple
	-will look up the definition of 'apple' and return various definitions in STDout`,
	Args: cobra.ExactArgs(1),


	Run: func(cmd *cobra.Command, args []string) {
		word := args[0]
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
				fmt.Println(wordtype.PartOfSpeech)
				fmt.Println(definition.Val)
				fmt.Printf("%v\n\n", definition.Example)
				fmt.Scanln()
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(lookupCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lookupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lookupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
