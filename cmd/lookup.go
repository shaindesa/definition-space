package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var lookupCmd = &cobra.Command{
	Use:   "lookup",
	Short: "Looks up definitions of words. Does not commit anything to memory.",
	Long: `
		definition-space lookup
	Use this command alongside a single real English word to lookup the word's definition(s) in the dictionary API.

	Usage: 
	definition-space lookup [word]

	`,

	Args: cobra.ExactArgs(1),


	Run: func(cmd *cobra.Command, args []string) {

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
				fmt.Println("(Press ENTER to continue listing definitions, or type anything else to exit)")
				fmt.Scanln(&response)
				if response != ""{
					break stopreading
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(lookupCmd)
}
