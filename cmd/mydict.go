/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// mydictCmd represents the mydict command
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
			if daysSince == 1{
				fmt.Printf("%v\t'%v'\t\tAdded %v day ago\n", key+1, val.Word, daysSince)
			} else {
				fmt.Printf("%v\t'%v'\t\tAdded %v days ago\n", key+1, val.Word, daysSince)
			}

			fmt.Printf("%v\n", val.POS)
			fmt.Printf("%v\n", val.Definition)
			if val.Example != ""{
				fmt.Printf("Example: \"%v\"\n", val.Example)
			}
			fmt.Println()
		}
			
	},
}

func init() {
	rootCmd.AddCommand(mydictCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mydictCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mydictCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
