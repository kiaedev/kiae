package gitp

import (
	"context"
	"time"

	"github.com/kiaedev/kiae/api/provider"
	"github.com/ktrysmt/go-bitbucket"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Bitbucket struct {
	*bitbucket.Client
}

func NewBitbucket(token string) Provider {
	return &Bitbucket{
		Client: bitbucket.NewOAuthbearerToken(token),
	}
}

func (g *Bitbucket) List(ctx context.Context) ([]*provider.Repo, error) {
	repos, err := g.Repositories.ListForAccount(&bitbucket.RepositoriesOptions{})
	if err != nil {
		return nil, err
	}

	results := make([]*provider.Repo, 0)
	for _, repo := range repos.Items {
		results = append(results, &provider.Repo{
			Name:  repo.Name,
			Intro: repo.Description,
			// GitUrl:    repo.Links["self"],
			// HttpUrl:   repo.Links[],
			CreatedAt: timeFormat(repo.CreatedOn),
			UpdatedAt: timeFormat(repo.UpdatedOn),
		})
	}

	return results, nil
}

func timeFormat(value string) *timestamppb.Timestamp {
	t, _ := time.Parse(time.RFC3339, value)
	return timestamppb.New(t)
}
