package user

import (
	"context"
	"fmt"

	//nolint
	"github.com/jackc/pgx/v5"
)

type Repository struct {
	conn *pgx.Conn
}

func NewRepository(dns string) (*Repository, error) {
	conn, err := pgx.Connect(context.Background(), dns)
	if err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}

	return &Repository{conn: conn}, nil
}

func (r *Repository) Insert(ctx context.Context, user *User) (*User, error) {
	var newUser User

	query := `INSERT INTO users (fname, lname, age, email, passwordHash) VALUES ($1, $2, $3, $4, $5) 
	RETURNING userId, fname, lname, age, email, passwordHash`
	err := r.conn.QueryRow(ctx, query,
		user.Fname, user.Lname, user.Age, user.Email, user.PasswordHash).Scan(&newUser.UserID, &newUser.Fname,
		&newUser.Lname, &newUser.Age, &newUser.Email, &newUser.PasswordHash)

	if err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}

	return &newUser, nil
}

func (r *Repository) Update(ctx context.Context, user *User) (*User, error) {
	var updatedUser User

	query := `UPDATE user SET fname= $1, lname=$2, age=$3, email=$4 $passwordHash=$5 WHERE 
	userId=$6 RETURNING userId, fname, lname, age, email, passwordHash`
	err := r.conn.QueryRow(ctx, query,
		user.Fname, user.Lname, user.Age, user.Email, user.PasswordHash, user.UserID).Scan(&updatedUser.Fname,
		&updatedUser.Lname, &updatedUser.Age, &updatedUser.Email, &updatedUser.PasswordHash)

	if err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}

	return &updatedUser, nil
}

func (r *Repository) Delete(ctx context.Context, userID int) error {
	query := "DELETE FROM users WHERE userID=$1"
	_, err := r.conn.Exec(ctx, query, userID)

	if err != nil {
		return fmt.Errorf("error: %w", err)
	}

	return nil
}

func (r *Repository) SelectAll(ctx context.Context) ([]User, error) {
	var users []User

	query := "SELECT userId, fname, lname, age, email, passwordHasg FROM users"
	rows, err := r.conn.Query(ctx, query)

	if err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}

	for rows.Next() {
		var user User
		err := rows.Scan(&user.UserID, &user.Fname, &user.Lname, &user.Age, &user.Email, &user.PasswordHash)

		if err != nil {
			continue
		}

		users = append(users, user)
	}

	return users, nil
}

func (r Repository) Select(ctx context.Context, userID int) (*User, error) {
	var user User

	query := "SELECT userId, fname, lname, age, email, passwordHash FROM users WHERE userId =$1"
	err := r.conn.QueryRow(ctx, query, userID).
		Scan(&user.UserID, &user.Fname, &user.Lname, &user.Age, &user.Email, &user.PasswordHash)

	if err != nil {
		return nil, fmt.Errorf("erorr: %w", err)
	}

	return &user, nil
}
