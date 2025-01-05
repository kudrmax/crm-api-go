package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"my/crm-golang/internal/models/contact"
	"my/crm-golang/internal/storage/postgres/contacts"
)

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	repo := contacts.New(db)

	newContact := &contact.Contact{
		Name:     "Alice Smith",
		Phone:    "1234567890",
		Telegram: "@alicesmith",
		Birthday: "1990-05-15",
	}

	if err := repo.Create(newContact); err != nil {
		log.Fatalf("Failed to create contact: %v", err)
	}
	fmt.Println("Contact created successfully")

	contactModel, err := repo.GetByName("Alice Smith")
	if err != nil {
		log.Fatalf("Failed to get contact: %v", err)
	}
	fmt.Printf("Contact retrieved: %+v\n", contactModel)

	updateData := &contact.ContactUpdateData{
		Phone:    "0987654321",
		Telegram: "@alice_new",
	}
	if err := repo.Update(contactModel, updateData); err != nil {
		log.Fatalf("Failed to update contact: %v", err)
	}
	fmt.Println("Contact updated successfully")

	if err := repo.DeleteByName("Alice Smith"); err != nil {
		log.Fatalf("Failed to delete contact: %v", err)
	}
	fmt.Println("Contact deleted successfully")
}
