package services

import (
	"book-crud/pkg/domain"
	"book-crud/pkg/models"
	"book-crud/pkg/types"
	"errors"
)

// parent struct to implement interface binding
type authorService struct {
	repo domain.IAuthorRepo
}

// interface binding
func AuthorServiceInstance(authorRepo domain.IAuthorRepo) domain.IAuthorService {
	return &authorService{
		repo: authorRepo,
	}
}

// all methods of interface are implemented
func (service *authorService) GetAuthors(authorID uint) ([]types.AuthorRequest, error) {
	var allAuthors []types.AuthorRequest
	authors := service.repo.GetAuthors(authorID)
	if len(authors) == 0 {
		return nil, errors.New("No author found")
	}
	for _, val := range authors {
		allAuthors = append(allAuthors, types.AuthorRequest{
			ID:            val.ID,
			AuthorName:    val.AuthorName,
			AuthorAddress: val.AuthorAddress,
			Phone:         val.Phone,
		})
	}
	return allAuthors, nil
}

func (service *authorService) CreateAuthor(author *models.AuthorDetail) error {
	if err := service.repo.CreateAuthor(author); err != nil {
		return errors.New("AuthorDetail was not created")
	}
	return nil
}

func (service *authorService) UpdateAuthor(author *models.AuthorDetail) error {
	if err := service.repo.UpdateAuthor(author); err != nil {
		return errors.New("AuthorDetail update was unsuccessful")
	}
	return nil
}

func (service *authorService) DeleteAuthor(authorID uint) error {
	if err := service.repo.DeleteAuthor(authorID); err != nil {
		return errors.New("AuthorDetail deletion was unsuccessful")
	}
	return nil
}
