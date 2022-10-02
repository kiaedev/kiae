package gitp

import (
	"context"

	"github.com/google/go-github/github"
	"github.com/kiaedev/kiae/api/provider"
	"golang.org/x/oauth2"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Github struct {
	*github.Client
}

func NewGithub(token string) Provider {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)

	return &Github{
		Client: github.NewClient(oauth2.NewClient(ctx, ts)),
	}
}

func (g *Github) List(ctx context.Context) ([]*provider.Repo, error) {
	repos, _, err := g.Repositories.List(ctx, "", &github.RepositoryListOptions{})
	if err != nil {
		return nil, err
	}

	results := make([]*provider.Repo, 0)
	for _, repo := range repos {
		results = append(results, &provider.Repo{
			Name:      repo.GetName(),
			FullName:  repo.GetFullName(),
			Intro:     repo.GetDescription(),
			GitUrl:    repo.GetGitURL(),
			HttpUrl:   repo.GetURL(),
			CreatedAt: timestamppb.New(repo.CreatedAt.Time),
			UpdatedAt: timestamppb.New(repo.UpdatedAt.Time),
		})
	}

	return results, nil
}
