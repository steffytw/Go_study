package main

import (
	"fmt"
	"os"
)

func main(){
	file,err := os.Open("test.txt")
	if err != nil{
		fmt.Println("File doesn't Exist")
	}else{
		fmt.Println(file.Name()+ " opened successfully!")
	}
}
