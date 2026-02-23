package main

import (
	"context"
	"fmt"
	"os"
	"service/discount/api/database/migrate"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	pgxConnAuthDb  *pgxpool.Pool
	dbUsersService string
	serverName     string
)

func init() {

	serverName = os.Getenv("SERVER_NAME")
	dbUsersService = os.Getenv("DB_SERVER_AUTH")

	if serverName == "" || dbUsersService == "" {
		exitf("SERVER_NAME and DB_SERVER_AUTH must be set")
	}
}

func main() {
	dbConnection()
	defer pgxConnAuthDb.Close()

	ctx := context.Background()

	// Buat table migrations kalau belum ada
	_, err := pgxConnAuthDb.Exec(ctx, migrate.CreateTableMigration())
	if err != nil {
		exitf("Failed create migrations table: %v", err)
	}

	// ambil daftar migrations
	migrations := migrate.GetMigrations()

	for _, m := range migrations {
		if isMigrated(ctx, m.Name) {
			fmt.Println("Already migrated:", m.Name)
			continue
		}

		fmt.Println("Migrating:", m.Name)
		_, err := pgxConnAuthDb.Exec(ctx, m.SQL)
		if err != nil {
			exitf("Migration failed (%s): %v", m.Name, err)
		}

		_, err = pgxConnAuthDb.Exec(ctx, "INSERT INTO migrations (name) VALUES ($1)", m.Name)
		if err != nil {
			exitf("Failed to insert migration record (%s): %v", m.Name, err)
		}

		fmt.Println("Migration success:", m.Name)
	}

	fmt.Println("All migrations done!")
}

func dbConnection() {
	var err error
	cfgAuthDb, err := pgxpool.ParseConfig(dbUsersService + " application_name=" + serverName)
	if err != nil {
		exitf("Unable to create db pool config: %v", err)
	}
	cfgAuthDb.MaxConns = 1000
	cfgAuthDb.MaxConnLifetime = 5 * 60 * 1000000000 // 5 minutes
	cfgAuthDb.MaxConnIdleTime = 2 * 60 * 1000000000 // 2 minutes

	pgxConnAuthDb, err = pgxpool.NewWithConfig(context.Background(), cfgAuthDb)
	if err != nil {
		exitf("Unable to connect to database: %v", err)
	}
}

func isMigrated(ctx context.Context, name string) bool {
	var exists bool
	err := pgxConnAuthDb.QueryRow(ctx, `
		SELECT EXISTS (
			SELECT 1 FROM migrations WHERE name = $1
		)
	`, name).Scan(&exists)
	if err != nil {
		exitf("Failed to check migration exists: %v", err)
	}
	return exists
}

func exitf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}
