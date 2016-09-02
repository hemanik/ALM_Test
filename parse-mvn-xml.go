package main

import (
	"encoding/xml"
        "fmt"
        "io/ioutil"
        "os" 
)

func parseXML(filename string) {

    type Failure struct {
  
        Message string `xml:"message,attr"`
        Type string `xml:"type,attr"` 
        Details string `xml:",chardata"` 
    } 

    type TestCase struct {
        Name string `xml:"name,attr"`
        F Failure `xml:"failure"` 
    }  

    type TestResult struct {
	TestSuite string `xml:"name,attr"`
	Tests string `xml:"tests,attr"`
	Failures string `xml:"failures,attr"`
	Errors string `xml:"errors,attr"`
	Skipped string `xml:"skipped,attr"`
	Time string `xml:"time,attr"`
        T []TestCase `xml:"testcase"` 
    }

    v := TestResult{}

    xmlFile, err := os.Open(filename)

    if err != nil {
	fmt.Println("Error opening file:", err)
	return
    }

    defer xmlFile.Close()

    b, _ := ioutil.ReadAll(xmlFile)
 	
    xml.Unmarshal(b, &v)
  
    fmt.Println("\nResults: \n")  
    fmt.Printf("Test Suite %q\n", v.TestSuite)
    fmt.Printf("Tests Run: %q\n", v.Tests)
    fmt.Printf("Failures: %q\n", v.Failures)
    fmt.Printf("Errors: %q\n", v.Errors)
    fmt.Printf("Skipped: %q\n", v.Skipped)
    fmt.Printf("Time Elapsed: %q\n", v.Time)
    
    for i, test := range v.T {
   
       fmt.Printf("\nTest Case %d: %q\n", i+1, test.Name)
       if v.Failures != "0" {
       	   fmt.Printf("Failure Message: %q\n", test.F.Message)
           fmt.Printf("Failure Type: %q\n", test.F.Type)
           fmt.Printf("Failure Details: %q\n", test.F.Details)
       }
    }	
}

func main() {
    sampleFile := "failure_test.xml"
    parseXML(sampleFile)
}
