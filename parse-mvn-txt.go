package main

import (
     "bufio"
     "regexp"
     "fmt"
     "os"
     "strings"
)

func parse-txt1() {

    type TestResult struct {
          Name string 
          Tests string 
          Failures string 
          Errors string 
          Skipped string 
          Time string 
     }

     //v := TestResult{}

     f, _ := os.Open("test.txt")
     scanner := bufio.NewScanner(f)
     
     for scanner.Scan() {
          line := scanner.Text()
         
          pattern := `\w+(\s\w+)?: [^,\s]*`
          tmp:= regexp.MustCompile(pattern)
          str := tmp.FindAllString(line, -1)
       
          values := []string{"Test set", "Tests run" , "Failures", "Errors", "Skipped", "Time elapsed"}

          for _, i := range str {
                          
               x := strings.Split(i, ":")
               
               for  _, b := range values { 
               
                    if x[0] == b {
                        //Add to Struct
                        fmt.Print("\n", x[1])
                        break 
                    }
               }
          }    
     }
}
	
