package main

import (
	//"encoding/json"
	"encoding/xml"
	//"fmt"
	"log"
	"os"

	"github.com/educlos/testrail"
	str2duration "github.com/xhit/go-str2duration/v2"
	"golang.org/x/net/html/charset"
)

// Testsuites struct below is autogenerated using the awesome https://www.onlinetool.io/xmltogo/ tool!
// Testsuites mirrors the structure of JUnit XML output
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

func readEnvVars() (string, string, string) {
	testRailServer := os.Getenv("TESTRAIL_SERVER")
	if testRailServer == "" {
		log.Fatalln("Environment variable TESTRAIL_SERVER not specified")
	}
	username := os.Getenv("USERNAME")
	if username == "" {
		log.Fatalln("Environment variable USERNAME not specified")
	}	
	password := os.Getenv("PASSWORD")
	if password == "" {
		log.Fatalln("Environment variable PASSWORD not specified")
	}
	return testRailServer,username,password
}

func readJunitXML(file *os.File) Testsuites {
	dec := xml.NewDecoder(file)
	dec.CharsetReader = charset.NewReaderLabel
	dec.Strict = false

	var doc Testsuites
	if err := dec.Decode(&doc); err != nil {
		log.Fatal(err)
	}
	return doc
	//junitJSON, err := json.Marshal(doc)
    //if err != nil {
    //    log.Fatal(err)
	//}
	//return junitJSON
}

func logJunitDetail(j Testsuites) {
	for i,tc := range j.Testsuite.Testcase {
		log.Printf("%v\n",j)
		log.Printf("%v\n",tc)
		log.Printf("Number of tests: %v\n", j.Testsuite.Tests)
		log.Printf("Number of failed tests: %v\n", j.Testsuite.Failures)
		log.Printf("Testcase %d: %v\n", i, tc)
		log.Printf("Testcase %d name: %v\n", i, tc.Name)
		log.Printf("Testcase %d failure message: %v\n", i,tc.Failure.Text)
		log.Println("-------------------------------------------")	
	}
}

func processResultsToTestRail(j Testsuites, testRailServer string, username string, password string) {
	logJunitDetail(j)
	log.Printf("TESTRAIL_SERVER: %s, USERNAME: %s, PASSWORD: %s\n", testRailServer, username, password)
	var trSendableResults testrail.SendableResults
	for _,tc := range j.Testsuite.Testcase {
		//tcDuration := fmt.Sprintf("%sms",tc.Time)
		r := testrail.SendableResult{
			//Elapsed: testrail.TimespanFromDuration(str2duration.ParseDuration(tcDuration)),
			StatusID: testrail.StatusPassed,
			Comment: tc.Failure.Text,
		}
		tr := testrail.Results{
			TestID: 1,
			SendableResult: r,
		}
		trSendableResults.append(trSendableResults,tr)
	}
	client := testrail.NewClient("https://example.testrail.com", username, password)
	result,err := client.AddResults(1,trSendableResults)
	if err != nil {
		log.Fatalf("Error adding results to Testrail: %v", err)
	}
	log.Println(result)
}

func main() {
	testRailServer, username, password := readEnvVars()
	junitDoc := readJunitXML(os.Stdin)
	processResultsToTestRail(junitDoc, testRailServer, username, password)	
}