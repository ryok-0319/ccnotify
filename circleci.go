package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// CircleCI represents information obtained from CircleCI
type CircleCI struct {
	Owner string
	PR    PullRequest
	Repo  string
}

// PullRequest represents GitHub pull request
type PullRequest struct {
	Revision string
	Number   int
}

// NewCircleCI constructs CircleCI
func NewCircleCI() (circleci *CircleCI, err error) {
	c := &CircleCI{
		Owner: os.Getenv("CIRCLE_PROJECT_USERNAME"),
		Repo:  os.Getenv("CIRCLE_PROJECT_REPONAME"),
		PR: PullRequest{
			Revision: os.Getenv("CIRCLE_SHA1"),
			Number:   0,
		},
	}

	pr := os.Getenv("CIRCLE_PULL_REQUEST")
	if pr == "" {
		return c, nil
	}
	re := regexp.MustCompile(`[1-9]\d*$`)
	c.PR.Number, err = strconv.Atoi(re.FindString(pr))
	if err != nil {
		return c, fmt.Errorf("%v: cannot get env", pr)
	}
	return c, nil
}
