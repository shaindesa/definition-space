package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a word from your local dictionary",
	Long: `Removes the first word in dictionary that matches word provided in argument.
	Usage: 'definition-space remove [word]'
	
	
	`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		CheckDir()
		file := os.Getenv("HOME") + "/.definition-space/dictionary.json"

		words, err := ReadJSON(file)
		if err != nil || len(words) == 0 {
			fmt.Println("You don't have a local dictionary yet. Get started by calling `definition-space add [word]`")
			return
		}

		found := false

		for ind, word := range words{
			if word.Word == args[0]{
				words = append(words[:ind], words[ind+1:]...)
				found = true
			}
		}

		if !found{
			fmt.Printf("Could not find word %v in local dictionary.\n\n", args[0])
			return
		}

		jsondata, err := json.Marshal(words)
		if err != nil{
			log.Fatal(err)
		}

		err = WriteJSON(file, jsondata)
		if err != nil{
			fmt.Println("Error removing word from dictionary")
			log.Fatal(err)
		}
		fmt.Printf("Successfully removed %v from local dictionary\n\n", args[0])

	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
