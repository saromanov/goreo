package publisher

import (
	"context"
)

type RepositoryRelease struct {
	ID              *int64  `json:"id,omitempty"`
	TagName         *string `json:"tag_name,omitempty"`
	TargetCommitish *string `json:"target_commitish,omitempty"`
	Name            *string `json:"name,omitempty"`
	Body            *string `json:"body,omitempty"`
	Draft           *bool   `json:"draft,omitempty"`
	Prerelease      *bool   `json:"prerelease,omitempty"`
}

// Publisher defines uploading to services
type Publisher interface {
	CreateRelease(ctx context.Context, req RepositoryRelease) (*RepositoryRelease, error)
	GetRelease(ctx context.Context, tag string) (*RepositoryRelease, error)
}
