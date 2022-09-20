package services

import "example.com/demo/models"

type BookService interface {
	CreateBook(*models.Book) error
	GetBook(*string) (*models.Book, error)
	GetAllBooks() ([]*models.Book, error)
	UpdateBook(*models.Book) error
}
