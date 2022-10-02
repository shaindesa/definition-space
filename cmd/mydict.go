/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type WordData struct{
	Value string `json:"value"`
	POS string `json:"partofspeech"`
	Definition string `json:"definition"`
}


func CheckDir() {
	homedir := os.Getenv("HOME")
	dir := homedir + "/.definition-space/"
	_, err := os.Stat(dir)
	if os.IsNotExist(err){
		fmt.Println("Local Dictionary not found. Creating a new directory for your dictionary.")
		os.Chdir(homedir)
		os.Mkdir(".definition-space", 0722)
		return
	}
}

func readJSON(fileName string) ([]WordData, error){
	file, err := os.Open(fileName)
	if err != nil{
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	outarr := []WordData{}

	decoder.Decode(&outarr)

	return outarr, nil
}

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

		words, err := readJSON(file)
		if err != nil{
			fmt.Println("You don't have a local dictionary yet. Get started by calling `definition-space add [word]`")
			return
		}

		for key, val := range words{
			fmt.Printf("%v\t'%v'\n", key, val.Value)
			fmt.Printf("%v\n", val.POS)
			fmt.Printf("%v\n\n", val.Definition)

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
