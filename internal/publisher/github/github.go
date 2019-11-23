package github

import (
	"context"
	"net/http"

	"github.com/google/go-github/github"
	"github.com/pkg/errors"
	"github.com/saromanov/goreo/internal/publisher"
	"golang.org/x/oauth2"
)

type client struct {
	api   *github.Client
	Owner string
	Repo  string
}

// New creates connect to Github
func New(token, owner, repo string) publisher.Publisher {
	ts := oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: token,
	})
	tc := oauth2.NewClient(context.TODO(), ts)

	githubClient := github.NewClient(tc)

	return &client{
		api:   githubClient,
		Owner: owner,
		Repo:  repo,
	}
}

func (g *client) CreateRelease(ctx context.Context, req publisher.RepositoryRelease) (*publisher.RepositoryRelease, error) {
	_, res, err := g.api.Repositories.CreateRelease(context.TODO(), g.Owner, g.Repo, &github.RepositoryRelease{
		Body:  req.Body,
		Name:  req.Name,
		Draft: req.Draft,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to create a release")
	}

	if res.StatusCode != http.StatusCreated {
		return nil, errors.Errorf("create release: invalid status: %s", res.Status)
	}

	return &req, nil
}

func (g *client) GetRelease(ctx context.Context, tag string) (*publisher.RepositoryRelease, error) {
	_, _, err := g.api.Repositories.GetReleaseByTag(context.TODO(), g.Owner, g.Repo, tag)
	if err != nil {
		return nil, errors.Wrap(err, "unable to find release")
	}

	return &publisher.RepositoryRelease{}, nil
}
