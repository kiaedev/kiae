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
			Name:      ret.GetName(),
			CreatedAt: timestamppb.New(ret.Commit.Commit.Committer.Date.Local()),
			UpdatedAt: timestamppb.New(ret.Commit.Commit.Committer.Date.Local()),
		})
	}

	return branches, nil
}

func (g *Github) ListCommits(ctx context.Context, fullName, refName string) ([]*provider.Commit, error) {
	owner, repo := getOwnerRepo(fullName)
	results, _, err := g.Repositories.ListCommits(ctx, owner, repo, &github.CommitsListOptions{SHA: refName})
	if err != nil {
		return nil, err
	}

	commits := make([]*provider.Commit, 0)
	for _, ret := range results {
		commits = append(commits, &provider.Commit{
			CommitId:       ret.Commit.GetSHA(),
			ShortId:        ret.Commit.GetSHA()[:7],
			Message:        ret.Commit.GetMessage(),
			CommitterName:  ret.Committer.GetName(),
			CommitterEmail: ret.Committer.GetEmail(),
			CreatedAt:      timestamppb.New(ret.Commit.Committer.Date.Local()),
			UpdatedAt:      timestamppb.New(ret.Commit.Committer.Date.Local()),
		})
	}

	return commits, nil
}

func getOwnerRepo(fullName string) (string, string) {
	items := strings.Split(fullName, "/")
	return items[0], items[1]
}
