package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"

	_ "github.com/lib/pq"
)

// Путь к папке с миграциями
var migrationsDir = "./migrations/postgresql/master"

func main() {
	// Настройка подключения к базе данных
	dsn := "host=localhost user=postgres password=postgres dbname=crm_golang port=5432 sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Чтение файлов в директории
	files, err := os.ReadDir(migrationsDir)
	if err != nil {
		log.Fatalf("Failed to read migrations directory: %v", err)
	}
	if len(files) == 0 {
		log.Fatalf("No migrations found in directory %s", migrationsDir)
	}

	// Отбор и сортировка файлов с расширением .sql
	var sqlFiles []string
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".sql" {
			sqlFiles = append(sqlFiles, file.Name())
		}
	}
	sort.Strings(sqlFiles)

	// Выполнение миграций
	for _, fileName := range sqlFiles {
		filePath := filepath.Join(migrationsDir, fileName)
		fmt.Printf("Applying migration: %s\n", fileName)

		// Чтение содержимого файла
		content, err := os.ReadFile(filePath)
		if err != nil {
			log.Fatalf("Failed to read migration file %s: %v", fileName, err)
		}

		// Выполнение SQL запроса
		if _, err := db.Exec(string(content)); err != nil {
			log.Fatalf("Failed to execute migration %s: %v", fileName, err)
		}

		fmt.Printf("Migration applied: %s\n", fileName)
	}

	fmt.Println("All migrations applied successfully!")
}
