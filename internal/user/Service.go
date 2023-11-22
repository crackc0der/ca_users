package user

import "context"

type repository interface {
	Insert(ctx context.Context, user *User) (*User, error)
}

type Service struct {
}
