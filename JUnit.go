package main

import (
	//"encoding/json"
	"encoding/xml"
	"errors"
	//"fmt"
	"log"
	"os"

	//"strconv"
	"time"

	"github.com/educlos/testrail"
	//str2duration "github.com/xhit/go-str2duration/v2"
	"github.com/davecgh/go-spew/spew"
	"golang.org/x/net/html/charset"
)

// Testsuites struct below is autogenerated using the awesome https://www.onlinetool.io/xmltogo/ tool!
// Testsuites mirrors the structure of JUnit XML output
type Testsuites struct {
	XMLName   xml.Name `xml:"testsuites"`
	Text      string   `xml:",chardata"`
	Testsuite []struct {
		Text       string `xml:",chardata"`
		Name       string `xml:"name,attr"`
		Errors     string `xml:"errors,attr"`
		Tests      string `xml:"tests,attr"`
		Failures   string `xml:"failures,attr"`
		Time       string `xml:"time,attr"`
		Timestamp  string `xml:"timestamp,attr"`
		Skipped    string `xml:"skipped,attr"`
		Properties struct {
			Text     string `xml:",chardata"`
			Property []struct {
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
			} `xml:"failure"`
			Skipped *string `xml:"skipped,omitempty"`  // Note that this is a *string, as an empty value will give a "" while non-existent will give nil. This lets us distinguish a <skipped /> from a non-existent tag
		} `xml:"testcase"`
	} `xml:"testsuite"`
}

func readEnvVars() (string, string, string, string, string) {
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
	projectName := os.Getenv("PROJECT_NAME")
	if password == "" {
		log.Fatalln("Environment variable PROJECT_NAME not specified")
	}
	suiteName := os.Getenv("SUITE_NAME")
	if suiteName == "" {
		log.Fatalln("Environment variable SUITE_NAME not specified")
	}
	return testRailServer, username, password, projectName, suiteName
}

func readJunitXML(file *os.File) Testsuites {
	dec := xml.NewDecoder(file)
	dec.CharsetReader = charset.NewReaderLabel
	dec.Strict = false

	var doc Testsuites
	if err := dec.Decode(&doc); err != nil {
		log.Panicf(err.Error())
	}
	return doc
}

func logJunitDetail(tss Testsuites) {
	log.Printf("%+v\n", tss.Testsuite)
	spew.Dump(tss.Testsuite)
	for _, ts := range tss.Testsuite {
		log.Printf("Number of tests: %v\n", ts.Tests)
		log.Printf("Number of failed tests: %v\n", ts.Failures)
		for i, tc := range ts.Testcase {
			log.Printf("%v\n", ts)
			log.Printf("%v\n", tc)

			log.Printf("Testcase %d: %v\n", i, tc)
			log.Printf("Testcase %d name: %v\n", i, tc.Name)
			log.Printf("Testcase %d failure message: %v\n", i, tc.Failure.Text)
			log.Println("-------------------------------------------")
		}
	}
}

func processResultsToTestRail(j Testsuites, client *testrail.Client, projectID int, suiteID int) {
	logJunitDetail(j)
	now := time.Now().Format("2006-01-02 15:04:05")

	for _, ts := range j.Testsuite {
		for i, tc := range ts.Testcase {

			//duration,err := str2duration.ParseDuration(fmt.Sprintf("%ss",tc.Time))
			//if err != nil {
			//	log.Fatalf("Error converting %v to duration\n",tc.Time)
			//}

			tcName := tc.Name
			testcaseID, err := getTestCaseID(client, projectID, suiteID, tcName)
			if err != nil {
				log.Panicf("Couldn't find test case '%s'\n", tc.Name)
			}
			var tcStatus int
			if tc.Failure.Text != "" {
				tcStatus = testrail.StatusFailed
			} else if tc.Skipped != nil {
				//tcStatus = testrail.StatusUntested  //StatusUntested results in a failure when posting - library bug??
				tcStatus = testrail.StatusRetest
			} else {
				tcStatus = testrail.StatusPassed
			}

			tsr := testrail.SendableResult{
				//Elapsed: *testrail.TimespanFromDuration(duration),
				StatusID: tcStatus,
				Comment:  tc.Failure.Text,
				Version:  now,
				Defects:  "",
				//AssignedToID: 1,
			}

			log.Printf("tsr: %v\n", tsr)
			result, err := client.AddResultForCase(projectID, testcaseID, tsr)
			if err != nil {
				log.Panicf("Error adding results for test case %d: %v\n", i, err)
			}
			log.Printf("Success adding results for test case %d: %v\n", i, result)
		}
	}
}

func getProjectID(client *testrail.Client, projectName string) (int, error) {
	projects, err := client.GetProjects()
	if err != nil {
		log.Panicf("Error reading projects: %v\n", err)
	}

	for _, p := range projects {
		//log.Println(p.ID)
		//log.Printf("project: %v\n", p)
		//log.Printf("project name: %v\n", p.Name)
		if p.Name == projectName {
			log.Printf("Found project '%s' is id %d\n", projectName, p.ID)
			return p.ID, nil
		}
	}
	return 0, errors.New("Couldn't find project")
}

func getSuiteID(client *testrail.Client, projectID int, suiteName string) (int, error) {
	suites, err := client.GetSuites(projectID)
	if err != nil {
		log.Panicf("Error reading suites: %v\n", err)
	}

	for _, s := range suites {
		//log.Println(s.ID)
		//log.Printf("suite: %v\n", s)
		//log.Printf("suite name: %v\n", s.Name)
		if s.Name == suiteName {
			log.Printf("Found suite '%s' for project '%d' is id %d\n", suiteName, projectID, s.ID)
			return s.ID, nil
		}
	}
	return 0, errors.New("Couldn't find suite")
}

func getTestCaseID(client *testrail.Client, projectID int, suiteID int, testcaseName string) (int, error) {
	testcases, err := client.GetCases(projectID, suiteID)
	if err != nil {
		log.Panicf("Error reading testcases: %v\n", err)
	}

	for _, tc := range testcases {
		//log.Println(s.ID)
		//log.Printf("suite: %v\n", s)
		//log.Printf("suite name: %v\n", s.Name)
		if tc.Title == testcaseName {
			log.Printf("Found testcase '%s' for suite '%d', project '%d' is id %d\n", testcaseName, suiteID, projectID, tc.ID)
			return tc.ID, nil
		}
	}
	//return 0, errors.New("Couldn't find test case")
	newTestcase, err := addTestCase(client, suiteID, testcaseName)
	return newTestcase.ID, err
}

func addTestCase(client *testrail.Client, suiteID int, testcaseName string) (testrail.Case, error) {
	now := time.Now().Format("2006-01-02 15:04:05")
	newTestCase := testrail.SendableCase{
		Title: testcaseName,
		Date:  now,
	}
	tc, err := client.AddCase(suiteID, newTestCase)
	return tc, err
}

func main() {
	testrailServer, username, password, projectName, suiteName := readEnvVars()
	junitDoc := readJunitXML(os.Stdin)
	client := testrail.NewClient(testrailServer, username, password)
	projectID, err := getProjectID(client, projectName)
	if err != nil {
		log.Panicf("Couldn't find project named '%s'\n", projectName)
	}
	log.Printf("Project ID for '%s' is %d\n", projectName, projectID)

	suiteID, err := getSuiteID(client, projectID, suiteName)
	if err != nil {
		log.Panicf("Couldn't find suite named '%s' for project '%s'\n", suiteName, projectName)
	}

	log.Printf("Suitename '%s' has id %d\n", suiteName, suiteID)
	processResultsToTestRail(junitDoc, client, projectID, suiteID)
	//log.Printf("Results: %v\n", success)
}
