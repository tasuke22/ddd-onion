package tag

import (
	"context"
)

type TagRepository interface {
	Store(ctx context.Context, tag *Tag) error
	FindByNames(ctx context.Context, names []string) ([]*Tag, error)
}
