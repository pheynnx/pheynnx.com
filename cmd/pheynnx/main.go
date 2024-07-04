package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"

	"github.com/pheynnx/pheynnx.com/internal/cache"
	"github.com/pheynnx/pheynnx.com/internal/controller/blog"
	"github.com/pheynnx/pheynnx.com/internal/database"
	"github.com/pheynnx/pheynnx.com/internal/orbit"
)

func main() {
	// load .env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	// initialize sqlx pool
	db, err := database.NewDBPool()
	if err != nil {
		log.Fatal(err)
	}

	// initialize post cache for templates to pull from
	// currently only shaving of filesystem io call and markdown parsing times
	// if data is moved back to the cloud, caching the posts in memory will eliminate waiting for networked io
	pc, err := cache.InitCache(db)
	if err != nil {
		log.Fatal(err)
	}

	// create top level router
	r := chi.NewRouter()

	// root level middleware stack
	r.Use(chiMiddleware.RealIP)
	// r.Use(middleware.Logger)
	r.Use(chiMiddleware.Recoverer)

	// serve static files
	// r.Get("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "web/root/favicon.ico")
	// })
	r.Get("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/root/robots.txt")
	})
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	// app routers
	r.Group(func(r chi.Router) {

		// I would like to keep backend analytics of site traffic in its own thread
		// r.Use(middleware.Metrics(db))

		r.Mount("/", blog.Routes(pc))
	})

	// start server
	orbit.Launch(r)
}
