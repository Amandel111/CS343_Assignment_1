package main

import (
		"fmt"
		"os"
		"log"
		"io/ioutil"
		"strings"
	)

var numWords = make(map[string]int)
func createDict (word string){
	if word != ""{
		lowercaseWord := strings.ToLower(word)
		allLetterWord := ""
		for i := 0; i < len(lowercaseWord); i++{
			// ascii for lowercase characters
			if lowercaseWord[i] > 96 && lowercaseWord[i] < 123{
				allLetterWord += string(lowercaseWord[i])
				//fmt.Print(" ", string(lowercaseWord[i]), " is a character.")
			}
		}
		//fmt.Print("only letters ", allLetterWord)
		//fmt.Print("lowercaseWord", lowercaseWord)
		numWords[allLetterWord] += 1
		//fmt.Print(lowercaseWord, numWords[lowercaseWord])
	}
}
func single_threaded(files []string) {
	// TODO: Your single-threaded implementation
	//for each element in files, if the word is separated by a space/pubctuation : if it exists, increment dictionary, othewise add to dictionary
	//need to loop thorugh each string in files
	//for each string, figure out how to access txt file connected to the string (as a filepath)
	//numWords := make(map[string]int)
	fmt.Print(files[0])
	//content, err:= os.ReadFile(files[0])
	for i := 0; i < len(files); i++{
		content, err:= os.ReadFile(files[i])
		if err != nil{
			panic(err)
		}
		stringContent := string(content)
		wordList := strings.Split(stringContent, " ")
		//fmt.Printf(string(numWords))
		for j := 0; j < len(wordList); j++{
			createDict(wordList[j])
		}
	err1 := ioutil.WriteFile("output/single.txt", content, 0777)
	if err1 != nil{
		log.Fatalf("%v", err)
		}
	}
	fmt.Print(numWords)
	//fmt.Print(output);
	//fmt.Print(string(content));
	//return content
	
}

func multi_threaded(files []string) {
	// TODO: Your multi-threaded implementation
}


func main() {
	// TODO: add argument processing and run both single-threaded and multi-threaded functions
	single_threaded([]string{"input/book.txt", "input/book2.txt"});
}