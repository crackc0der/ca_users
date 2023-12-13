package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) Insert(ctx context.Context, user *User) (*User, error) {
	args := m.Called(ctx, user)
	return args.Get(0).(*User), args.Error(0)
}

func (m *MockRepository) Update(ctx context.Context, user *User) (*User, error) {
	args := m.Called(ctx, user)
	return args.Get(0).(*User), args.Error(0)
}

func (m *MockRepository) Delete(ctx context.Context, userId uuid.UUID) error {
	args := m.Called(ctx, userId)
	return args.Error(0)
}

func (m *MockRepository) SelectAll(ctx context.Context) ([]User, error) {
	args := m.Called(ctx)
	return args.Get(0).([]User), args.Error(0)
}

func (m *MockRepository) Select(ctx context.Context, userId uuid.UUID) (*User, error) {
	args := m.Called(ctx, userId)
	return args.Get(0).(*User), args.Error(1)
}
