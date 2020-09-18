package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html/charset"
)

// Testsuites mirrors the structure of JUnit XML
type Testsuites struct {
	XMLName   xml.Name `xml:"testsuites"`
	Text      string   `xml:",chardata"`
	Testsuite struct {
		Text       string `xml:",chardata"`
		Tests      string `xml:"tests,attr"`
		Failures   string `xml:"failures,attr"`
		Time       string `xml:"time,attr"`
		Name       string `xml:"name,attr"`
		Properties struct {
			Text     string `xml:",chardata"`
			Property struct {
				Text  string `xml:",chardata"`
				Name  string `xml:"name,attr"`
				Value string `xml:"value,attr"`
			} `xml:"property"`
		} `xml:"properties"`
		Testcase []struct {
			Text      string `xml:",chardata"`
			Classname string `xml:"classname,attr"`
			Name      string `xml:"name,attr"`
			Time      string `xml:"time,attr"`
			Failure   struct {
				Text    string `xml:",chardata"`
				Message string `xml:"message,attr"`
				Type    string `xml:"type,attr"`
			} `xml:"failure"`
		} `xml:"testcase"`
	} `xml:"testsuite"`
} 


func main() {
        dec := xml.NewDecoder(os.Stdin)
        dec.CharsetReader = charset.NewReaderLabel
        dec.Strict = false

        var doc Testsuites
        if err := dec.Decode(&doc); err != nil {
                log.Fatal(err)
        }
        b, err := json.Marshal(doc)
        if err != nil {
                log.Fatal(err)
        }
        fmt.Println(string(b))
}