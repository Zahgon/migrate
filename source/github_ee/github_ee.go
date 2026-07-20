package github_ee

import (
	"github.com/golang-migrate/migrate/v4/source"

	"github.com/google/go-github/v39/github"
)

func init() {
	source.Register("github-ee", &GithubEE{})
}

type GithubEE struct {
	source.Driver
}

func (g *GithubEE) Open(url string) (source.Driver, error) {
	_ = "STUB: not implemented"
	return *new(source.Driver), nil
}

func (g *GithubEE) createGithubClient(host, username, password string, verifyTLS bool) (*github.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseBool(val string, fallback bool) bool { _ = "STUB: not implemented"; return false }
