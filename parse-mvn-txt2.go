package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

type Failure_Details struct {
	TestCase string
	Type     string
	Message  string
	Location []string
}

type TestResult struct {
	Name     string
	Tests    string
	Failures string
	Errors   string
	Skipped  string
	Time     string
	Failure_Details
}

func parseTxt(filename string) (TestResult, error) {

	v := TestResult{}

	f, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error opening file:", err)
		return v, err
	}

	name := regexp.MustCompile("Test set: ([a-zA-Z0-9_.]+)")
	v.Name = name.FindStringSubmatch(string(f))[1]

	tests := regexp.MustCompile("Tests run: ([a-zA-Z0-9_.]+)")
	v.Tests = tests.FindStringSubmatch(string(f))[1]

	failures := regexp.MustCompile("Failures: ([a-zA-Z0-9_.]+)")
	v.Failures = failures.FindStringSubmatch(string(f))[1]

	errors := regexp.MustCompile("Errors: ([a-zA-Z0-9_.]+)")
	v.Errors = errors.FindStringSubmatch(string(f))[1]

	skipped := regexp.MustCompile("Skipped: ([a-zA-Z0-9_.]+)")
	v.Skipped = skipped.FindStringSubmatch(string(f))[1]

	time := regexp.MustCompile("Time elapsed: ([a-zA-Z0-9_.]+)")
	v.Time = time.FindStringSubmatch(string(f))[1]

	fmt.Println("\nResults: \n")

	fmt.Printf("Test Suite: %q\n", v.Name)
	fmt.Printf("Tests Run: %q\n", v.Tests)
	fmt.Printf("Failures: %q\n", v.Failures)
	fmt.Printf("Errors: %q\n", v.Errors)
	fmt.Printf("Skipped: %q\n", v.Skipped)
	fmt.Printf("Time Elapsed: %q\n", v.Time)

	if v.Failures != "0" {
		failedTestCase := regexp.MustCompile(`.*\(.*\)`)
		v.Failure_Details.TestCase = failedTestCase.FindString(string(f))

		failureType := regexp.MustCompile(`(.*Error):(.*)`)
		v.Failure_Details.Type = failureType.FindStringSubmatch(string(f))[1]
		v.Failure_Details.Message = failureType.FindStringSubmatch(string(f))[2]

		failureLocation := regexp.MustCompile("at.*")
		v.Failure_Details.Location = failureLocation.FindAllString(string(f), -1)

		fmt.Printf("\nFailed Test Case: %q\n", v.Failure_Details.TestCase)
		fmt.Printf("Failure Type: %q\n", v.Failure_Details.Type)
		fmt.Printf("Failure Message: %q\n", v.Failure_Details.Message)
		fmt.Printf("Failure Details: %q\n", v.Failure_Details.Location)
	}

	return v, nil
}

func main() {
	sampleFile := "test.txt"
        //parseTxt(sampleFile)
	result, _ := parseTxt(sampleFile)
        fmt.Println(result)
}
