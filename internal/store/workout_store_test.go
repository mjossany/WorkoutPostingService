package store

import (
	"database/sql"
	"testing"
)

func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("pgx", "host=localhost user=postgres password=postgres dbname=postgres port=5433 sslmode=disable")
	if err != nil {
		t.Fatalf("opening test db: %v", err)
	}

	err = Migrate(db, "../../migrations")
	if err != nil {
		t.Fatalf("migrating test db error: %w", err)
	}

	_, err = db.Exec(`TRUNCATE workouts, workout_entries CASCADE`)
	if err != nil {
		t.Fatalf("truncating tables error: %w", err)
	}

	return db
}
