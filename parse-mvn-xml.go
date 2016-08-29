package main

import (
	"encoding/xml"
        "fmt"
        "io/ioutil"
        "os" 
)

func main() {

    type TestResult struct {
	Name string `xml:"name,attr"`
	Tests string `xml:"tests,attr"`
	Failures string `xml:"failures,attr"`
	Errors string `xml:"errors,attr"`
	Skipped string `xml:"skipped,attr"`
	Time string `xml:"time,attr"`
    }

    v := TestResult{}

    xmlFile, err := os.Open("test.xml")

    if err != nil {
	fmt.Println("Error opening file:", err)
	return
    }

    defer xmlFile.Close()

    b, _ := ioutil.ReadAll(xmlFile)
 	
    xml.Unmarshal(b, &v)
 
    fmt.Printf("Name: %q\n", v.Name)
    fmt.Printf("Tests Run: %q\n", v.Tests)
    fmt.Printf("Failures: %q\n", v.Failures)
    fmt.Printf("Errors: %q\n", v.Errors)
    fmt.Printf("Skipped: %q\n", v.Skipped)
    fmt.Printf("Time Elapsed: %q\n", v.Time)
	
}

