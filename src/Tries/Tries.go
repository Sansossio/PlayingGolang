package main;

import (
	"fmt"
	"strings"
	"encoding/json"
);
// Const
const total = 26
// Trie struct
type trie struct {
	letters [total]*trie
	isEndWord bool
}
type jsonStruct struct {
	Letter string
	Childs []jsonStruct
}
// Parser word
func parserWord(word string) (int, string) {
	return len(word), strings.ToLower(word)
}
// Byte Code
func getWordCode(word string, position int) byte {
	return word[position] - 'a';
}
// New instance
func getNewTrie() *trie {
	newTrie := trie{ isEndWord: false }
	return &newTrie
}
// Insert string
func insert(trieExplorer *trie, word string) {
	// Properties
	length, myWord := parserWord(word)
	// Bucle
	for i := 0; i < length; i++ {
		// Get code
		index := getWordCode(myWord, i)
		// Comprobation
		if trieExplorer.letters[index] == nil {
			trieExplorer.letters[index] = getNewTrie()
		}
		trieExplorer = trieExplorer.letters[index]
	}
	trieExplorer.isEndWord = true
}
// Insert array
func insertArray(trieExplorer *trie, word []string) {
	// Properties
	length := len(word);
	// Loop
	for i := 0; i < length; i++ {
		insert(trieExplorer, word[i])
	}
}
// Search
func search(trieExplorer *trie, word string) bool {
	// Properties
	length, myWord := parserWord(word)
	// Bucle
	for i := 0; i < length; i++ {
		// Get code
		index := getWordCode(myWord, i)
		// Comprobation
		if trieExplorer.letters[index] == nil {
			return false
		}
		trieExplorer = trieExplorer.letters[index]
	}
	return trieExplorer.isEndWord
}
// Show trie
func showTrie(trieExplorer *trie) []jsonStruct {
	// Count
	count := 0;
	var indexs []int;
	for i := 0; i < total; i++ {
		if (trieExplorer.letters[i] != nil) {
			// Slice
			indexs = append(indexs, i)
			// Increase
			count++
		}
	}
	// Response
	response := make([]jsonStruct, count)
	// For loop
	for i := 0; i < count; i++ {
		// Index
		myIndex := indexs[i]
		// Comprobe
		if (trieExplorer.letters[myIndex] != nil) {
			// Get childs
			childs := showTrie(trieExplorer.letters[myIndex])
			// Response
			myLetter := string(myIndex + 'a')
			// Create json struct
			myStruct := jsonStruct{ Letter: myLetter, Childs: childs }
			// Create json
			if (myLetter != "") {
				response[i] = myStruct
			}
		}
	}
	// Response
	return response
}
// Main
func main() {
	// Instance
	myTrie := getNewTrie();
	// Insert
	allWords := []string{
		"hello",
		"world",
		"word",
		"want",
	}
	insertArray(myTrie, allWords)
	// Search
	result := search(myTrie, "world")
	fmt.Println(result)
	// Show trie structure
	show := showTrie(myTrie)
	print, _ := json.Marshal(show)
	fmt.Println(string(print))
}