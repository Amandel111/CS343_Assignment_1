package main

import (
		"fmt"
		"os"
		"log"
		//"io/ioutil"
		"strings"
		"regexp"
		"strconv"
	)

//global variables
var numWordSingle = make(map[string]int)
var numWordDouble = make(map[string]int)
func createDict (word string, dict map[string] int){
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
		dict[allLetterWord] += 1
	}
}
func single_threaded(files []string) {

	//read from files and pass to diciotnary creator 
	for i := 0; i < len(files); i++{
		content, err:= os.ReadFile(files[i])
		if err != nil{
			panic(err)
		}
		stringContent := string(content)
		re1 := regexp.MustCompile(`\p{P}|[^\S+]`)
	
		wordList := re1.Split(stringContent, -1)
		fmt.Printf("wordList W", wordList)
		for j := 0; j < len(wordList); j++{
			createDict(wordList[j], numWordSingle)
		}
	}
	write_to_file("output/single.txt", numWordSingle)


}

func write_to_file(filepath string, dict map[string] int){
	file, err := os.Create(filepath)
	if err != nil {
		fmt.Println(err)
		return
	}

	for key, value := range dict{
		s := key + ": " + strconv.FormatInt(int64(value), 10) + "\n"
		file.WriteString(s)
	}
	
}



func multi_threaded(files []string) {
	// TODO: Your multi-threaded implementation
	const NUM_THREADS = 2;
	//dictionary for words as a shared resource, will have mutual exclusion via lock/unlock
	//we have the number of threads we want, NUM_THREADS
	//we loop thorugh file
	//we have a file, split it into NUM_THREADS number of files/strings that we pass
	//we call the goroutine thread Num_THREADS number of times, this fills in the diciotnary
	for i := 0; i < len(files); i++{

		//divide file based on length
		file, err := os.Open( files[i]) 
	if err != nil {
    	log.Fatal(err)
	}
	fi, err := file.Stat()
	if err != nil {
    	log.Fatal(err)
	}
	fmt.Println( "file size" , fi.Size() )
	sizeOFileChunk := fi.Size() / NUM_THREADS 
	fmt.Print("size of chunk", sizeOFileChunk)
		/*content, err:= os.ReadFile(files[i])
		if err != nil{
			panic(err)
		}
		stringContent := string(content)
		re1 := regexp.MustCompile(`\p{P}|[^\S+]`)
	
		wordList := re1.Split(stringContent, -1)
		fmt.Print(wordList)
	}
*/
}
}


func main() {
	// TODO: add argument processing and run both single-threaded and multi-threaded functions
	single_threaded([]string{"input/book.txt", "input/book2.txt"})
	multi_threaded([]string{"input/book.txt", "input/book2.txt"})
}