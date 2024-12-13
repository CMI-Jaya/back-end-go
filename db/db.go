package db

import (
	"context"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"go-project/config"

	_ "github.com/jackc/pgx/v5/stdlib" // Import driver PostgreSQL
)

var DB *sql.DB

func ConnectDB(cfg *config.Config) {
	var err error
	DB, err = sql.Open("pgx", cfg.GetDBConnectionString())
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Set database connection pool parameters
	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(25)
	DB.SetConnMaxLifetime(5 * time.Minute)

	// Test connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = DB.PingContext(ctx); err != nil {
		log.Fatalf("Database connection error: %v", err)
	}

	log.Println("Connected to PostgreSQL database successfully!")
}

func RunMigrations(migrationsDir string) {

	files, err := ioutil.ReadDir(migrationsDir)
	if err != nil {
		log.Fatalf("Gagal untuk membaca migrations %v", err)
	}

	// run sql
	for _, file := range files {
		if file.IsDir() || filepath.Ext(file.Name()) != ".sql" {
			continue
		}

		filepath := filepath.Join(migrationsDir, file.Name())
		content, err := os.ReadFile(filepath)
		if err != nil {
			log.Fatalf("Gagal membaca migration file %v", file.Name(), err)
		}

		fmt.Println("Running Migrations %s: %v", file.Name(), err)
		_, err = DB.Exec(string(content))
		if err != nil {
			log.Fatalf("Failed execute migration %s: %v", file.Name(), err)
		}
		fmt.Println("Migration success !", file.Name())
	}
}

func SaveUser(email, hashedPassword, role string) error {
	query := "INSERT INTO users (email, password, role) VALUES ($1, $2, $3)"
	_, err := DB.Exec(query, email, hashedPassword, role)
	if err != nil {
		log.Printf("Error saat menyimpan user: %v", err)
		return err
	}
	return nil
}

func GetDB() *sql.DB {
	return DB
}
