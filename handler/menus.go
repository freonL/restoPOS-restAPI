package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/freonL/restoPOS-restAPI/model"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetMenus(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	// var limit, offset int64
	limit, _ := strconv.Atoi(r.FormValue("limit"))
	offset, _ := strconv.Atoi(r.FormValue("offset"))
	if limit <= 0 {
		limit = 10
	}

	if offset < 0 {
		offset = 0
	}
	// fmt.Println(limit, offset)

	tbl := []model.Item{}

	db.Preload("Category").Limit(limit).Offset(offset).Find(&tbl)
	respondJSON(w, http.StatusOK, tbl)
}

func GetMenu(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	rec := model.Item{}

	if err := db.First(&rec, id).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, rec)
}

func CreateMenu(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
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

func UpdateMenu(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	rec := getMenuOr404(db, id, w, r)
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

func DeleteMenu(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	rec := getMenuOr404(db, id, w, r)
	if rec == nil {
		return
	}
	if err := db.Delete(&rec).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func DisableMenu(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	rec := getMenuOr404(db, id, w, r)

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

func EnableMenu(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	rec := getMenuOr404(db, name, w, r)
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

// getMenuOr404 gets a item instance if exists, or respond the 404 error otherwise
func getMenuOr404(db *gorm.DB, id string, w http.ResponseWriter, r *http.Request) *model.Item {
	rec := model.Item{}
	if err := db.First(&rec, id).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &rec
}
