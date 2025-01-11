package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"my/crm-golang/internal/api/handlers/contact_logs_create"
	"my/crm-golang/internal/api/handlers/contact_logs_create_empty"
	"my/crm-golang/internal/api/handlers/contact_logs_get_all_list"
	"my/crm-golang/internal/api/handlers/contact_logs_get_one"
	"my/crm-golang/internal/api/handlers/contacts_create"
	"my/crm-golang/internal/api/handlers/contacts_delete"
	"my/crm-golang/internal/api/handlers/contacts_get_all"
	"my/crm-golang/internal/api/handlers/contacts_get_last_names"
	"my/crm-golang/internal/api/handlers/contacts_get_one"
	"my/crm-golang/internal/api/handlers/contacts_get_similar"
	"my/crm-golang/internal/api/handlers/contacts_update"
	"my/crm-golang/internal/api/handlers/info"
	"my/crm-golang/internal/api/handlers/not_implemented"
	"my/crm-golang/internal/services/contact_logs"
	"my/crm-golang/internal/services/contacts"
	logsrepo "my/crm-golang/internal/storage/postgres/contact_logs"
	contactsrepo "my/crm-golang/internal/storage/postgres/contacts"
)

func main() {
	log.Print("Started!")

	app := NewApp()
	log.Print("App was created!")

	app.chiRouter.Get("/__info/", info.New().Handle)

	app.chiRouter.Get("/api/v2/contacts/get/", contacts_get_all.New(app.contactService).Handle)
	app.chiRouter.Post("/api/v2/contacts/create/", contacts_create.New(app.contactService).Handle)
	app.chiRouter.Get("/api/v2/contacts/get_lasts/", contacts_get_last_names.New(app.contactService).Handle)
	app.chiRouter.Get("/api/v2/contacts/{name}/get/", contacts_get_one.New(app.contactService).Handle)
	app.chiRouter.Get("/api/v2/contacts/{name}/get_similar/", contacts_get_similar.New(app.contactService).Handle)
	app.chiRouter.Put("/api/v2/contacts/{name}/update/", contacts_update.New(app.contactService).Handle)
	app.chiRouter.Delete("/api/v2/contacts/{name}/delete/", contacts_delete.New(app.contactService).Handle)

	app.chiRouter.Post("/api/v2/contacts/{name}/logs/create/", contact_logs_create.New(app.contactLogService, app.contactService).Handle)
	app.chiRouter.Post("/api/v2/contacts/{name}/logs/create/empty/", contact_logs_create_empty.New(app.contactLogService, app.contactService).Handle)
	app.chiRouter.Get("/api/v2/contacts/{name}/logs/get_all/list/", contact_logs_get_all_list.New(app.contactLogService, app.contactService).Handle)
	app.chiRouter.Get("/api/v2/logs/last_logs/", not_implemented.New().Handle)
	app.chiRouter.Get("/api/v2/logs/{log_id}/get/", contact_logs_get_one.New().Handle)
	app.chiRouter.Put("/api/v2/logs/{log_id}/update/", not_implemented.New().Handle)    // contact_logs_update
	app.chiRouter.Delete("/api/v2/logs/{log_id}/delete/", not_implemented.New().Handle) // contact_logs_delete

	app.chiRouter.Get("/api/v2/stats/count_of_interactions/", not_implemented.New().Handle)
	app.chiRouter.Get("/api/v2/stats/count_of_interactions/{name}/", not_implemented.New().Handle)
	app.chiRouter.Get("/api/v2/stats/days_count_since_last_interaction/", not_implemented.New().Handle)
	app.chiRouter.Get("/api/v2/stats/days_count_since_last_interaction/{name}", not_implemented.New().Handle)

	log.Print("Router set up!")

	log.Println("Starting server on :8000...")
	if err := http.ListenAndServe(":8000", app.chiRouter); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

type App struct {
	chiRouter         *chi.Mux
	contactService    *contacts.Service
	contactLogService *contact_logs.Service
}

func NewApp() *App {
	// database
	dsn := "host=localhost user=dev_u password=dev_p dbname=dev_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	repositoryContacts := contactsrepo.New(db)
	repositoryContactLogs := logsrepo.New(db)
	log.Print("Connected to database!")

	// services
	contactService := contacts.NewService(repositoryContacts)
	contactLogService := contact_logs.NewService(repositoryContactLogs)

	// chi router
	chiRouter := chi.NewRouter()
	chiRouter.Use(middleware.Logger)

	// app
	return &App{
		contactService:    contactService,
		contactLogService: contactLogService,
		chiRouter:         chiRouter,
	}
}

//func getRouter() http.Handler {
//	r := chi.NewRouter()
//	r.Use(AdminOnly)
//	r.Get("/", adminIndex)
//	r.Get("/accounts", adminListAccounts)
//	return r
//}
