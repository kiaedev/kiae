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
	results, _, err := g.Branches.ListBranches(fullName, &gitlab.ListBranchesOptions{})
	if err != nil {
		return nil, err
	}

	branches := make([]*provider.Branch, 0)
	for _, ret := range results {
		branches = append(branches, &provider.Branch{
			Name:   ret.Name,
			Commit: pCommitFromGl(ret.Commit),
		})
	}

	return branches, nil
}

func (g *Gitlab) ListTags(ctx context.Context, fullName string) ([]*provider.Tag, error) {
	results, _, err := g.Tags.ListTags(fullName, &gitlab.ListTagsOptions{})
	if err != nil {
		return nil, err
	}

	commits := make([]*provider.Tag, 0)
	for _, ret := range results {
		commits = append(commits, &provider.Tag{
			Name:   ret.Name,
			Commit: pCommitFromGl(ret.Commit),
		})
	}

	return commits, nil
}

func pCommitFromGl(commit *gitlab.Commit) *provider.Commit {
	return &provider.Commit{
		Sha1:           commit.ID,
		ShortId:        commit.ShortID,
		Message:        commit.Message,
		CommitterName:  commit.CommitterName,
		CommitterEmail: commit.CommitterEmail,
		CreatedAt:      timestamppb.New(commit.CreatedAt.Local()),
		UpdatedAt:      timestamppb.New(commit.CommittedDate.Local()),
	}
}
