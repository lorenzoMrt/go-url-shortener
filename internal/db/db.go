package db

import "database/sql"

func CreateTable(db *sql.DB) error {
	var query = `CREATE TABLE IF NOT EXISTS urls (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    short_url TEXT NOT NULL,
    original_url TEXT NOT NULL
);`
	_, err := db.Exec(query)
	return err
}

func StoreUrl(db *sql.DB, shortUrl string, originalUrl string) error {
	query := `INSERT INTO urls (short_url, original_url) VALUES (?, ?)`
	_, err := db.Exec(query, shortUrl, originalUrl)
	return err
}

func GetOriginalUrl(db *sql.DB, shortUrl string) (string, error) {
	var originalUrl string
	query := `SELECT original_url FROM urls WHERE short_url = ? LIMIT 1`
	err := db.QueryRow(query, shortUrl).Scan(&originalUrl)
	if err != nil {
		return "", err
	}
	return originalUrl, nil
}
