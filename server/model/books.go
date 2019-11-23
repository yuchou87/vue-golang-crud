package model

import (
	"database/sql"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Status bool   `json:"status"`
}

func (b *Book) GetBook(db *sql.DB) error {
	return db.QueryRow("SELECT title, author, status FROM books WHERE id=$1", b.ID).Scan(&b.Title, &b.Author, &b.Status)
}

func (b *Book) UpdateBook(db *sql.DB) error {
	_, err := db.Exec("UPDATE books SET title=$1, author=$2, status=$3 WHERE id=$4", b.Title, b.Author, b.Status, b.ID)
	return err
}

func (b *Book) DeleteBook(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM books WHERE id=$1", b.ID)
	return err
}

func (b *Book) CreateBook(db *sql.DB) error {
	err := db.QueryRow("INSERT INTO books(title, author, status) VALUES($1, $2, $3) RETURNING id", b.Title, b.Author, b.Status).Scan(&b.ID)
	if err != nil {
		return err
	}
	return nil
}

func GetBooks(db *sql.DB, start, count int) ([]Book, error) {
	rows, err := db.Query("SELECT id, title, author, status FROM books LIMIT $1 OFFSET $2", count, start)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := []Book{}
	for rows.Next() {
		var b Book
		if err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.Status); err != nil {
			return nil, err
		}
		books = append(books, b)
	}
	return books, nil
}
