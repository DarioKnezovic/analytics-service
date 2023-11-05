package database

import (
	"fmt"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
	"time"
)

// Migration represents a record in the migrations table
type Migration struct {
	ID        int
	Name      string
	CreatedAt time.Time
}

func PerformAutoMigrations(db *gorm.DB) error {
	migrationsDir := "./pkg/database/migrations/"
	migrationFiles, err := ioutil.ReadDir(migrationsDir)
	if err != nil {
		log.Fatal("Failed to read migration files:", err)
		return err
	}

	executedMigrations, err := fetchAllMigrations(db)
	if err != nil {
		log.Fatal("Failed to read executed migrations:", err)
		return err
	}

	// Start a new database transaction
	tx := db.Begin()
	if tx.Error != nil {
		log.Fatal("Failed to start a transaction:", tx.Error)
		return tx.Error
	}

	// Execute each migration file with a .sql extension
	for _, file := range migrationFiles {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".sql") {
			migrationFilePath := filepath.Join(migrationsDir, file.Name())
			if !checkIfMigrationExecuted(executedMigrations, file.Name()) {
				err := executeMigration(tx, migrationFilePath)
				if err != nil {
					// Rollback the transaction and return the error
					tx.Rollback()
					log.Fatal("Failed to execute migration:", err)
					return err
				}
				err = saveMigration(tx, file.Name())
				if err != nil {
					// Rollback the transaction and return the error
					tx.Rollback()
					log.Fatal("Failed to save migration:", err)
					return err
				}
			}
		}
	}

	// Commit the transaction if all migrations succeed
	if err := tx.Commit().Error; err != nil {
		// If there's an error during commit, rollback the transaction
		tx.Rollback()
		log.Fatal("Failed to commit transaction:", err)
		return err
	}

	log.Print("Auto migrations have been successfully finished")

	return nil
}

// CheckIfMigrationExecuted checks if the target migration (fileName) has been executed
func checkIfMigrationExecuted(executedMigrations []Migration, fileName string) bool {
	for _, migration := range executedMigrations {
		if migration.Name == fileName {
			return true // The migration has been executed
		}
	}
	return false // The migration has not been executed
}

func saveMigration(tx *gorm.DB, fileName string) error {
	query := fmt.Sprintf("INSERT INTO migrations (name) VALUES ('%s');", fileName)
	err := tx.Exec(query).Error
	return err
}

// FetchAllMigrations retrieves all migration records from the migrations table
func fetchAllMigrations(db *gorm.DB) ([]Migration, error) {
	var migrations []Migration
	if err := db.Table("migrations").Find(&migrations).Error; err != nil {
		log.Println("Error fetching migrations:", err)
		return nil, err
	}
	return migrations, nil
}

func executeMigration(tx *gorm.DB, migrationFilePath string) error {
	// Read the SQL content from the migration file
	content, err := ioutil.ReadFile(migrationFilePath)
	if err != nil {
		return err
	}

	// Execute the SQL query within the transaction
	result := tx.Exec(string(content))

	if result.Error != nil {
		return result.Error
	}

	return nil
}
