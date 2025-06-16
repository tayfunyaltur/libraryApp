package handlers

import (
	"library-backend/internal/models"
	"library-backend/internal/service"
	"library-backend/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type BookHandler struct {
	service   *service.BookService
	validator *validator.Validate
}

func NewBookHandler(service *service.BookService) *BookHandler {
	return &BookHandler{
		service:   service,
		validator: validator.New(),
	}
}

// GetBooks retrieves all books with pagination and filtering
// @Summary      Get all books
// @Description  Get a list of all books with optional filtering and pagination
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        title    query     string  false  "Filter by title"
// @Param        author   query     string  false  "Filter by author"
// @Param        year     query     int     false  "Filter by year"
// @Param        limit    query     int     false  "Number of items per page (default 10, max 100)"
// @Param        offset   query     int     false  "Number of items to skip (default 0)"
// @Success      200      {object}  models.BooksResponse
// @Failure      400      {object}  models.ErrorResponse
// @Failure      500      {object}  models.ErrorResponse
// @Router       /books [get]
func (h *BookHandler) GetBooks(c *gin.Context) {
	var filter models.BookFilter

	if err := c.ShouldBindQuery(&filter); err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid query parameters", "INVALID_QUERY", err.Error())
		return
	}

	// Set defaults
	if filter.Limit <= 0 {
		filter.Limit = 10
	}
	if filter.Limit > 100 {
		filter.Limit = 100
	}

	response, err := h.service.GetAllBooks(&filter)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to fetch books", "DATABASE_ERROR", err.Error())
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetBook retrieves a single book by ID
// @Summary      Get book by ID
// @Description  Get a single book by its ID
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Book ID"
// @Success      200  {object}  models.BookResponse
// @Failure      400  {object}  models.ErrorResponse
// @Failure      404  {object}  models.ErrorResponse
// @Failure      500  {object}  models.ErrorResponse
// @Router       /books/{id} [get]
func (h *BookHandler) GetBook(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid book ID", "INVALID_BOOK_ID", err.Error())
		return
	}

	book, err := h.service.GetBookByID(uint(id))
	if err != nil {
		if err.Error() == "book not found" {
			utils.SendError(c, http.StatusNotFound, "Book not found", "BOOK_NOT_FOUND")
			return
		}
		utils.SendError(c, http.StatusInternalServerError, "Failed to fetch book", "DATABASE_ERROR", err.Error())
		return
	}

	c.JSON(http.StatusOK, models.BookResponse{
		Success: true,
		Data:    book,
	})
}

// CreateBook creates a new book
// @Summary      Create a new book
// @Description  Create a new book with the provided information
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        book  body      models.CreateBookRequest  true  "Book information"
// @Success      201   {object}  models.BookResponse
// @Failure      400   {object}  models.ValidationErrorResponse
// @Failure      500   {object}  models.ErrorResponse
// @Router       /books [post]
func (h *BookHandler) CreateBook(c *gin.Context) {
	var req models.CreateBookRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid request format", "INVALID_REQUEST", err.Error())
		return
	}

	if err := h.validator.Struct(&req); err != nil {
		utils.SendValidationError(c, err)
		return
	}

	book, err := h.service.CreateBook(&req)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to create book", "DATABASE_ERROR", err.Error())
		return
	}

	c.JSON(http.StatusCreated, models.BookResponse{
		Success: true,
		Data:    book,
		Message: "Book created successfully",
	})
}

// UpdateBook updates an existing book
// @Summary      Update book
// @Description  Update an existing book by ID
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        id    path      int                       true  "Book ID"
// @Param        book  body      models.UpdateBookRequest  true  "Updated book information"
// @Success      200   {object}  models.BookResponse
// @Failure      400   {object}  models.ValidationErrorResponse
// @Failure      404   {object}  models.ErrorResponse
// @Failure      500   {object}  models.ErrorResponse
// @Router       /books/{id} [put]
func (h *BookHandler) UpdateBook(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid book ID", "INVALID_BOOK_ID", err.Error())
		return
	}

	var req models.UpdateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid request format", "INVALID_REQUEST", err.Error())
		return
	}

	if err := h.validator.Struct(&req); err != nil {
		utils.SendValidationError(c, err)
		return
	}

	book, err := h.service.UpdateBook(uint(id), &req)
	if err != nil {
		if err.Error() == "book not found" {
			utils.SendError(c, http.StatusNotFound, "Book not found", "BOOK_NOT_FOUND")
			return
		}
		utils.SendError(c, http.StatusInternalServerError, "Failed to update book", "DATABASE_ERROR", err.Error())
		return
	}

	c.JSON(http.StatusOK, models.BookResponse{
		Success: true,
		Data:    book,
		Message: "Book updated successfully",
	})
}

// DeleteBook deletes a book
// @Summary      Delete book
// @Description  Delete a book by ID
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        id  path      int  true  "Book ID"
// @Success      200 {object}  models.SuccessResponse
// @Failure      400 {object}  models.ErrorResponse
// @Failure      404 {object}  models.ErrorResponse
// @Failure      500 {object}  models.ErrorResponse
// @Router       /books/{id} [delete]
func (h *BookHandler) DeleteBook(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid book ID", "INVALID_BOOK_ID", err.Error())
		return
	}

	if err := h.service.DeleteBook(uint(id)); err != nil {
		if err.Error() == "book not found" {
			utils.SendError(c, http.StatusNotFound, "Book not found", "BOOK_NOT_FOUND")
			return
		}
		utils.SendError(c, http.StatusInternalServerError, "Failed to delete book", "DATABASE_ERROR", err.Error())
		return
	}

	utils.SendSuccess(c, http.StatusOK, "Book deleted successfully", nil)
}

// SearchBooks searches for books
// @Summary      Search books
// @Description  Search for books by title, author, or description
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        q  query     string  true  "Search query"
// @Success      200 {object}  models.BooksResponse
// @Failure      400 {object}  models.ErrorResponse
// @Failure      500 {object}  models.ErrorResponse
// @Router       /books/search [get]
func (h *BookHandler) SearchBooks(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		utils.SendError(c, http.StatusBadRequest, "Search query is required", "MISSING_QUERY")
		return
	}

	books, err := h.service.SearchBooks(query)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Search failed", "DATABASE_ERROR", err.Error())
		return
	}

	c.JSON(http.StatusOK, models.BooksResponse{
		Success: true,
		Data:    books,
		Total:   int64(len(books)),
	})
}
