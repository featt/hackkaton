package handler

import (
	"backend/pkg/service"
	"encoding/json"
	"log"
	"net/http"
)

type TechnicalHandler struct {
	service *service.TechnicalService
}

func NewTechnicalHandler(service *service.TechnicalService) *TechnicalHandler {
	return &TechnicalHandler{service: service}
}

func (h *TechnicalHandler) CheckDatabaseConnection(w http.ResponseWriter, r *http.Request) {
	err := h.service.CheckDatabaseConnection(r.Context())
	if err != nil {
		log.Printf("Database connection error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error connecting to the database"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Database connection is successful"))
}

func (h *TechnicalHandler) GetAllTables(w http.ResponseWriter, r *http.Request) {
	tables, err := h.service.GetAllTables(r.Context())
	if err != nil {
		log.Printf("Error retrieving tables: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error retrieving tables"))
		return
	}

	tableData := make(map[string][]map[string]interface{})

	for _, table := range tables {
		rows, err := h.service.GetTableData(r.Context(), table, 10)
		if err != nil {
			log.Printf("Error retrieving data for table %s: %v\n", table, err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error retrieving table data"))
			return
		}
		tableFields, err := h.service.GetTableFields(r.Context(), table)
		if err != nil {
			log.Printf("Error retrieving fields for table %s: %v\n", table, err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error retrieving table fields"))
			return
		}
		tableRows := make([]map[string]interface{}, 0)
		for i := 0; i < len(rows); i += len(tableFields) {
			rowData := make(map[string]interface{})
			for j, field := range tableFields {
				rowData[field] = rows[i+j]
			}
			tableRows = append(tableRows, rowData)
		}
		tableData[table] = tableRows
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tableData)
}
