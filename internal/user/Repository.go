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

func (r *Repository) Insert() (int, error) {
	var userId int
	err := r.conn.QueryRow(context.Background(), "INSERT INTO users (fname, lname, age, email, passwordHash RETURNING userId").Scan(&userId)
	if err != nil {
		return -1, err
	}

	return userId, nil
}

func (r *Repository) Update(user *User) (*User, error) {
	var u User
	err := r.conn.QueryRow(context.Background(), "UPDATE user SET fname= $1, lname=$2, age=$3, email=$4 $passwordHash=$5 WHERE userId=$6 RETURNING *",
		user.Fname, user.Lname, user.Age, user.Email, user.PasswordHash, user.UserID).Scan(&u.Fname, &u.Lname, &u.Age, &u.Email, &u.PasswordHash)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *Repository) Delete(userId int) error {
	_, err := r.conn.Exec(context.Background(), "DELETE FROM users WHERE userID=$1", userId)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) SelectAll() ([]User, error) {
	var users []User
	rows, err := r.conn.Query(context.Background(), "SELECT userId, fname, lname, age, email, passwordHasg FROM users")
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

func (r Repository) Select(userId int) (*User, error) {
	var user User
	err := r.conn.QueryRow(context.Background(), "SELECT userId, fname, lname, age, email, passwordHash FROM users WHERE userId =$1", userId).
		Scan(&user.UserID, &user.Fname, &user.Lname, &user.Age, &user.Email, &user.PasswordHash)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
