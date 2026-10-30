package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"example.com/pz9-auth/internal/http/handlers"
	"example.com/pz9-auth/internal/platform/config"
	"example.com/pz9-auth/internal/repo"
)

func main() {
	cfg := config.Load()
	if cfg.DB_DSN == "" {
		log.Fatal("DB_DSN is not set")
	}

	db, err := repo.Open(cfg.DB_DSN)
	if err != nil {
		log.Fatal("db connect: ", err)
	}
	_ = db.Exec("SET timezone TO 'UTC'").Error

	users := repo.NewUserRepo(db)
	if err := users.AutoMigrate(); err != nil {
		log.Fatal("migrate: ", err)
	}
	log.Println("migration ok: table users")

	auth := handlers.NewAuthHandler(users, cfg.BcryptCost)

	r := chi.NewRouter()
	r.Use(middleware.Logger) // логирует метод, путь и статус (без тела запроса и паролей)
	r.Post("/auth/register", auth.Register)
	r.Post("/auth/login", auth.Login)

	log.Printf("listening on %s (bcrypt cost=%d)", cfg.Addr, cfg.BcryptCost)
	log.Fatal(http.ListenAndServe(cfg.Addr, r))
}
