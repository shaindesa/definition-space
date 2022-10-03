/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// cleardictCmd represents the cleardict command
var cleardictCmd = &cobra.Command{
	Use:   "cleardict",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cleardictCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cleardictCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
