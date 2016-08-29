package main

import (
     "bufio"
     "regexp"
     "fmt"
     "os"
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

     f, _ := os.Open("test.txt")
     scanner := bufio.NewScanner(f)
     
     //scanner.Split(bufio.ScanWords)
 
     for scanner.Scan() {
         line := scanner.Text()
         pattern := `\w+(\s\w+)?: [^,\s]*` 
         re := regexp.MustCompile(pattern)
         
       // fmt.Println(line)
       // re := regexp.MustCompile(`[A-Za-z]+(\s[A-Za-z]+)?:`)
       // re := regexp.MustCompile(`\w+(\s\w+)?: [^,\s]*`)
         fmt.Printf("%q\n", re.FindAllString(line, -1))
     }

     //fmt.Printf("%q\n", strings.Split("a,b,c", ","))
 
}
	