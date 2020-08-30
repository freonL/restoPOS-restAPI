package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

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

func authHandler(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, pass, _ := r.BasicAuth()
		// fmt.Println(user, pass)
		// fmt.Println(r.Header)
		if r.Method == "GET" {
			next.ServeHTTP(w, r)

		} else {
			if user == "admin" && pass == "admin" {
				next.ServeHTTP(w, r)
			} else {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(401)
				w.Write([]byte("not Admin"))
			}
		}

	})
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
	a.Router.Use(authHandler)
	a.setRouters()
}

// Set all required routers
func (a *App) setRouters() {
	// Routing for handling the projects
	a.Get("/menus", a.GetMenus)
	a.Get("/menus/{id:[0-9]+}", a.GetMenu)
	a.Post("/menus", a.CreateMenu)
	a.Put("/menus/{id:[0-9]+}", a.UpdateMenu)
	a.Delete("/menus/{id:[0-9]+}", a.DeleteMenu)
	a.Put("/menus/{id:[0-9]+}/disable", a.DisableMenu)
	a.Put("/menus/{id:[0-9]+}/enable", a.EnableMenu)

	a.Get("/categories", a.GetCategories)
	a.Get("/categories/{id:[0-9]+}", a.GetCategory)
	a.Post("/categories", a.CreateCategories)
	a.Put("/categories/{id:[0-9]+}", a.UpdateCategories)
	a.Delete("/categories/{id:[0-9]+}", a.DeleteCategories)
	a.Put("/categories/{id:[0-9]+}/disable", a.DisableCategories)
	a.Put("/categories/{id:[0-9]+}/enable", a.EnableCategories)
}

// Wrap the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
	a.Router.HandleFunc(path, f).Methods("GET").Queries("limit", "{[0-9]*?}").Queries("offset", "{[0-9]*?}")
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

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
