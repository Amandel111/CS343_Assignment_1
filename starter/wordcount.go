package main

import (
		"fmt"
		"os"
	)

func single_threaded(files []string) {
	// TODO: Your single-threaded implementation
	//for each element in files, if the word is separated by a space/pubctuation : if it exists, increment dictionary, othewise add to dictionary
	//need to loop thorugh each string in files
	//for each string, figure out how to access txt file connected to the string (as a filepath)
	fmt.Print(files[0]);
	content, err:= os.ReadFile(files[0]);
	if err != nil{
		panic(err)
	}
	fmt.Print(string(content));
	//return content
	
}

func multi_threaded(files []string) {
	// TODO: Your multi-threaded implementation
}


func main() {
	// TODO: add argument processing and run both single-threaded and multi-threaded functions
	single_threaded([]string{"input/big.txt", "input/book2.txt"});
}