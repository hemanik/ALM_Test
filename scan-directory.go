package main

import (
        "fmt"
	"io/ioutil"
	"log"
        "strings"
)

func main() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
                if strings.Contains(file.Name(), ".xml"){
                     parseXML(file.Name())  
                }
                if strings.Contains(file.Name(), ".txt"){
                     parseTxt(file.Name())
                }   
	}
}
