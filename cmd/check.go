package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Lookup a definition in your local dictionary",
	Long: `Go through the JSON file in $HOME/.definition-space/dictionary.json and look for all words that match the word given in as an argument. For example:

definition-space check [word]


`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		CheckDir()
		file := os.Getenv("HOME") + "/.definition-space/dictionary.json"

		words, err := ReadJSON(file)
		if err != nil{
			fmt.Println("You don't have a local dictionary yet. Get started by calling `definition-space add [word]`")
			return
		}

		word := strings.ToLower(args[0])
		found := false

		for key, val := range words{
			if val.Word == word{
				fmt.Printf("%v\t'%v'\n", key, val.Word)
				fmt.Printf("%v\n", val.POS)
				fmt.Printf("%v\n", val.Definition)
				if val.Example != ""{
					fmt.Printf("Example: \"%v\"\n", val.Example)
				}
				fmt.Println()
				found = true
			}
		}

		if !found{
			fmt.Println("Word not found in local dictionary")
		}
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
