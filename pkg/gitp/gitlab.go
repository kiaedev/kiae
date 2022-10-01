package gitp

import (
	"context"

	"github.com/kiaedev/kiae/api/provider"
	"github.com/xanzy/go-gitlab"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Gitlab struct {
	*gitlab.Client
}

func NewGitlab(token string) Provider {
	client, _ := gitlab.NewOAuthClient(token)
	return &Gitlab{
		Client: client,
	}
}

func (g *Gitlab) List(ctx context.Context) ([]*provider.Repo, error) {
	projects, _, err := g.Projects.ListProjects(&gitlab.ListProjectsOptions{Owned: gitlab.Bool(true)})
	if err != nil {
		return nil, err
	}

	results := make([]*provider.Repo, 0)
	for _, proj := range projects {
		results = append(results, &provider.Repo{
			Name:      proj.Name,
			Intro:     proj.Description,
			GitUrl:    proj.SSHURLToRepo,
			HttpUrl:   proj.HTTPURLToRepo,
			CreatedAt: timestamppb.New(proj.CreatedAt.Local()),
			UpdatedAt: timestamppb.New(proj.LastActivityAt.Local()),
		})
	}

	return results, nil
}
