[![Gitpod ready-to-code](https://img.shields.io/badge/Gitpod-ready--to--code-blue?logo=gitpod)](https://gitpod.io/#https://github.com/monch1962/junit2testrail)
[![Build Status](https://dev.azure.com/monch1962/monch1962/_apis/build/status/monch1962.junit2testrail?branchName=master)](https://dev.azure.com/monch1962/monch1962/_build/latest?definitionId=12&branchName=master)

# junit2testrail
Process JUnit test results into TestRail.

junit2testrail will read JUnit XML results from `stdin` and post those results to the nominated TestRail server. TestRail-specific configuration is supplied via environment variables:
- `TESTRAIL_SERVER` is the full URL of the TestRail instance
- `USERNAME` is the username of the user that will be used to post updates
- `PASSWORD` is the password of that user
- `PROJECT_NAME` is the project name on the TestRail instance to be updating
- `SUITE_NAME` is the suite name within the TestRail project to be updated

## Example usage

`$ cat junit-sample.xml | TESTRAIL_SERVER=... USERNAME=... PASSWORD=... PROJECT_NAME="My Test Project" SUITE_NAME="Master" go run JUnit.go`

or

`$ go build`

`$ cat junit-sample.xml | TESTRAIL_SERVER=... USERNAME=... PASSWORD=... PROJECT_NAME="My Test Project" SUITE_NAME="Master" ./junit2testrail`

or

`$ docker build . -t junit2testrail:latest`

`$ docker run -e ...`

## TODO
- decide how to handle JUnit test cases that don't exist within the nominated project/suite
  - ignore them?
  - error out?
  - create the test case within the project/suite and update it?
- couldn't get batches of updates working due to some strangeness in the TestRail API - at the moment I'm doing one update per testcase, which isn't ideal
- write some `--help` documentation
- create some detailed docs
- consider adding support for Zephyr, XRay, ALM etc. (shouldn't be hard, just don't need it yet)
- work out how to drive this from Kubernetes-hosted infra, as part of a completely automated testing capability (e.g. test execution driven by Gitops)
- consider whether this should be installable as a Knative FaaS
- build in CI to build/test & deploy to a Docker registry (docker.io, GCP, AWS, Azure, ...)
