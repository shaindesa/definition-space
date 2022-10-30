/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// wordsCmd represents the words command
var wordsCmd = &cobra.Command{
	Use:   "words",
	Short: "Prints a comma-separated list of all of the words in your local dictionary",
	Long: `Queries your local dictionary at $HOME/.definition-space/dictionary.json and prints the words found there, comma separated.
	Example usage: definition-space words
	
	`,
	Run: func(cmd *cobra.Command, args []string) {
		CheckDir()
		file := os.Getenv("HOME") + "/.definition-space/dictionary.json"

		words, err := ReadJSON(file)
		if err != nil || len(words) == 0 {
			fmt.Println("You don't have a local dictionary yet. Get started by calling `definition-space add [word]`")
			return
		}

		for _, val := range words{
			fmt.Printf("%v, ", val.Word)
		}
		fmt.Printf("\n\n")
		
	},
}

func init() {
	rootCmd.AddCommand(wordsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// wordsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// wordsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
