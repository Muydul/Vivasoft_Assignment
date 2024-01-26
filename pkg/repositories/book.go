package repositories

import (
	"book-crud/pkg/domain"
	"book-crud/pkg/models"
	"gorm.io/gorm"
)

// parent struct to implement interface binding
type bookRepo struct {
	db *gorm.DB
}

// interface binding
func BookDBInstance(d *gorm.DB) domain.IBookRepo {
	return &bookRepo{
		db: d,
	}
}

// all methods of interface are implemented
func (repo *bookRepo) GetBooks(bookID uint) []models.BookDetail {
	var book []models.BookDetail
	var err error

	if bookID != 0 {
		err = repo.db.Where("id = ?", bookID).Find(&book).Error
	} else {
		err = repo.db.Find(&book).Error
	}
	if err != nil {
		return []models.BookDetail{}
	}
	return book
}
func (repo *bookRepo) CreateBook(book *models.BookDetail) error {
	if err := repo.db.Create(book).Error; err != nil {
		return err
	}
	return nil
}

func (repo *bookRepo) UpdateBook(book *models.BookDetail) error {
	if err := repo.db.Save(book).Error; err != nil {
		return err
	}
	return nil
}
func (repo *bookRepo) DeleteBook(bookID uint) error {
	var Book models.BookDetail
	if err := repo.db.Where("id = ?", bookID).Delete(&Book).Error; err != nil {
		return err
	}
	return nil
}
