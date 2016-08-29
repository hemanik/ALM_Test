package main

import (
     "regexp"
     "fmt"
     "io/ioutil"
)

func main() {

    type TestResult struct {
          Name string 
          Tests string 
          Failures string 
          Errors string 
          Skipped string 
          Time string 
     }

     v := TestResult{}

     f, err := ioutil.ReadFile("test.txt")

     if err != nil {
        fmt.Println("Error opening file:", err)
        return
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

    
    fmt.Printf("Name: %q\n", v.Name)
    fmt.Printf("Tests Run: %q\n", v.Tests)
    fmt.Printf("Failures: %q\n", v.Failures)
    fmt.Printf("Errors: %q\n", v.Errors)
    fmt.Printf("Skipped: %q\n", v.Skipped)
    fmt.Printf("Time Elapsed: %q\n", v.Time) 

}
	
