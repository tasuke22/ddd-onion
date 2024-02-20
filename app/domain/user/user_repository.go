//go:generate mockgen -package $GOPACKAGE -source $GOFILE -destination mock_$GOFILE
package user

import "context"

type UserRepository interface {
	Store(ctx context.Context, user *User) error
	FindByName(ctx context.Context, name string) (*User, error)
}
