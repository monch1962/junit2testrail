package main

import (
	"os"
	"testing"

	"github.com/educlos/testrail"
)

  func TestRailAccess(t *testing.T){

    username := "monch1962@gmail.com"
	password := os.Getenv("PASSWORD") // sent to me in email...
	
	var result testrail.SendableResult
	result.StatusID = 2

    client := testrail.NewClient("https://monch1962a.testrail.com", username, password)

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