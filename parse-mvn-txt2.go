package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

type FailureDetails struct {
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
	FailureDetails
}

func displayTxtResult(v TestResult) {

	fmt.Println("\nResults: \n")

	fmt.Printf("Test Suite: %q\n", v.Name)
	fmt.Printf("Tests Run: %q\n", v.Tests)
	fmt.Printf("Failures: %q\n", v.Failures)
	fmt.Printf("Errors: %q\n", v.Errors)
	fmt.Printf("Skipped: %q\n", v.Skipped)
	fmt.Printf("Time Elapsed: %q\n", v.Time)

	if v.Failures != "0" {

		fmt.Printf("\nFailed Test Case: %q\n", v.FailureDetails.TestCase)
		fmt.Printf("Failure Type: %q\n", v.FailureDetails.Type)
		fmt.Printf("Failure Message: %q\n", v.FailureDetails.Message)
		fmt.Printf("Failure Details: %q\n", v.FailureDetails.Location)
	}
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

	if v.Failures != "0" {

		failedTestCase := regexp.MustCompile(`.*\(.*\)`)
		v.FailureDetails.TestCase = failedTestCase.FindString(string(f))

		failureType := regexp.MustCompile(`(.*Error)(:.*)`)
		v.FailureDetails.Type = failureType.FindStringSubmatch(string(f))[1]
		v.FailureDetails.Message = failureType.FindStringSubmatch(string(f))[2]

		failureLocation := regexp.MustCompile("at.*")
		v.FailureDetails.Location = failureLocation.FindAllString(string(f), -1)
	}

	return v, nil
}

func main() {

	sampleFile := "failure_test2.txt"
	result, _ := parseTxt(sampleFile)
	displayTxtResult(result)
}
