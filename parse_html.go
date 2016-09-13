package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

func parseHTML(r io.Reader) {
	d := html.NewTokenizer(r)
	for {
		// token type
		tokenType := d.Next()
		//fmt.Println(tokenType)
		if tokenType == html.ErrorToken {
			return
		}
		token := d.Token()
		fmt.Println("Token is: ", token)
		switch tokenType {
		case html.StartTagToken:
			if token.Data == "title" {
				fmt.Println("This is title token !!")

				next := d.Next()
				fmt.Println("Next:", next)
				//token := d.Token()
				token1 := d.Token()
				fmt.Println("Suite Name: ", token1)
			}
		//	fmt.Println(token.Data)
		// <tag>
		// type Token struct {
		//     Type     TokenType
		//     DataAtom atom.Atom
		//     Data     string
		//     Attr     []Attribute
		// }
		//
		// type Attribute struct {
		//     Namespace, Key, Val string
		// }
		case html.TextToken: // text between start and end tag
			if token.Data == "h1" {
				//fmt.Println(token)
			}
		case html.EndTagToken: // </tag>
		case html.SelfClosingTagToken: // <tag/>

		}
	}
}

func main() {

	filename := "com.udacity.gradle.test.PersonTest.html"

	htmlFile, err := os.Open(filename)

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer htmlFile.Close()

	//	b, _ := ioutil.ReadAll(htmlFile)

	parseHTML(htmlFile)

}
