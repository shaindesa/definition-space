/*
Copyright © 2022 Shain D shaindesa@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "definition-space",
	Short: "Lookup and learn words using a dictionary API.",
	Long: `definition-space allows you to lookup words from your CLI and add them to a local dictionary.
	The dictionary is initialised in $HOME/.definition-space/dictionary.json when you first call 'definition-space add [word]'
	The API used is https://api.dictionaryapi.dev/api/v2/entries/en/.
	After adding words to your dictionary, you can check a word with 'check', print the whole thing with 'mydict', learn words with 'learn', remove words with 'remove', and clear everything with 'cleardict'.


	`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


