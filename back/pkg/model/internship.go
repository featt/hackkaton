package model

import (
	"time"
)

type Internship struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	FileURL     string    `json:"file_url"` // URL ссылка на файл, хранящийся в системе хранения файлов
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
