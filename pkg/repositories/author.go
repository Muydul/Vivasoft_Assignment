package repositories

import (
	"book-crud/pkg/domain"
	"book-crud/pkg/models"

	"gorm.io/gorm"
)

// parent struct to implement interface binding
type authorRepo struct {
	db *gorm.DB
}

// interface binding
func AuthorDBInstance(d *gorm.DB) domain.IAuthorRepo {
	return &authorRepo{
		db: d,
	}
}

// all methods of interface are implemented
func (repo *authorRepo) GetAuthors(authorID uint) []models.AuthorDetail {
	var authors []models.AuthorDetail
	var err error

	if authorID != 0 {
		err = repo.db.Where("id = ?", authorID).Find(&authors).Error
	} else {
		err = repo.db.Find(&authors).Error
	}
	if err != nil {
		return []models.AuthorDetail{}
	}
	return authors
}

func (repo *authorRepo) CreateAuthor(author *models.AuthorDetail) error {
	if err := repo.db.Create(author).Error; err != nil {
		return err
	}
	return nil
}

func (repo *authorRepo) UpdateAuthor(author *models.AuthorDetail) error {
	if err := repo.db.Save(author).Error; err != nil {
		return err
	}
	return nil
}

func (repo *authorRepo) DeleteAuthor(authorID uint) error {
	var author models.AuthorDetail
	if err := repo.db.Where("id = ?", authorID).Delete(&author).Error; err != nil {
		return err
	}
	return nil
}
