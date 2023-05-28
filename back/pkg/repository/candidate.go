package repository

import (
	"backend/pkg/model"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type CandidateRepository struct {
	db *pgxpool.Pool
}

func NewCandidateRepository(db *pgxpool.Pool) *CandidateRepository {
	return &CandidateRepository{db: db}
}

func (r *CandidateRepository) Create(ctx context.Context, candidate *model.CandidateModel) error {
	query := `
		INSERT INTO candidates (first_name, last_name, middle_name, gender, city, district, date_of_birth,
			citizenship, education, email, phone_number, institution, institution_city, faculty,
			specialization, graduation_year, education_level, work_experience, internship_field, work_schedule,
			profile_pic, vk_profile, telegram_profile) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23)
	`

	_, err := r.db.Exec(ctx, query,
		candidate.FirstName, candidate.LastName, candidate.MiddleName, candidate.Gender, candidate.City, candidate.District,
		candidate.DateOfBirth, candidate.Citizenship, candidate.Education, candidate.Email, candidate.PhoneNumber, candidate.Institution,
		candidate.InstitutionCity, candidate.Faculty, candidate.Specialization, candidate.GraduationYear, candidate.EducationLevel,
		candidate.WorkExperience, candidate.InternshipField, candidate.WorkSchedule, candidate.ProfilePic, candidate.VKProfile, candidate.TelegramProfile)

	return err
}
