package gitp

import (
	"context"
	"net/url"

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

func (g *Gitlab) ListRepos(ctx context.Context) ([]*provider.Repo, error) {
	projects, _, err := g.Projects.ListProjects(&gitlab.ListProjectsOptions{Owned: gitlab.Bool(true)})
	if err != nil {
		return nil, err
	}

	results := make([]*provider.Repo, 0)
	for _, proj := range projects {
		results = append(results, &provider.Repo{
			Name:      proj.Path,
			FullName:  proj.PathWithNamespace,
			Intro:     proj.Description,
			GitUrl:    proj.SSHURLToRepo,
			HttpUrl:   proj.HTTPURLToRepo,
			CreatedAt: timestamppb.New(proj.CreatedAt.Local()),
			UpdatedAt: timestamppb.New(proj.LastActivityAt.Local()),
		})
	}

	return results, nil
}

func (g *Gitlab) ListBranches(ctx context.Context, fullName string) ([]*provider.Branch, error) {
	results, _, err := g.Branches.ListBranches(url.PathEscape(fullName), &gitlab.ListBranchesOptions{})
	if err != nil {
		return nil, err
	}

	branches := make([]*provider.Branch, 0)
	for _, ret := range results {
		branches = append(branches, &provider.Branch{
			Name:      ret.Name,
			Default:   ret.Default,
			CreatedAt: timestamppb.New(ret.Commit.CreatedAt.Local()),
			UpdatedAt: timestamppb.New(ret.Commit.CommittedDate.Local()),
		})
	}

	return branches, nil
}

func (g *Gitlab) ListCommits(ctx context.Context, fullName, refName string) ([]*provider.Commit, error) {
	results, _, err := g.Commits.ListCommits(url.PathEscape(fullName), &gitlab.ListCommitsOptions{RefName: gitlab.String(refName)})
	if err != nil {
		return nil, err
	}

	commits := make([]*provider.Commit, 0)
	for _, ret := range results {
		commits = append(commits, &provider.Commit{
			CommitId:       ret.ID,
			ShortId:        ret.ShortID,
			Message:        ret.Message,
			CommitterName:  ret.CommitterName,
			CommitterEmail: ret.CommitterEmail,
			CreatedAt:      timestamppb.New(ret.CreatedAt.Local()),
			UpdatedAt:      timestamppb.New(ret.CommittedDate.Local()),
		})
	}

	return commits, nil
}
