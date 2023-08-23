package handlers

import (
	"encoding/json"
	"mini-clickup/internal/api/models"
	"net/http"

	"gorm.io/gorm"
)

func RemoveTask(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var tsk models.Task
	err := json.NewDecoder(r.Body).Decode(&tsk)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	err = db.Delete(&models.TaskLog{}, tsk.User).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = db.Delete(&models.Task{}, tsk.User).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
