package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type TestResultXML struct {
	Result         string `xml:"result,attr"`
	TimeOut        string `xml:"timeout,attr"`
	Tests          string `xml:"completed"`
	Failures       string `xml:"failures"`
	Errors         string `xml:"errors"`
	Skipped        string `xml:"skipped"`
	FailureMessage string `xml:"failureMessage"`
}

func displayXMLResult(v TestResultXML) {

	fmt.Println("\nResults: \n")
	fmt.Printf("Result Code %q\n", v.Result)
	fmt.Printf("Tests Run: %q\n", v.Tests)
	fmt.Printf("Failures: %q\n", v.Failures)
	fmt.Printf("Errors: %q\n", v.Errors)
	fmt.Printf("Skipped: %q\n", v.Skipped)
	fmt.Printf("Time Out: %q\n", v.TimeOut)
	fmt.Printf("Failure Message: %q\n", v.FailureMessage)
}

func parseXML(filename string) (TestResultXML, error) {

	v := TestResultXML{}

	xmlFile, err := os.Open(filename)

	if err != nil {
		fmt.Println("Error opening file:", err)
		return v, err
	}

	defer xmlFile.Close()

	b, _ := ioutil.ReadAll(xmlFile)

	xml.Unmarshal(b, &v)

	return v, nil
}

func main() {

	sampleFile := "failsafe-summary.xml"
	result, _ := parseXML(sampleFile)
	displayXMLResult(result)

}
