/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	Args: cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		CheckDir()

		w, defcount, err := GetDefs(args[0])
		if err != nil{
			log.Fatal(err)
		}
		//loop through the definitions, print them out
		count := 0

stopreading:
		for _, wordgroup := range w[0].WordGroup{
			for _, definition := range wordgroup.Definitions{

				count++
				fmt.Printf("Definition %v / %v\n", count, defcount)
				fmt.Println(wordgroup.PartOfSpeech)
				fmt.Println(definition.Val)
				if definition.Example != ""{
					fmt.Printf("Example: \"%v\"\n\n", definition.Example)
				} else{
					fmt.Printf("(Example not provided)\n\n")
				}

				// Ask user whether they want to continue
				var response string
				fmt.Printf("\tPress ENTER to continue listing definitions\n\tType `a` to add to local dictionary\n\tType anything else to exit\n\n")
				fmt.Scanln(&response)
				switch{
					case response == "":
						continue
					case response == "a":
						newword := WordData{
							Word: args[0],
							POS: wordgroup.PartOfSpeech,
							Example: definition.Example,
							Definition: definition.Val}

						err := AddWord(newword)
						if err != nil{
							log.Fatal("Error adding word to dictionary")
						}
						fmt.Printf("Added %v to your local dictionary! Use 'definition-space mydict' to view\n", args[0])
						break stopreading
					default:
						break stopreading
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
