package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/freonL/restoPOS-restAPI/handler"
	"github.com/freonL/restoPOS-restAPI/model"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// App initialize with predefined configuration
func (a *App) Initialize() {
	godotenv.Load(".env")
	Type := os.Getenv("DB_TYPE")
	Host := os.Getenv("DB_HOST")
	Port := os.Getenv("DB_PORT")
	User := os.Getenv("DB_USER")
	Pass := os.Getenv("DB_PASS")
	Name := os.Getenv("DB_NAME")

	var db *gorm.DB
	var err error
	if Type == "postgres" {
		SSL := os.Getenv("DB_SSL")
		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			Host,
			Port,
			User,
			Pass,
			Name,
			SSL,
		)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	} else {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			User,
			Pass,
			Host,
			Port,
			Name,
		)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	}

	if err != nil {
		log.Fatal("Could not connect database")
	}
	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

// Set all required routers
func (a *App) setRouters() {
	// Routing for handling the projects
	a.Get("/api/menus", a.GetItems)
	a.Post("/api/menus", a.CreateItem)
	a.Get("/api/menus/{id}", a.GetItem)
	a.Put("/api/menus/{id}", a.UpdateItem)
	a.Delete("/api/menus/{id}", a.DeleteItem)
	a.Put("/api/menus/{id}/disable", a.DisableItem)
	a.Put("/api/menus/{id}/enable", a.EnableItem)
}

// Wrap the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Wrap the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Wrap the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Wrap the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

func (a *App) GetItems(w http.ResponseWriter, r *http.Request) {
	handler.GetItems(a.DB, w, r)
}

func (a *App) CreateItem(w http.ResponseWriter, r *http.Request) {
	handler.CreateItem(a.DB, w, r)
}

func (a *App) GetItem(w http.ResponseWriter, r *http.Request) {
	handler.GetItem(a.DB, w, r)
}

func (a *App) UpdateItem(w http.ResponseWriter, r *http.Request) {
	handler.UpdateItem(a.DB, w, r)
}

func (a *App) DeleteItem(w http.ResponseWriter, r *http.Request) {
	handler.DeleteItem(a.DB, w, r)
}

func (a *App) DisableItem(w http.ResponseWriter, r *http.Request) {
	handler.DisableItem(a.DB, w, r)
}

func (a *App) EnableItem(w http.ResponseWriter, r *http.Request) {
	handler.EnableItem(a.DB, w, r)
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
