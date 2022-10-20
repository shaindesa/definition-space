# definition-space
## Introduction
`definition-space` is a simple CLI tool written in Go for keeping a local dictionary of words that you can look up and learn. 

Often the knowledge we gain from looking up words in the dictionary is ephemeral. `definition-space` addresses this problem. In `definition-space`, words can be 'learned' by calling `definition-space learn`. Learn gives you a reminder on the definitions of words you have recently added. This rote technique is useful for improving your retention and increasing your vocabulary.

`definition-space` is lightweight and simple to use. It requires an internet connection to add words to the dictionary, but does not require one when reading your local dictionary. Keep it open while reading a difficult book or article and call on it when the words get tough.

The tool queries the API hosted at https://api.dictionaryapi.dev/api/v2/entries/en/ and stores words in JSON format in `$HOME/.definition-space`.

## Usage
You will need a recent version of Go installed on your machine to build the binary file from the source code. 

Clone directory using git:

`git clone https://github.com/shaindesa/definition-space/`

`cd` into the `definition-space` directory and run

`go build .`

An executable binary file will be created. Move this file into your `$PATH` to run `definition-space` from your command line.

## Commands
### lookup
`definition-space lookup [word]`

Looks up the word given as an argument in the dictionary API. It will return all of the definitions one-by-one. This command commits nothing to local memory.

### add
`definition-space add [word]`

Returns definitions in the same way `lookup` does, but giving the user an additional option to add definitions to their local dictionary.

### check
`definition-space check [word]`

Queries the user's local dictionary and returns the definition stored for all matching words.

### remove
`definition-space remove [word]`

Removes word from user's local dictionary.

### mydict
`definition-space mydict`

Prints all words and their definitions in the user's local dictionary to the output.

### cleardict
`definition-space cleardict`

Clears all entries in user's local dictionary.

### learn
`definition-space learn`

Chooses up to five words in the user's local dictionary that have not been `learn`ed four times already. These words have their definitions printed to help the user with retaining their meanings. 



