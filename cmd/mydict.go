package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var mydictCmd = &cobra.Command{
	Use:   "mydict",
	Short: "Prints out your entire local dictionary into STDOut",
	Long: `Finds your dictionary file in $HOME/.defniition-space and prints out its contents. Will initialise the directory if necessary.
	Usage: definition-space mydict
	
	`,
	Run: func(cmd *cobra.Command, args []string) {
		CheckDir()
		file := os.Getenv("HOME") + "/.definition-space/dictionary.json"

		words, err := ReadJSON(file)
		if err != nil || len(words) == 0 {
			fmt.Println("You don't have a local dictionary yet. Get started by calling `definition-space add [word]`")
			return
		}

		for key, val := range words{
			daysSince := (time.Now().Unix() - val.TimeAdded) / 86400
			fmt.Printf("%v\t\"%v\"\n", key+1, val.Word)
			fmt.Printf("%v\n", val.POS)
			fmt.Printf("%v\n", val.Definition)
			if val.Example != ""{
				fmt.Printf("Example: \"%v\"\n", val.Example)
			}
			if daysSince == 1{
				fmt.Printf("(Added %v day ago)\n\n", daysSince)
			} else {
				fmt.Printf("(Added %v days ago)\n\n",  daysSince)
			}
		}
			
	},
}

func init() {
	rootCmd.AddCommand(mydictCmd)
}
