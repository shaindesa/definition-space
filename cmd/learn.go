package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var learnCmd = &cobra.Command{
	Use:   "learn",
	Short: "Bring up to five words from your dictionary to help memorise them",
	Long: `learn is used to help you memorise the words you have added to your local dictionary. Call learn with 'definition-space learn' and you will be given up to five random words in your dictionary.
Each word in your dictionary has a JSON tag of 'learns.' Everytime learn selects a word it adds 1 to that total. When total learns reaches 4, the word will no longer show up in learn.

Usage: definition-space learn

`,
	Run: func(cmd *cobra.Command, args []string) {
		CheckDir()
		file := os.Getenv("HOME") + "/.definition-space/dictionary.json"

		words, err := ReadJSON(file)
		if err != nil || len(words) == 0 {
			fmt.Println("You don't have a local dictionary yet. Get started by calling `definition-space add [word]`")
			return
		}

		eligible_words := []WordData{}

		for _, val := range words{
			if val.Learns < 4 {
				eligible_words = append(eligible_words, val)
			}
		}

		if len(eligible_words) == 0{
			fmt.Println("You've learned all the words in your dictionary! Add some more with `definition-space add [word]`")
			return
		}

		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(eligible_words), func(i, j int) { eligible_words[i], eligible_words[j] = eligible_words[j], eligible_words[i]})
		
		var learning_words []WordData
		if len(eligible_words) >= 5{
			learning_words = eligible_words[:5]
		} else{
			learning_words = eligible_words
		}

		fmt.Printf("\nWord learning mode shows you up to five randomly selected words from your local dictionary. Go over your definitions here to help remember them!\n\n")
		
		for k, v := range learning_words{
			daysSince := (time.Now().Unix() - v.TimeAdded) / 86400
			fmt.Printf("(%v/%v)\t\"%v\"\n", k+1, len(learning_words), v.Word)
			fmt.Println(v.Definition)
			if v.Example != ""{
				fmt.Println(v.Example)
			}
			if daysSince == 1{
				fmt.Printf("(Added %v day ago)\n\n", daysSince)
			} else{
				fmt.Printf("(Added %v days ago)\n\n", daysSince)
			}

			// Ask user whether they want to continue
			var response string
			fmt.Println("(Press ENTER to continue learning, or type anything else to exit)")
			fmt.Scanln(&response)
			if response != ""{
				break
			}
			v.Learns++

			for ind, word := range words{
				if word.Word == v.Word && word.Definition == v.Definition{
					other_words := words[ind+1:]
					words = append(words[:ind], v)
					words = append(words, other_words...)
				}
			}
		}

		jsondata, err := json.Marshal(words)
		if err != nil{
			log.Fatal(err)
		}

		WriteJSON(file, jsondata)
	},
}

func init() {
	rootCmd.AddCommand(learnCmd)
}
