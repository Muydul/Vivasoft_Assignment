package domain

import (
	"book-crud/pkg/models"
	"book-crud/pkg/types"
)

// for database repository operation (call from service)
type IAuthorRepo interface {
	GetAuthors(authorID uint) []models.AuthorDetail
	CreateAuthor(author *models.AuthorDetail) error
	UpdateAuthor(author *models.AuthorDetail) error
	DeleteAuthor(authorID uint) error
}

// for service operation (response to controller | call from controller)
type IAuthorService interface {
	GetAuthors(authorID uint) ([]types.AuthorRequest, error)
	CreateAuthor(author *models.AuthorDetail) error
	UpdateAuthor(author *models.AuthorDetail) error
	DeleteAuthor(authorID uint) error
}
