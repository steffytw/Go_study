package main

import (
	"fmt"
	"os"
	"errors"
)

func main(){
	file1,err := os.Open("test.txt")
	if err != nil{
		fmt.Println("File doesn't Exist")
	}else{
		fmt.Println(file1.Name()+ " opened successfully!")
	}

	sampleErr1 := errors.New("New error occured")
	fmt.Println(sampleErr1)

	sampleErr2 := fmt.Errorf("Error occured in %s","db connection")// error with formatting of error messages
	fmt.Println(sampleErr2) 

	file2 ,_ := os.Open("sample.text") // underscore (‘_’) operator can be used to ignore the error
	fmt.Println(file2)
}
