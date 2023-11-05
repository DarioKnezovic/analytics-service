package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

const (
	MigrationsDirectory = "./pkg/database/migrations/"
)

func main() {
	// Check if the migration name is provided as a command-line argument
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run create_migration.go <migration_name>")
		os.Exit(1)
	}
	fmt.Printf("Creating a new migration... \n")

	// Get the provided migration name from the command line
	migrationName := os.Args[1]

	// Determine the next available migration number
	nextMigrationNumber := getNextMigrationNumber()

	// Create the migration filename
	migrationFileName := fmt.Sprintf("%03d_%s.sql", nextMigrationNumber, migrationName)

	// Create the migration file
	if err := createMigrationFile(migrationFileName); err != nil {
		fmt.Println("Failed to create migration file:", err)
		os.Exit(1)
	}

	fmt.Printf("Created migration file: %s\n", migrationFileName)
}

// getNextMigrationNumber returns the next available migration number
func getNextMigrationNumber() int {
	files, err := filepath.Glob(filepath.Join(MigrationsDirectory, "*.sql"))
	if err != nil {
		fmt.Println("Error while reading migration files:", err)
		os.Exit(1)
	}

	return len(files) + 1
}

// createMigrationFile creates a new migration SQL file with the given name
func createMigrationFile(fileName string) error {
	filePath := filepath.Join(MigrationsDirectory, fileName)

	// Create the migration file with a timestamp-based comment
	fileContent := fmt.Sprintf("-- Migration created at %s\n\n", time.Now().Format(time.RFC3339))

	// Write the content to the file
	return ioutil.WriteFile(filePath, []byte(fileContent), 0644)
}
