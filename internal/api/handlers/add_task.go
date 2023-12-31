package handlers

import (
	"encoding/json"
	"mini-clickup/internal/api/models"
	"net/http"

	"gorm.io/gorm"
)

func AddTask(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var tsk models.Task
	err := json.NewDecoder(r.Body).Decode(&tsk)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = db.Create(&tsk).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
