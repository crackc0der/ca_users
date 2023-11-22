package user

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Repository struct {
	conn *pgx.Conn
}

func NewRepository(dns string) *Repository {
	conn, err := pgx.Connect(context.Background(), dns)
	if err != nil {
		return nil
	}
	return &Repository{conn: conn}
}

func (r *Repository) Insert(ctx context.Context, user *User) (*User, error) {
	var userId int
	err := r.conn.QueryRow(ctx, "INSERT INTO users (fname, lname, age, email, passwordHash) VALUES ($1, $2, $3, $4, $5) RETURNING userId",
		user.Fname, user.Lname, user.Age, user.Email, user.PasswordHash).Scan(&userId)
	if err != nil {
		return nil, err
	}

	updatedUser, _ := r.Select(ctx, userId)
	return updatedUser, nil
}

func (r *Repository) Update(ctx context.Context, user *User) (*User, error) {
	var u User
	err := r.conn.QueryRow(ctx, "UPDATE user SET fname= $1, lname=$2, age=$3, email=$4 $passwordHash=$5 WHERE userId=$6 RETURNING *",
		user.Fname, user.Lname, user.Age, user.Email, user.PasswordHash, user.UserID).Scan(&u.Fname, &u.Lname, &u.Age, &u.Email, &u.PasswordHash)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *Repository) Delete(ctx context.Context, userID int) error {
	_, err := r.conn.Exec(ctx, "DELETE FROM users WHERE userID=$1", userID)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) SelectAll(ctx context.Context) ([]User, error) {
	var users []User
	rows, err := r.conn.Query(ctx, "SELECT userId, fname, lname, age, email, passwordHasg FROM users")
	if err != nil {
		return nil, err
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

func (r Repository) Select(ctx context.Context, userId int) (*User, error) {
	var user User
	err := r.conn.QueryRow(ctx, "SELECT userId, fname, lname, age, email, passwordHash FROM users WHERE userId =$1", userId).
		Scan(&user.UserID, &user.Fname, &user.Lname, &user.Age, &user.Email, &user.PasswordHash)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
