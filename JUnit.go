package junit2testrail

import "encoding/xml"

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
