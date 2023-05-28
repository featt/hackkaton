package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

type TechnicalRepository struct {
	db *pgxpool.Pool
}

func NewTechnicalRepository(db *pgxpool.Pool) *TechnicalRepository {
	return &TechnicalRepository{db: db}
}

func (r *TechnicalRepository) CheckDatabaseConnection(ctx context.Context) error {
	_, err := r.db.Exec(ctx, "SELECT 1")
	return err
}

func (r *TechnicalRepository) GetAllTables(ctx context.Context) ([]string, error) {
	rows, err := r.db.Query(ctx, "SELECT table_name FROM information_schema.tables WHERE table_schema = 'public'")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tables []string
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			return nil, err
		}
		tables = append(tables, tableName)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tables, nil
}

func (r *TechnicalRepository) GetTableData(ctx context.Context, tableName string, limit int) ([]string, error) {
	query := fmt.Sprintf("SELECT * FROM %s LIMIT %d", tableName, limit)

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns := rows.FieldDescriptions()
	columnNames := make([]string, len(columns))
	for i, col := range columns {
		columnNames[i] = string(col.Name)
	}

	data := make([]string, 0)

	for rows.Next() {
		values := make([]interface{}, len(columnNames))
		columnPointers := make([]interface{}, len(columnNames))
		for i := range columnNames {
			columnPointers[i] = &values[i]
		}

		if err := rows.Scan(columnPointers...); err != nil {
			return nil, err
		}

		rowData := make([]string, len(columnNames))
		for i, v := range values {
			if v != nil {
				rowData[i] = fmt.Sprintf("%v", v)
			} else {
				rowData[i] = "NULL"
			}
		}

		data = append(data, rowData...)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

func (r *TechnicalRepository) GetTableFields(ctx context.Context, tableName string) ([]string, error) {
	query := fmt.Sprintf("SELECT column_name FROM information_schema.columns WHERE table_name = '%s'", tableName)

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	fields := make([]string, 0)

	for rows.Next() {
		var field string
		if err := rows.Scan(&field); err != nil {
			return nil, err
		}
		fields = append(fields, field)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return fields, nil
}
