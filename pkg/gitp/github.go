package gitp

import (
	"context"
	"strings"

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

func (g *Github) ListRepos(ctx context.Context) ([]*provider.Repo, error) {
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

func (g *Github) ListBranches(ctx context.Context, fullName string) ([]*provider.Branch, error) {
	owner, repo := getOwnerRepo(fullName)
	results, _, err := g.Repositories.ListBranches(ctx, owner, repo, &github.ListOptions{})
	if err != nil {
		return nil, err
	}

	branches := make([]*provider.Branch, 0)
	for _, ret := range results {
		branches = append(branches, &provider.Branch{
			Name:   ret.GetName(),
			Commit: pCommitFromGh(ret.Commit.GetCommit()),
		})
	}

	return branches, nil
}

func (g *Github) ListTags(ctx context.Context, fullName string) ([]*provider.Tag, error) {
	owner, repo := getOwnerRepo(fullName)
	results, _, err := g.Repositories.ListTags(ctx, owner, repo, &github.ListOptions{})
	if err != nil {
		return nil, err
	}

	tags := make([]*provider.Tag, 0)
	for _, ret := range results {
		tags = append(tags, &provider.Tag{
			Name:   ret.GetName(),
			Commit: pCommitFromGh(ret.Commit),
		})
	}

	return tags, nil
}

func getOwnerRepo(fullName string) (string, string) {
	items := strings.Split(fullName, "/")
	return items[0], items[1]
}

func pCommitFromGh(commit *github.Commit) *provider.Commit {
	return &provider.Commit{
		Sha1:           commit.GetSHA(),
		ShortId:        commit.GetSHA()[:7],
		Message:        commit.GetMessage(),
		CommitterName:  commit.Committer.GetName(),
		CommitterEmail: commit.Committer.GetEmail(),
		CreatedAt:      timestamppb.New(commit.Committer.Date.Local()),
		UpdatedAt:      timestamppb.New(commit.Committer.Date.Local()),
	}
}
