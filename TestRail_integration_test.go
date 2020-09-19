// +build integration

package main

import (
	//"fmt"
	"log"
	"os"
	"testing"

	"github.com/educlos/testrail"
)

  func TestRailAccess(t *testing.T){

	testrailServer := os.Getenv("TESTRAIL_SERVER") // e.g. https://id.testrail.io
    username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD") // sent to me in email...
	
	var result testrail.SendableResult
	result.StatusID = 5

    client := testrail.NewClient(testrailServer, username, password)

    projectID := 1
    suiteID := 1
	cases, err := client.GetCases(projectID, suiteID)
	if err != nil {
		t.Fatalf("%v\n",err)
	}

    for _, c := range cases{
      t.Logf("Test ID:%d\n",c.ID)
	}
	
	res, err := client.AddResult(2, result)
	if err != nil {
		t.Fatalf("AddResult error: %v\n",err)
	}
	t.Logf("AddResult: %v\n", res)
  }

  func TestCaseAccess(t *testing.T) {
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	testrailServer := os.Getenv("TESTRAIL_SERVER") // e.g. https://id.testrail.io

    client := testrail.NewClient(testrailServer, username, password)

    projectID := 1
    suiteID := 1
	cases, err := client.GetCases(projectID, suiteID)
	if err != nil {
		log.Fatalf("Error reading test cases: %v\n", err)
	}

    for _, c := range cases{
	  log.Println(c.ID)
	  log.Printf("case: %v\n", c)
    }
  }

  func TestProjectAccess(t *testing.T) {
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	testrailServer := os.Getenv("TESTRAIL_SERVER") // e.g. https://id.testrail.io

    client := testrail.NewClient(testrailServer, username, password)

	projects, err := client.GetProjects()
	if err != nil {
		log.Fatalf("Error reading projects: %v\n", err)
	}

    for _, p := range projects{
	  log.Println(p.ID)
	  log.Printf("project: %v\n", p)
	  suites, err := client.GetSuites(p.ID)
	  if err != nil {
		  log.Fatalf("Error reading suites from project %v\n", p)
	  }
	  for _, s := range suites {
		  log.Printf("suite: %v\n", s)
		  log.Printf("suite ID: %v, %v\n", s.ID, s.Name)
	  }
	}
	
  }