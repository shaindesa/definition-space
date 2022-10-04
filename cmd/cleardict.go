package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var cleardictCmd = &cobra.Command{
	Use:   "cleardict",
	Short: "Clear all entries in your local dictionary.",
	Long: `Replaces the JSON in $HOME/.definition-space/dictionary.json with "[]", effectively clearing all entries from your local dictionary. Will ask user for confirmation.
	
	Usage: 'definition-space cleardict'
	
	`,
	Run: func(cmd *cobra.Command, args []string) {
		CheckDir()
		
		var confirmation string
		fmt.Println("Are you sure you want to clear your entire local dictionary? This action cannot be reversed (y/N)")
		fmt.Scanln(&confirmation)
		if confirmation != "Y" && confirmation != "y"{
			fmt.Println("Exiting. Dictionary NOT modified.")
			return
		}
		err := WriteJSON(os.Getenv("HOME") + "/.definition-space/dictionary.json", []byte("[]"))
		if err != nil{
			fmt.Println("There was an error in clearing the dictionary.")
			log.Fatal(err)
		}
		fmt.Println("Local dictionary has been cleared.")


	},
}

func init() {
	rootCmd.AddCommand(cleardictCmd)
}
