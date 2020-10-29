package main

import (
	"bufio"
	"log"
	"os"

	"github.com/ryok-0319/ccnotify/github"

	"github.com/urfave/cli"
)

const (
	name        = "ccnotify"
	description = "Notify from CircleCI to GitHub"

	version = "0.0.1"
)

func main() {
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "print version",
	}

	var token string

	app := &cli.App{
		Name:    name,
		Version: version,
		Usage:   description,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "token",
				Aliases:     []string{"t"},
				Usage:       "GitHub access token",
				Destination: &token,
				EnvVars:     []string{"GITHUB_ACCESS_TOKEN"},
			},
		},
		Action: func(c *cli.Context) error {
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			input := scanner.Text()

			circleci, err := NewCircleCI()
			if err != nil {
				return err
			}

			client, err := github.NewClient(github.Config{
				Token: token,
				Owner: circleci.Owner,
				Repo:  circleci.Repo,
				PR: github.PullRequest{
					Revision: circleci.PR.Revision,
					Number:   circleci.PR.Number,
					Title:    "Title",
					Message:  "Message",
				},
			})
			if err != nil {
				return err
			}

			err = client.Notify.Notify(input)
			if err != nil {
				return err
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
