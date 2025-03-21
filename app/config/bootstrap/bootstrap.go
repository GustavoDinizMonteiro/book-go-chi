package bootstrap

import (
	"database/sql"
	"fmt"
	"net/url"
	"os"

	_ "github.com/lib/pq"
)

func InitDB() (*sql.DB, error) {
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		return nil, fmt.Errorf("DB_URL environment variable is required")
	}

	parsedURL, err := url.Parse(dbURL)
	if err != nil {
		return nil, fmt.Errorf("error parsing DB_URL: %v", err)
	}

	host := parsedURL.Hostname()
	port := parsedURL.Port()
	user := parsedURL.User.Username()
	password, _ := parsedURL.User.Password()
	dbname := parsedURL.Path[1:]

	queryParams := parsedURL.Query()
	sslmode := queryParams.Get("sslmode")
	if sslmode == "" {
		sslmode = "require"
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host,
		port,
		user,
		password,
		dbname,
		sslmode,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
