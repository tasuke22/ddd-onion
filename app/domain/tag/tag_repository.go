package tag

import (
	"context"
)

type TagRepository interface {
	Store(ctx context.Context, tag *Tag) error
	FindByName(ctx context.Context, name string) (*Tag, error)
	FindByNames(ctx context.Context, names []string) ([]*Tag, error)
}
