/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
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

// learnCmd represents the learn command
var learnCmd = &cobra.Command{
	Use:   "learn",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// learnCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// learnCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
