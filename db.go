package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Role      string    `json:"role"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"createdAt"`
}

func openDB() (*sql.DB, error) {
	dsn := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// hash calculates the bcrypt hash of a plaintext password.
func hash(plain string) ([]byte, error) {
	cost := 11
	return bcrypt.GenerateFromPassword([]byte(plain), cost)
}

// listUsers returns a list of all users.
func listUsers() ([]*User, error) {
	query := `select id, name, email, role, active from users`
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := openDB()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error leyendo usuarios de la base de datos:\n%s", err.Error()))
	}

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*User

	for rows.Next() {
		u := &User{}
		err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Role, &u.Active)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// updatePassword changes a user's password in the database.
func updatePassword(email string, password string) error {
	hashed, err := hash(password)
	if err != nil {
		return errors.New(fmt.Sprintf("Error encriptando la password:\n%s", err.Error()))
	}

	query := `
		update users set password = $1
		where email = $2
	`

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := openDB()
	if err != nil {
		return errors.New(fmt.Sprintf("Error conectando a la base de datos:\n%s", err.Error()))
	}

	result, err := db.ExecContext(ctx, query, hashed, email)
	if err != nil {
		return errors.New(fmt.Sprintf("Error actualizando la base de datos:\n%s", err.Error()))
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return errors.New(fmt.Sprintf("Error actualizando la base de datos:\n%s", err.Error()))
	}
	if affected != 1 {
		return errors.New(fmt.Sprintf("No se actualizó ninguna cuenta, seguro que el email %q existe? (Rows affected: %d)", email, affected))
	}

	return nil
}

// createAdmin creates a new user account with role `admin`.
func createAdmin(name, email, password string) error {
	hashed, err := hash(password)
	if err != nil {
		return errors.New(fmt.Sprintf("Error encriptando la password:\n%s", err.Error()))
	}

	query := `
		insert into users (name, email, password, role)
		values ($1, $2, $3, 'admin')
	`

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := openDB()
	if err != nil {
		return errors.New(fmt.Sprintf("Error conectando a la base de datos:\n%s", err.Error()))
	}

	result, err := db.ExecContext(ctx, query, name, email, hashed)
	if err != nil {
		return errors.New(fmt.Sprintf("Error actualizando la base de datos:\n%s", err.Error()))
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return errors.New(fmt.Sprintf("Error actualizando la base de datos:\n%s", err.Error()))
	}
	if affected != 1 {
		return errors.New(fmt.Sprintf("No se actualizó ninguna cuenta (Rows affected: %d)", affected))
	}

	return nil
}
