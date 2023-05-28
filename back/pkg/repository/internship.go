package repository

import (
	"backend/pkg/model"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type InternshipRepository struct {
	db *pgxpool.Pool
}

func NewInternshipRepository(db *pgxpool.Pool) *InternshipRepository {
	return &InternshipRepository{db: db}
}

func (r *InternshipRepository) CreateInternship(ctx context.Context, internship model.Internship) (int64, error) {
	var id int64
	err := r.db.QueryRow(ctx,
		"INSERT INTO internships (user_id, title, description, start_date, end_date, file_url) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		internship.UserID,
		internship.Title,
		internship.Description,
		internship.StartDate,
		internship.EndDate,
		internship.FileURL,
	).Scan(&id)
	return id, err
}

func (r *InternshipRepository) GetInternshipByID(ctx context.Context, id int64) (model.Internship, error) {
	var internship model.Internship
	err := r.db.QueryRow(ctx,
		"SELECT id, user_id, title, description, start_date, end_date, file_url FROM internships WHERE id=$1",
		id,
	).Scan(&internship.ID, &internship.UserID, &internship.Title, &internship.Description, &internship.StartDate, &internship.EndDate, &internship.FileURL)
	return internship, err
}

func (r *InternshipRepository) UpdateInternship(ctx context.Context, id int64, internship model.Internship) error {
	_, err := r.db.Exec(ctx,
		"UPDATE internships SET user_id=$1, title=$2, description=$3, start_date=$4, end_date=$5, file_url=$6 WHERE id=$7",
		internship.UserID,
		internship.Title,
		internship.Description,
		internship.StartDate,
		internship.EndDate,
		internship.FileURL,
		id,
	)
	return err
}

func (r *InternshipRepository) DeleteInternship(ctx context.Context, id int64) error {
	_, err := r.db.Exec(ctx,
		"DELETE FROM internships WHERE id=$1",
		id,
	)
	return err
}

func (r *InternshipRepository) GetAllInternships(ctx context.Context) ([]model.Internship, error) {
	rows, err := r.db.Query(ctx,
		"SELECT id, user_id, title, description, start_date, end_date, file_url FROM internships")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var internships []model.Internship
	for rows.Next() {
		var internship model.Internship
		err := rows.Scan(&internship.ID, &internship.UserID, &internship.Title, &internship.Description, &internship.StartDate, &internship.EndDate, &internship.FileURL)
		if err != nil {
			return nil, err
		}
		internships = append(internships, internship)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return internships, nil
}
