package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v32/github"
)

// CommentService handles communication with the comment related
// methods of GitHub API
type CommentService service

// PostOptions specifies the optional parameters to post comments to a pull request
type PostOptions struct {
	Number   int
	Revision string
}

// Post posts comment
func (g *CommentService) Post(body string, opt PostOptions) error {
	if opt.Number != 0 {
		_, _, err := g.client.API.IssuesCreateComment(
			context.Background(),
			opt.Number,
			&github.IssueComment{Body: &body},
		)
		return err
	}
	if opt.Revision != "" {
		_, _, err := g.client.API.RepositoriesCreateComment(
			context.Background(),
			opt.Revision,
			&github.RepositoryComment{Body: &body},
		)
		return err
	}
	return fmt.Errorf("github.comment.post: Number or Revision is required")
}
