package handler

import (
	"encoding/json"
	"net/http"

	"github.com/freonL/restoPOS-restAPI/model"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetCategories(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	tbl := []model.ItemCategory{}
	db.Find(&tbl)
	respondJSON(w, http.StatusOK, tbl)
}

func GetCategory(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	rec := model.ItemCategory{}

	if err := db.First(&rec, id).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, rec)
}

func CreateCategories(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	rec := model.ItemCategory{}

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

func UpdateCategories(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	rec := getCategoriesOr404(db, id, w, r)
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

func DeleteCategories(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	rec := getCategoriesOr404(db, id, w, r)
	if rec == nil {
		return
	}
	if err := db.Delete(&rec).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func DisableCategories(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	rec := getCategoriesOr404(db, id, w, r)

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

func EnableCategories(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	rec := getCategoriesOr404(db, name, w, r)
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

// getCategoriesOr404 gets a item instance if exists, or respond the 404 error otherwise
func getCategoriesOr404(db *gorm.DB, id string, w http.ResponseWriter, r *http.Request) *model.ItemCategory {
	rec := model.ItemCategory{}
	if err := db.First(&rec, id).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &rec
}
