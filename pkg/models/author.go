package models

// struct for database schema
type AuthorDetail struct {
	ID            uint `gorm:"primaryKey;autoIncrement"`
	AuthorName    string
	AuthorAddress string
	Phone         string
}
