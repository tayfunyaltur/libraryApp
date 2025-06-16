package service

import (
	"errors"
	"library-backend/internal/models"
	"library-backend/pkg/database"

	"gorm.io/gorm"
)

type BookService struct {
	db *database.Database
}

func NewBookService(db *database.Database) *BookService {
	return &BookService{db: db}
}

func (s *BookService) GetAllBooks(filter *models.BookFilter) (*models.BooksResponse, error) {
	var books []models.Book
	var total int64

	query := s.db.Model(&models.Book{})

	// Apply filters
	if filter.Title != "" {
		query = query.Where("title ILIKE ?", "%"+filter.Title+"%")
	}
	if filter.Author != "" {
		query = query.Where("author ILIKE ?", "%"+filter.Author+"%")
	}
	if filter.Year != nil {
		query = query.Where("year = ?", *filter.Year)
	}

	// Get total count
	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}

	// Apply pagination
	if filter.Limit > 0 {
		query = query.Limit(filter.Limit)
	}
	if filter.Offset > 0 {
		query = query.Offset(filter.Offset)
	}

	// Execute query
	if err := query.Order("created_at DESC").Find(&books).Error; err != nil {
		return nil, err
	}

	return &models.BooksResponse{
		Success: true,
		Data:    books,
		Total:   total,
		Page:    filter.Offset/filter.Limit + 1,
		Limit:   filter.Limit,
	}, nil
}

func (s *BookService) GetBookByID(id uint) (*models.Book, error) {
	var book models.Book

	if err := s.db.First(&book, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("book not found")
		}
		return nil, err
	}

	return &book, nil
}

func (s *BookService) CreateBook(req *models.CreateBookRequest) (*models.Book, error) {
	book := req.ToModel()

	if err := s.db.Create(book).Error; err != nil {
		return nil, err
	}

	return book, nil
}

func (s *BookService) UpdateBook(id uint, req *models.UpdateBookRequest) (*models.Book, error) {
	var book models.Book

	// Find existing book
	if err := s.db.First(&book, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("book not found")
		}
		return nil, err
	}

	// Apply updates
	req.ApplyToModel(&book)

	// Save changes
	if err := s.db.Save(&book).Error; err != nil {
		return nil, err
	}

	return &book, nil
}

func (s *BookService) DeleteBook(id uint) error {
	result := s.db.Delete(&models.Book{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("book not found")
	}

	return nil
}

func (s *BookService) SearchBooks(query string) ([]models.Book, error) {
	var books []models.Book

	err := s.db.Where("title ILIKE ? OR author ILIKE ? OR description ILIKE ?",
		"%"+query+"%", "%"+query+"%", "%"+query+"%").
		Order("created_at DESC").
		Find(&books).Error

	return books, err
}
