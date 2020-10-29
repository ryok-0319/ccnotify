package github

import (
	"context"

	"github.com/google/go-github/v32/github"
)

// API is GitHub API interface
type API interface {
	IssuesCreateComment(ctx context.Context, number int, comment *github.IssueComment) (*github.IssueComment, *github.Response, error)
	RepositoriesCreateComment(ctx context.Context, sha string, comment *github.RepositoryComment) (*github.RepositoryComment, *github.Response, error)
}

// GitHub represents the attribute information necessary for requesting GitHub API
type GitHub struct {
	*github.Client
	owner, repo string
}

// IssuesCreateComment is a wrapper of https://godoc.org/github.com/google/go-github/github#IssuesService.CreateComment
func (g *GitHub) IssuesCreateComment(ctx context.Context, number int, comment *github.IssueComment) (*github.IssueComment, *github.Response, error) {
	return g.Client.Issues.CreateComment(ctx, g.owner, g.repo, number, comment)
}

// RepositoriesCreateComment is a wrapper of https://godoc.org/github.com/google/go-github/github#RepositoriesService.CreateComment
func (g *GitHub) RepositoriesCreateComment(ctx context.Context, sha string, comment *github.RepositoryComment) (*github.RepositoryComment, *github.Response, error) {
	return g.Client.Repositories.CreateComment(ctx, g.owner, g.repo, sha, comment)
}
