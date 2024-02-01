package main

import (
		"fmt"
		"os"
		"log"
		"io"
		"strings"
		"regexp"
		"strconv"
		"sync"
	)

//global variables
var numWordSingle = make(map[string]int)
var numWordDouble = make(map[string]int)
func createDict (wordList []string, dict map[string] int){
	for j := 0; j < len(wordList); j++{
		if wordList[j] != ""{
			lowercaseWord := strings.ToLower(wordList[j])
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
		s := key + " " + strconv.FormatInt(int64(value), 10) + "\n"
		file.WriteString(s)
	}
	
}

func multi_write_to_file(filepath string, dict map[string] int, m *sync.Mutex){
	file, err := os.Create(filepath)
	if err != nil {
		fmt.Println(err)
		return
	}

	m.Lock()
	for key, value := range dict{
		s := key + " " + strconv.FormatInt(int64(value), 10) + "\n"
		file.WriteString(s)
	}
	m.Unlock()
	
}



func read_file_chunk(chunkSize int64, startByte int64, filePath string, wg *sync.WaitGroup, m *sync.Mutex){
	defer wg.Done()
	//read from full file the designated chunk of bytes into the buffer
	fileContent, err:= os.ReadFile(filePath)
		if err != nil{
			panic(err)
		}
	stringFileContent := string(fileContent)
	reader := strings.NewReader(stringFileContent) 
	r := io.NewSectionReader(reader, startByte, chunkSize)

	buf := make([]byte, chunkSize)
	n, err := r.Read(buf) 
    if err != nil { 
        panic(err) 
    } 
	fmt.Printf("n: %v\n", n) 

	//split content into a wordlist 
	fileChunkWords := string(buf)
	re1 := regexp.MustCompile(`\p{P}|[^\S+]`)
	wordList := re1.Split(fileChunkWords, -1)

	//add to dicitonary 
	for j := 0; j < len(wordList); j++{
		if wordList[j] != ""{
			lowercaseWord := strings.ToLower(wordList[j])
			m.Lock()
			numWordDouble[lowercaseWord] += 1
			m.Unlock()
		}
	}
}


func multi_threaded(files []string) {
	const NUM_THREADS = 3;
	var wg sync.WaitGroup
	var m sync.Mutex 

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
		sizeOfFile := fi.Size();
		sizeOfFileChunk := sizeOfFile / NUM_THREADS 

		//Loop through a file and call thread for each chunk
		for startByte := int64(0); startByte < sizeOfFile; startByte = startByte + sizeOfFileChunk{
			wg.Add(1)
			//if remaining bytes of the file is smaller than file chunk edge case
			if sizeOfFile<= int64(startByte) + sizeOfFileChunk{
				//if the remaining byte is smaller than the expected chunk size 
				//fmt.Print("remaining bytes less than size")
					sizeOfFileChunk = sizeOfFile - startByte
				}
			// checks if start byte at the end of the file
			if startByte >= sizeOfFile{
				//fmt.Print("end of search")
			}else{
				go read_file_chunk(sizeOfFileChunk, startByte, files[i], &wg, &m)
				}
			}
			wg.Wait()
		}
		multi_write_to_file("output/multi.txt", numWordDouble, &m);
}


func main() {
	// TODO: add argument processing and run both single-threaded and multi-threaded functions
	single_threaded([]string{"input/book.txt", "input/book2.txt", "input/big.txt"})
	multi_threaded([]string{"input/book.txt", "input/book2.txt", "input/big.txt"})
}