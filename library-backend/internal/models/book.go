package models

import (
	"time"

	"gorm.io/gorm"
)

// Book GORM Model - Database table will be created from this
type Book struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Title       string         `json:"title" gorm:"type:varchar(255);not null;index" validate:"required,min=1,max=255"`
	Author      string         `json:"author" gorm:"type:varchar(255);not null;index" validate:"required,min=1,max=255"`
	Year        int            `json:"year" gorm:"not null;index;check:year >= 1000 AND year <= 2024" validate:"required,min=1000,max=2024"`
	ISBN        string         `json:"isbn,omitempty" gorm:"type:varchar(13);uniqueIndex" validate:"omitempty,len=13"`
	Description string         `json:"description,omitempty" gorm:"type:text"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"` // Soft delete
}

// TableName specifies the table name
func (Book) TableName() string {
	return "books"
}

// DTOs (Data Transfer Objects)
type CreateBookRequest struct {
	Title       string `json:"title" validate:"required,min=1,max=255"`
	Author      string `json:"author" validate:"required,min=1,max=255"`
	Year        int    `json:"year" validate:"required,min=1000,max=2024"`
	ISBN        string `json:"isbn,omitempty" validate:"omitempty,len=13"`
	Description string `json:"description,omitempty"`
}

type UpdateBookRequest struct {
	Title       *string `json:"title,omitempty" validate:"omitempty,min=1,max=255"`
	Author      *string `json:"author,omitempty" validate:"omitempty,min=1,max=255"`
	Year        *int    `json:"year,omitempty" validate:"omitempty,min=1000,max=2024"`
	ISBN        *string `json:"isbn,omitempty" validate:"omitempty,len=13"`
	Description *string `json:"description,omitempty"`
}

// BookFilter for search and filtering
type BookFilter struct {
	Title  string `form:"title" json:"title,omitempty"`
	Author string `form:"author" json:"author,omitempty"`
	Year   *int   `form:"year" json:"year,omitempty"`
	Limit  int    `form:"limit" json:"limit,omitempty"`
	Offset int    `form:"offset" json:"offset,omitempty"`
}

// API Response structures
type BookResponse struct {
	Success bool   `json:"success"`
	Data    *Book  `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}

type BooksResponse struct {
	Success bool   `json:"success"`
	Data    []Book `json:"data"`
	Total   int64  `json:"total"`
	Page    int    `json:"page,omitempty"`
	Limit   int    `json:"limit,omitempty"`
	Message string `json:"message,omitempty"`
}

// Convert DTO to Model
func (req *CreateBookRequest) ToModel() *Book {
	return &Book{
		Title:       req.Title,
		Author:      req.Author,
		Year:        req.Year,
		ISBN:        req.ISBN,
		Description: req.Description,
	}
}

// Apply updates to model
func (req *UpdateBookRequest) ApplyToModel(book *Book) {
	if req.Title != nil {
		book.Title = *req.Title
	}
	if req.Author != nil {
		book.Author = *req.Author
	}
	if req.Year != nil {
		book.Year = *req.Year
	}
	if req.ISBN != nil {
		book.ISBN = *req.ISBN
	}
	if req.Description != nil {
		book.Description = *req.Description
	}
}
