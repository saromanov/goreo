package publisher

import (
	"context"
)

type RepositoryRelease struct {
}

// Publisher defines uploading to services
type Publisher interface {
	CreateRelease(ctx context.Context, req RepositoryRelease) (*RepositoryRelease, error)
	GetRelease(ctx context.Context, tag string) (*RepositoryRelease, error)
}
