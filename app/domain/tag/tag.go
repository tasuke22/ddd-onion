package tag

import "github.com/tasuke/go-pkg/ulid"

type Tag struct {
	id   string
	name string
}

func (t *Tag) ID() string {
	return t.id
}
func (t *Tag) Name() string {
	return t.name
}

func NewTag(
	name string,
) (*Tag, error) {
	return &Tag{
		id:   ulid.NewULID(),
		name: name,
	}, nil
}

func ReconstructTag(
	id string,
	name string,
) (*Tag, error) {
	return &Tag{
		id:   id,
		name: name,
	}, nil
}
