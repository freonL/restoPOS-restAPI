package handler

import (
	"encoding/json"
	"net/http"

	"github.com/freonL/restoPOS-restAPI/model"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetItems(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	tbl := []model.Item{}
	db.Find(&tbl)
	respondJSON(w, http.StatusOK, tbl)
}

func GetItem(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	rec := model.Item{}

	if err := db.First(&rec, id).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, rec)
}

func CreateItem(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	rec := model.Item{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&rec); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&rec).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, rec)
}

func UpdateItem(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	rec := getItemOr404(db, id, w, r)
	if rec == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&rec); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&rec).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, rec)
}

func DeleteItem(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	rec := getItemOr404(db, id, w, r)
	if rec == nil {
		return
	}
	if err := db.Delete(&rec).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func DisableItem(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	rec := getItemOr404(db, id, w, r)

	if rec == nil {
		return
	}
	rec.Disable()
	if err := db.Save(&rec).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, rec)
}

func EnableItem(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	rec := getItemOr404(db, name, w, r)
	if rec == nil {
		return
	}
	rec.Enable()
	if err := db.Save(&rec).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, rec)
}

// getEmployeeOr404 gets a employee instance if exists, or respond the 404 error otherwise
func getItemOr404(db *gorm.DB, id string, w http.ResponseWriter, r *http.Request) *model.Item {
	rec := model.Item{}
	if err := db.First(&rec, id).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &rec
}
