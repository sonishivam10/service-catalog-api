package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://user:password@db:5432/service_catalog?sslmode=disable"
	}

	var db *sqlx.DB
	var err error

	// Retry until DB is ready
	for i := 0; i < 10; i++ {
		db, err = sqlx.Connect("postgres", dsn)
		if err == nil {
			break
		}
		log.Println("⏳ Waiting for DB...")
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		log.Fatal("❌ Failed to connect to DB:", err)
	}

	// Upsert services
	for i := 1; i <= 10; i++ {
		name := fmt.Sprintf("Service %02d", i)
		description := fmt.Sprintf("Description for %s", name)

		// Check if service already exists
		var count int
		err := db.Get(&count, `SELECT COUNT(*) FROM services WHERE name = $1`, name)
		if err != nil {
			log.Fatalf("Failed to check existence of %s: %v", name, err)
		}

		var serviceID uuid.UUID
		if count == 0 {
			// Insert new service
			serviceID = uuid.New()
			_, err = db.Exec(`INSERT INTO services (id, name, description) VALUES ($1, $2, $3)`, serviceID, name, description)
			if err != nil {
				log.Fatalf("Failed to insert %s: %v", name, err)
			}
			log.Printf("✅ Inserted %s", name)
		} else {
			// Get existing service ID
			err = db.Get(&serviceID, `SELECT id FROM services WHERE name = $1`, name)
			if err != nil {
				log.Fatalf("Failed to fetch ID for %s: %v", name, err)
			}
		}

		// Insert versions only if they don’t exist
		versions := []string{"v1.0.0", "v1.1.0", "v2.0.0"}
		for _, version := range versions {
			var versionCount int
			err = db.Get(&versionCount, `SELECT COUNT(*) FROM versions WHERE service_id = $1 AND version = $2`, serviceID, version)
			if err != nil {
				log.Fatalf("Check version %s for %s failed: %v", version, name, err)
			}

			if versionCount == 0 {
				_, err = db.Exec(`INSERT INTO versions (id, service_id, version) VALUES ($1, $2, $3)`, uuid.New(), serviceID, version)
				if err != nil {
					log.Fatalf("Failed to insert version %s for %s: %v", version, name, err)
				}
				log.Printf("   ↳ Added version %s", version)
			}
		}
	}

	log.Println("✅ All seed data inserted.")
}
