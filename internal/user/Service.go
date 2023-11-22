package user

import "context"

type repository interface {
	Insert(ctx context.Context, user *User) (*User, error)
	Update(ctx context.Context, user *User) (*User, error)
	Delete(ctx context.Context) error
	SelectAll(ctx context.Context) ([]User, error)
	Select(ctx, userId int) (*User, error)
}

func NewService() (*Service, error) {

}

type Service struct {
}
