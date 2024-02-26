package repository

import (
	"context"
	"github.com/tasuke/go-onion/domain/tag"
	"github.com/tasuke/go-onion/infrastructure/db"
	"github.com/tasuke/go-onion/infrastructure/db/dbgen"
)

type tagRepository struct{}

func NewTagRepository() tag.TagRepository {
	return &tagRepository{}
}

func (tr tagRepository) Store(ctx context.Context, tag *tag.Tag) error {
	query := db.GetQuery(ctx)
	if err := query.UpsertTag(ctx, dbgen.UpsertTagParams{
		ID:   tag.ID(),
		Name: tag.Name(),
	}); err != nil {
		return err
	}
	return nil
}

func (tr tagRepository) FindByNames(ctx context.Context, names []string) ([]*tag.Tag, error) {
	query := db.GetQuery(ctx)
	foundTags, err := query.FindByNames(ctx, names)
	if err != nil {
		return nil, err
	}
	result := make([]*tag.Tag, len(foundTags))
	for i, foundTag := range foundTags {
		reconstructedTag, err := tag.ReconstructTag(foundTag.ID, foundTag.Name)
		if err != nil {
			return nil, err
		}
		result[i] = reconstructedTag
	}

	return result, nil
}
