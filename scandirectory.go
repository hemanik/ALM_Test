package main

import (
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

		if strings.Contains(file.Name(), ".xml") {
			result, _ := parseXML(file.Name())
			displayXMLResult(result)
		}

		if strings.Contains(file.Name(), ".txt") {
			result, _ := parseTxt(file.Name())
			displayTxtResult(result)
		}
	}
}
