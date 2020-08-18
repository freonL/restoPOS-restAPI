package app

import (
	"net/http"

	"github.com/freonL/restoPOS-restAPI/handler"
)

func (a *App) GetMenus(w http.ResponseWriter, r *http.Request) {
	handler.GetMenus(a.DB, w, r)
}

func (a *App) CreateMenu(w http.ResponseWriter, r *http.Request) {
	handler.CreateMenu(a.DB, w, r)
}

func (a *App) GetMenu(w http.ResponseWriter, r *http.Request) {
	handler.GetMenu(a.DB, w, r)
}

func (a *App) UpdateMenu(w http.ResponseWriter, r *http.Request) {
	handler.UpdateMenu(a.DB, w, r)
}

func (a *App) DeleteMenu(w http.ResponseWriter, r *http.Request) {
	handler.DeleteMenu(a.DB, w, r)
}

func (a *App) DisableMenu(w http.ResponseWriter, r *http.Request) {
	handler.DisableMenu(a.DB, w, r)
}

func (a *App) EnableMenu(w http.ResponseWriter, r *http.Request) {
	handler.EnableMenu(a.DB, w, r)
}
