package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"my/crm-golang/internal/models/contact"
	"my/crm-golang/internal/services/contacts"
	contactsrepo "my/crm-golang/internal/storage/postgres/contacts"
)

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	repo := contactsrepo.New(db)

	service := contacts.NewService(repo)

	newContact := &contact.Contact{
		Name:     "Max",
		Phone:    "123",
		Telegram: "@max",
		Birthday: "1990-05-15",
	}

	if err := service.Create(newContact); err != nil {
		fmt.Printf("Failed to create contact: %v", err)
	}
	fmt.Println("Contact created successfully")

	contactModel, err := service.GetByName("Alice Smith")
	if err != nil {
		fmt.Printf("Failed to get contact: %v", err)
	}
	fmt.Printf("Contact retrieved: %+v\n", contactModel)

	updateData := &contact.ContactUpdateData{
		//Phone:    "4566",
		//Telegram: "@new_max_max",
	}
	if err := service.Update("Max", updateData); err != nil {
		fmt.Printf("Failed to update contact: %v", err)
	}
	fmt.Println("Contact updated successfully")

	if err := service.DeleteByName("Max"); err != nil {
		log.Fatalf("Failed to delete contact: %v", err)
	}
	fmt.Println("Contact deleted successfully")
}
