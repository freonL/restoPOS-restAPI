package app

import (
	"net/http"

	"github.com/freonL/restoPOS-restAPI/handler"
)

func (a *App) GetCategories(w http.ResponseWriter, r *http.Request) {
	handler.GetCategories(a.DB, w, r)
}

func (a *App) CreateCategories(w http.ResponseWriter, r *http.Request) {
	handler.CreateCategories(a.DB, w, r)
}

func (a *App) GetCategory(w http.ResponseWriter, r *http.Request) {
	handler.GetCategory(a.DB, w, r)
}

func (a *App) UpdateCategories(w http.ResponseWriter, r *http.Request) {
	handler.UpdateCategories(a.DB, w, r)
}

func (a *App) DeleteCategories(w http.ResponseWriter, r *http.Request) {
	handler.DeleteCategories(a.DB, w, r)
}

func (a *App) DisableCategories(w http.ResponseWriter, r *http.Request) {
	handler.DisableCategories(a.DB, w, r)
}

func (a *App) EnableCategories(w http.ResponseWriter, r *http.Request) {
	handler.EnableCategories(a.DB, w, r)
}
