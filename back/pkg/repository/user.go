// repository/user.go
package repository

import (
	"backend/pkg/model"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, user model.User) (int64, error) {
	var id int64
	err := r.db.QueryRow(ctx,
		"INSERT INTO users (name, lastname, email, password, role) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		user.Name,
		user.LastName,
		user.Email,
		user.Password,
		string(user.Role),
	).Scan(&id)
	return id, err
}

func (r *UserRepository) GetUserByEmailAndPassword(ctx context.Context, email string, password string) (model.User, error) {
	var user model.User
	err := r.db.QueryRow(ctx,
		"SELECT id, name, email, password, role FROM users WHERE email=$1 AND password=$2",
		email,
		password,
	).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role)
	return user, err
}

func (r *UserRepository) GetUserByID(ctx context.Context, id int64) (model.User, error) {
	var user model.User
	err := r.db.QueryRow(ctx,
		"SELECT id, name, email, password, role, loyalty_points FROM users WHERE id=$1",
		id,
	).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role, &user.LoyaltyPoints)
	return user, err
}

func (r *UserRepository) UpdateUser(ctx context.Context, id int64, user model.User) error {
	_, err := r.db.Exec(ctx,
		"UPDATE users SET name=$1, email=$2, password=$3, role=$4 WHERE id=$5",
		user.Name,
		user.Email,
		user.Password,
		string(user.Role),
		id,
	)
	return err
}

func (r *UserRepository) DeleteUser(ctx context.Context, id int64) error {
	_, err := r.db.Exec(ctx,
		"DELETE FROM users WHERE id=$1",
		id,
	)
	return err
}

func (r *UserRepository) UpdateUserRole(ctx context.Context, id int64, role string) error {
	_, err := r.db.Exec(ctx,
		"UPDATE users SET role=$1 WHERE id=$2",
		role,
		id,
	)
	return err
}

func (r *UserRepository) GetAllUsers(ctx context.Context) ([]model.User, error) {
	rows, err := r.db.Query(ctx, "SELECT id, name, email, password, role FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) UpdateUserLoyalty(ctx context.Context, id int64, loyaltyPoints int) error {
	_, err := r.db.Exec(ctx,
		"UPDATE users SET loyalty_points=$1 WHERE id=$2",
		loyaltyPoints,
		id,
	)
	if err != nil {
		log.Printf("Error executing database update: %v", err)
		return err
	}
	return nil
}
