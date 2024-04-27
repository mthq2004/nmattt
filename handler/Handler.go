package handler

import (
	"encoding/json"
	"net/http"

	"github.com/congmanh18/NMATTT_AESRSA/model"
	"gorm.io/gorm"
)

func ReadTableData(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Set content type header to application/json
		w.Header().Set("Content-Type", "application/json")
		// Allow all origins
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// Allow only POST and OPTIONS
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		// Allow only Content-Type header
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		if r.Method != "GET" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		// Đọc dữ liệu từ bảng
		var data []model.Data
		if err := db.Find(&data).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Trả về dữ liệu dưới dạng JSON
		if err := json.NewEncoder(w).Encode(data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func DeleteByID(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		// Allow all origins
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// Allow only POST and OPTIONS
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		// Allow only Content-Type header
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		if r.Method != "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		var requestBody map[string]string
		if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
			http.Error(w, "Invalid JSON request body", http.StatusBadRequest)
			return
		}
		idParam, ok := requestBody["id"]
		if !ok {
			http.Error(w, "ID is missing in the request body", http.StatusBadRequest)
			return
		}
		if err := db.Where("id = ?", idParam).Delete(&model.Data{}).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		response := struct {
			Message string `json:"message"`
		}{
			Message: "Record deleted successfully",
		}
		json.NewEncoder(w).Encode(response)
	}
}
