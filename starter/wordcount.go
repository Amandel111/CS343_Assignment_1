package main

import (
		"fmt"
		"os"
		//"log"
		"io"
		"strings"
		"regexp"
		"strconv"
		//"sync"
	)

//global variables
var numWordSingle = make(map[string]int)
var numWordDouble = make(map[string]int)
func createDict (wordList []string, dict map[string] int){
	for j := 0; j < len(wordList); j++{
		if wordList[j] != ""{
			lowercaseWord := strings.ToLower(wordList[j])
			// allLetterWord := ""
			// for i := 0; i < len(lowercaseWord); i++{
			// 	// ascii for lowercase characters
			// 	// if lowercaseWord[i] > 96 && lowercaseWord[i] < 123{
			// 		allLetterWord += string(lowercaseWord[i])
			// 		//fmt.Print(" ", string(lowercaseWord[i]), " is a character.")
			// 	// }
			// }
			dict[lowercaseWord] += 1
		}
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
		createDict(wordList, numWordSingle)
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


func read_file_chunk(chunkSize int64, startByte int64, filePath string){
	// responsible for reading chunk of file
	fmt.Print("reading chunk")
	content, err:= os.ReadFile(filePath)
		if err != nil{
			panic(err)
		}
		stringContent := string(content)
	reader := strings.NewReader(stringContent) 
	r := io.NewSectionReader(reader, startByte, chunkSize)
	buf := make([]byte, chunkSize)
	n, err := r.Read(buf) 
    if err != nil { 
        panic(err) 
    } 
	fmt.Printf("Content in buffer: %s\n", string(buf)) 
	fmt.Printf("n: %v\n", n) 
}

// func read_list_chunk(wordListSection [] string, dict map[string] int){
// 	for i:= 0; i<len(wordListSection); i++{
// 	if wordListSection[i] != ""{
// 		lowercaseWord := strings.ToLower(wordListSection)
// 		sync.Mutex.Lock()
// 		dict[lowercaseWord] += 1
// 		sync.Mutex.Unlock()
// 	}
// 	fmt.Printf("List Chunk:", wordListSection)
// 	}
// }

func multi_threaded(files []string) {
	// TODO: Your multi-threaded implementation
	const NUM_THREADS = 3;
		//read from files and pass to diciotnary creator 
		for i := 0; i < len(files); i++{
			content, err:= os.ReadFile(files[i])
			if err != nil{
				panic(err)
			}
			stringContent := string(content)
			re1 := regexp.MustCompile(`\p{P}|[^\S+]`)
		
			wordList := re1.Split(stringContent, -1)
			lenWordList := len(wordList)

			for i:= 0; i < lenWordList; i+=lenWordList/NUM_THREADS{
				// endSplice := (i+lenWordList/NUM_THREADS)
				if i+lenWordList/NUM_THREADS > lenWordList{
					fmt.Printf("hello")
					// endSplice = lenWordList
				}
				//read_list_chunk(wordList[i:endSplice])
			}
			// fmt.Printf("wordList W", wordList)
			// for j := 0; j < len(wordList); j++{
			// 	createDict(wordList[j], numWordSingle)
			// }
		}
	//dictionary for words as a shared resource, will have mutual exclusion via lock/unlock
	//we have the number of threads we want, NUM_THREADS
	//we loop thorugh file
	//we have a file, split it into NUM_THREADS number of files/strings that we pass
	//we call the goroutine thread Num_THREADS number of times, this fills in the diciotnary
	// for i := 0; i < len(files); i++{
	// 	// startByte is where the thread should start reading from
	// 	startByte := int64(0)
	// 	//divide file based on length
	// 	file, err := os.Open( files[i]) 
	// if err != nil {
    // 	log.Fatal(err)
	// }
	// fi, err := file.Stat()
	// if err != nil {
    // 	log.Fatal(err)
	// }
	// fmt.Println( "file size" , fi.Size() )
	// sizeOfFileChunk := fi.Size() / NUM_THREADS 
	// fmt.Print("size of chunk", sizeOfFileChunk)
	// NOTE: DO WE NEED TO ADD A DOUBLE FOR LOOP TO LOOP THROUGH THE ENTIRE FILE
	// if remaining bytes of the file is smaller than file chunk edge case
	// if fi.Size() <= int64(startByte) + sizeOfFileChunk{
	// 	fmt.Print("remaining bytes less than size")
	// 	sizeOfFileChunk = fi.Size() - startByte
	// }
	// checks if its at the end of the file
	// if startByte >= fi.Size(){
	// 	fmt.Print("end of search")
	// }else{
	// 	fmt.Print("else statement")
	// 	read_file_chunk(sizeOfFileChunk, startByte, files[i])
	// 	startByte = startByte + sizeOfFileChunk
		// reader := strings.NewReader(files[i]) 
		// r := io.NewSectionReader(reader, startByte, sizeOfFileChunk)
		// buf := make([]byte, sizeOfFileChunk)
	// }
	// check if startByte > sizeOfFileChunk
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
//}

}


func main() {
	// TODO: add argument processing and run both single-threaded and multi-threaded functions
	single_threaded([]string{"input/book.txt", "input/book2.txt"})
	//multi_threaded([]string{"input/book.txt", "input/book2.txt"})
}