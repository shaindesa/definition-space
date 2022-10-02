package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

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

func ReadJSON(fileName string) ([]WordData, error){
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


func AddWord(newword WordData) error{

	filename := os.Getenv("HOME") + "/.definition-space/dictionary.json"

	var words []WordData
	words, err := ReadJSON(filename)
	if err != nil{
		fmt.Println("Cannot find data file. Creating new data file.")
		words = []WordData{}
	}

	words = append(words, newword)

	jsondata, err := json.Marshal(words)
	if err != nil{
		log.Fatal(err)
	}

	return WriteJSON(filename, jsondata)
}

func WriteJSON(filename string, jsondata []byte) error{
	err := os.WriteFile(filename, jsondata, 0644)
	if err != nil{
		return err
	}
	return nil
}
