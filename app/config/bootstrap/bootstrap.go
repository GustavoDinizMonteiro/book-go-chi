package bootstrap

import (
	"database/sql"
	"fmt"
	"net/url"
	"os"

	_ "github.com/lib/pq"
)

func InitDB() (*sql.DB, error) {
	// Pega a URL de conexão do banco de dados
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		return nil, fmt.Errorf("DB_URL environment variable is required")
	}

	// Analisa a URL
	parsedURL, err := url.Parse(dbURL)
	if err != nil {
		return nil, fmt.Errorf("error parsing DB_URL: %v", err)
	}

	// Extrai os componentes da URL
	host := parsedURL.Hostname()
	port := parsedURL.Port()
	user := parsedURL.User.Username()
	password, _ := parsedURL.User.Password()
	dbname := parsedURL.Path[1:] // Remove o "/" inicial

	// Formata o DSN
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbname,
	)

	// Conecta ao banco de dados
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// Verifica se a conexão foi bem-sucedida
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}