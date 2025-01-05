package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"my/crm-golang/internal/api/handlers/contacts_create"
	"my/crm-golang/internal/api/handlers/contacts_delete"
	"my/crm-golang/internal/api/handlers/contacts_get_all"
	"my/crm-golang/internal/api/handlers/contacts_get_last_names"
	"my/crm-golang/internal/api/handlers/contacts_get_one"
	"my/crm-golang/internal/api/handlers/contacts_get_similar"
	"my/crm-golang/internal/api/handlers/contacts_update"
	"my/crm-golang/internal/api/handlers/info"
	"my/crm-golang/internal/services/contacts"
	contactsrepo "my/crm-golang/internal/storage/postgres/contacts"
)

type App struct {
	chiRouter      *chi.Mux
	contactService *contacts.Service
}

func main() {
	fmt.Print("Hello!")

	app := NewApp()
	app.chiRouter.Get("/api/v2/__info/", info.New().Handle)
	app.chiRouter.Get("/api/v2/contacts/get/", contacts_get_all.New(app.contactService).Handle)
	app.chiRouter.Post("/api/v2/contacts/create/", contacts_create.New(app.contactService).Handle)
	app.chiRouter.Get("/api/v2/contacts/get_lasts/", contacts_get_last_names.New(app.contactService).Handle)
	app.chiRouter.Get("/api/v2/contacts/{name}/get/", contacts_get_one.New(app.contactService).Handle)
	app.chiRouter.Get("/api/v2/contacts/{name}/get_similar/", contacts_get_similar.New(app.contactService).Handle)
	app.chiRouter.Put("/api/v2/contacts/{name}/update/", contacts_update.New(app.contactService).Handle)
	app.chiRouter.Delete("/api/v2/contacts/{name}/delete/", contacts_delete.New(app.contactService).Handle)

	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", app.chiRouter); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

func NewApp() *App {
	// database
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	repository := contactsrepo.New(db)

	// services
	contactService := contacts.NewService(repository)

	// chi router
	chiRouter := chi.NewRouter()
	chiRouter.Use(middleware.Logger)

	// app
	return &App{
		contactService: contactService,
		chiRouter:      chiRouter,
	}
}

//func getRouter() http.Handler {
//	r := chi.NewRouter()
//	r.Use(AdminOnly)
//	r.Get("/", adminIndex)
//	r.Get("/accounts", adminListAccounts)
//	return r
//}
