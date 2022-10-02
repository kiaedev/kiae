package gitp

import (
	"context"
	"fmt"

	"github.com/kiaedev/kiae/api/provider"
)

type Provider interface {
	ListRepos(ctx context.Context) ([]*provider.Repo, error)
	ListBranches(ctx context.Context, fullName string) ([]*provider.Branch, error)
	ListCommits(ctx context.Context, fullName, refName string) ([]*provider.Commit, error)
}

type Constructor func(token string) Provider

var (
	repoProviders = map[string]Constructor{
		"github": NewGithub,
		"gitlab": NewGitlab,
	}
)

func Select(provider string, token string) (Provider, error) {
	constructor, ok := repoProviders[provider]
	if !ok {
		return nil, fmt.Errorf("provider %v not found", provider)
	}

	return constructor(token), nil
}
