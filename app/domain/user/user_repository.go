//go:generate mockgen -package $GOPACKAGE -source $GOFILE -destination mock_$GOFILE
package user

import "context"

type UserRepository interface {
	Save(ctx context.Context, user *User) error
	FindByName(ctx context.Context, name string) (*User, error)
	FindByUserID(ctx context.Context, id string) (*User, error)
	UpdateUser(ctx context.Context, user *User) error
}
