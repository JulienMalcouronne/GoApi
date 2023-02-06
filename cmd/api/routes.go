package main

import (
	"net/http"
	"vue-api/internal/data"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	mux.Get("/users/login", app.Login)
	mux.Post("/users/login", app.Login)
	mux.Get("/users/all", func(w http.ResponseWriter, r *http.Request) {
		var users data.User
		all, err := users.GetAll()
		if err != nil {
			app.errorLog.Println(err)
			return
		}
		app.writeJSON(w, http.StatusOK, all)
	})

	mux.Get("/users/add", func(w http.ResponseWriter, r *http.Request) {
		var u = data.User{
			Email:     "you@therezzzzaa.com",
			FirstName: "You",
			LastName:  "There",
			Password:  "password",
		}

		app.infoLog.Println("Adding user...")

		id, err := app.models.User.Insert(u)
		if err != nil {
			app.errorLog.Println(err)
			app.errorJSON(w, err, http.StatusForbidden)
			return
		}

		app.infoLog.Println("Got back id of", id)
		newUser, _ := app.models.User.GetOne(id)
		app.writeJSON(w, http.StatusOK, newUser)
	})

	return mux
}
