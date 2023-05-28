package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/xuri/excelize/v2"
)

type ExcelHandler struct {
}

func NewExcelHandler() *ExcelHandler {
	return &ExcelHandler{}
}

type Data struct {
	Info []map[string]string `json:"info"`
}

func (h *ExcelHandler) SaveExcel(w http.ResponseWriter, r *http.Request) {
	var data Data
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	f := excelize.NewFile()
	for i, info := range data.Info {
		for key, value := range info {
			f.SetCellValue("Sheet1", key+strconv.Itoa(i+2), value)
		}
	}

	if err := f.SaveAs("./file.xlsx"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *ExcelHandler) DownloadExcel(w http.ResponseWriter, r *http.Request) {
	f, err := excelize.OpenFile("./file.xlsx")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data Data
	rows, _ := f.GetRows("Sheet1")
	for i, row := range rows {
		if i == 0 {
			continue // Skip the header row
		}
		info := make(map[string]string)
		for j, colCell := range row {
			info[rows[0][j]] = colCell
		}
		data.Info = append(data.Info, info)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
