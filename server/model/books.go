package model

import (
	"database/sql"
	"errors"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Read   bool   `json:"read"`
}

func (b *Book) getBooks(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (b *Book) updateBook(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (b *Book) deleteBook(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (b *Book) createBook(db *sql.DB) error {
	return errors.New("Not implemented")
}

func getBooks(db *sql.DB, start, count int) ([]Book, error) {
	return nil, errors.New("Not implemented")
}
