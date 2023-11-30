package user

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

type Repository struct {
	conn *pgx.Conn
}

func NewRepository(dns string) (*Repository, error) {
	conn, err := pgx.Connect(context.Background(), dns)
	if err != nil {
		return nil, fmt.Errorf("failed create new repository in method NewRepository: %w", err)
	}

	return &Repository{conn: conn}, nil
}

func (r *Repository) Insert(ctx context.Context, user *User) (*User, error) {
	var newUser User

	password, err := r.passwordHash(user.PasswordHash)

	if err != nil {
		return nil, fmt.Errorf("error encrypting password: %w", err)
	}

	query := `INSERT INTO users (fname, lname, age, email, passwordHash) VALUES ($1, $2, $3, $4, $5) 
	RETURNING userId, fname, lname, age, email, passwordHash`
	err = r.conn.QueryRow(ctx, query,
		user.Fname, user.Lname, user.Age, user.Email, password).Scan(&newUser.UserID, &newUser.Fname,
		&newUser.Lname, &newUser.Age, &newUser.Email, &newUser.PasswordHash)

	if err != nil {
		return nil, fmt.Errorf("error in method insert: %w", err)
	}

	return &newUser, nil
}

func (r *Repository) Update(ctx context.Context, user *User) (*User, error) {
	var updatedUser User

	password, err := r.passwordHash(user.PasswordHash)

	if err != nil {
		return nil, fmt.Errorf("error encrypting password: %w", err)
	}

	query := `UPDATE users SET fname = $1, lname = $2, age = $3, email = $4, passwordHash = $5 WHERE 
	userId=$6 RETURNING fname, lname, age, email, passwordHash`
	err = r.conn.QueryRow(ctx, query,
		user.Fname, user.Lname, user.Age, user.Email, password, user.UserID).Scan(&updatedUser.Fname,
		&updatedUser.Lname, &updatedUser.Age, &updatedUser.Email, &updatedUser.PasswordHash)

	if err != nil {
		return nil, fmt.Errorf("error in method update: %w", err)
	}

	return &updatedUser, nil
}

func (r *Repository) Delete(ctx context.Context, userID int) error {
	query := "DELETE FROM users WHERE userID=$1"
	_, err := r.conn.Exec(ctx, query, userID)

	if err != nil {
		return fmt.Errorf("error in method Delete: %w", err)
	}

	return nil
}

func (r *Repository) SelectAll(ctx context.Context) ([]User, error) {
	var users []User

	query := "SELECT userId, fname, lname, age, email, passwordHash FROM users"
	rows, err := r.conn.Query(ctx, query)

	if err != nil {
		return nil, fmt.Errorf("error in method SelectAll: %w", err)
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

	query := "SELECT userID, fname, lname, age, email, passwordHash FROM users WHERE userId = $1"
	err := r.conn.QueryRow(ctx, query, userID).
		Scan(&user.UserID, &user.Fname, &user.Lname, &user.Age, &user.Email, &user.PasswordHash)

	if err != nil {
		return nil, fmt.Errorf("error in method Select: %w", err)
	}

	return &user, nil
}

func (r *Repository) passwordHash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("password encryption error: %w", err)
	}

	return string(hashedPassword), nil
}
