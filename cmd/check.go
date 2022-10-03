/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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

		if found == false{
			fmt.Println("Word not found in local dictionary")
		}
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
