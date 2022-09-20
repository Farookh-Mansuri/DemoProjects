package controllers

import "example.com/demo/services"

type BookController struct {
	BookService services.BookService
}
