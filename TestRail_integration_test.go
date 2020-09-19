// +build integration

package main

import (
	"os"
	"testing"

	"github.com/educlos/testrail"
)

  func TestRailAccess(t *testing.T){

	testrailServer := os.Getenv("TESTRAIL_SERVER") // e.g. https://id.testrail.io
    username := os.Getenv("USER")
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