package main

import (
	"fmt"
        "os"
	"encoding/xml"
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
  
        data := `
                <?xml version="1.0" encoding="UTF-8" ?>
                <testsuite tests="1" failures="0" name="org.hemani.javabrains.AppTest" time="0.004" errors="0" skipped="0">
			  <properties>
    				<property name="java.runtime.name" value="Java(TM) SE Runtime Environment"/>
    				<property name="sun.boot.library.path" value="/usr/java/jdk1.8.0_60/jre/lib/amd64"/>
    				<property name="java.vm.version" value="25.60-b23"/>
    				<property name="java.vm.vendor" value="Oracle Corporation"/>
    				<property name="maven.multiModuleProjectDirectory" value="/home/hemani/myapp/MavenTestApp"/>
			  </properties>
  			  <testcase classname="org.hemani.javabrains.AppTest" name="testApp" time="0.004"/>
		</testsuite>
        `
	/*
        err := xml.Unmarshal([]byte(data), &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
        */

        xmlFile, err := os.Open("test.xml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer xmlFile.Close()
	
	//var q Query
	xml.Unmarshal(xmlFile, &v)

        fmt.Printf("Test Result Struct: %v\n", v)

      	fmt.Printf("Name: %q\n", v.Name)
	fmt.Printf("Tests: %q\n", v.Tests)
	fmt.Printf("Failures: %q\n", v.Failures)
	fmt.Printf("Errors: %q\n", v.Errors)
	fmt.Printf("Skipped: %q\n", v.Skipped)
	fmt.Printf("Time: %q\n", v.Time)
	
}

