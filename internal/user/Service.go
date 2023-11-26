package user

import (
	"context"
	"fmt"
)

type repository interface {
	Insert(ctx context.Context, user *User) (*User, error)
	Update(ctx context.Context, user *User) (*User, error)
	Delete(ctx context.Context, userID int) error
	SelectAll(ctx context.Context) ([]User, error)
	Select(ctx context.Context, userID int) (*User, error)
}

func NewService(repo repository) *Service {
	return &Service{repo: repo}
}

type Service struct {
	repo repository
}

func (s *Service) AddUser(ctx context.Context, user *User) (*User, error) {
	fmt.Println(user)
	addedUser, err := s.repo.Insert(ctx, user)

	if err != nil {
		return nil, fmt.Errorf("error insering user: %w", err)
	}

	return addedUser, nil
}

func (s *Service) UpdateUser(ctx context.Context, user *User) (*User, error) {
	updatedUser, err := s.repo.Update(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("error updating user: %w", err)
	}

	return updatedUser, nil
}

func (s *Service) DeleteUser(ctx context.Context, userID int) error {
	err := s.repo.Delete(ctx, userID)
	if err != nil {
		return fmt.Errorf("error deleting user: %w", err)
	}

	return nil
}

func (s *Service) GetAllUsers(ctx context.Context) ([]User, error) {
	users, err := s.repo.SelectAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting all users: %w", err)
	}

	return users, nil
}

func (s *Service) GetUser(ctx context.Context, userID int) (*User, error) {
	user, err := s.repo.Select(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("error ger user: %w", err)
	}

	return user, nil
}
