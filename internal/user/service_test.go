package user

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddUser(t *testing.T) {
}

func TestUpdateUser(t *testing.T) {
}

func TestDeleteUser(t *testing.T) {
}

func TestGetAllUsers(t *testing.T) {
}

func TestGetUser(t *testing.T) {
	repo := new(MockRepository)

	GetUser := func(id uuid.UUID, user *User, err error) {
		repo.On("Select", mock.Anything, id).Return(user, err).Once()
	}

	tests := []struct {
		name    string
		setup   func()
		id      uuid.UUID
		want    *User
		wantErr error
	}{
		{
			name: "success",
			setup: func() {
				GetUser(
					uuid.MustParse("ccae37ea-d41e-4371-a3a3-89203b9e2608"),
					&User{
						UserID:       uuid.MustParse("ccae37ea-d41e-4371-a3a3-89203b9e2608"),
						Fname:        "Maga",
						Lname:        "Magov",
						Age:          33,
						Email:        "maga@gmail.com",
						PasswordHash: "asdh327d23b7328",
					},
					nil,
				)
			},
			id: uuid.MustParse("ccae37ea-d41e-4371-a3a3-89203b9e2608"),
			want: &User{
				UserID:       uuid.MustParse("ccae37ea-d41e-4371-a3a3-89203b9e2608"),
				Fname:        "Maga",
				Lname:        "Magov",
				Age:          33,
				Email:        "maga@gmail.com",
				PasswordHash: "asdh327d23b7328",
			},
			wantErr: nil,
		},
		{
			name: "success",
			setup: func() {
				GetUser(
					uuid.MustParse("ccae37ea-d41e-4371-a3a3-89203b9e2608"),
					&User{
						UserID:       uuid.MustParse("ccae37ea-d41e-4371-a3a3-89203b9e2608"),
						Fname:        "Daga",
						Lname:        "Dagov",
						Age:          34,
						Email:        "daga@gmail.com",
						PasswordHash: "asdh327d23b7328",
					},
					nil,
				)
			},
			id: uuid.MustParse("ccae37ea-d41e-4371-a3a3-89203b9e2608"),
			want: &User{
				UserID:       uuid.MustParse("ccae37ea-d41e-4371-a3a3-89203b9e2608"),
				Fname:        "Daga",
				Lname:        "Dagov",
				Age:          34,
				Email:        "daga@gmail.com",
				PasswordHash: "asdh327d23b7328",
			},
			wantErr: nil,
		},
	}

	svc := NewService(repo)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()

			got, err := svc.GetUser(context.Background(), tt.id)
			if err != nil && assert.Error(t, tt.wantErr, err.Error()) {
				assert.EqualError(t, err, tt.wantErr.Error())
			} else {
				assert.NoError(t, tt.wantErr)
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
