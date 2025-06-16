package service

import (
	"fmt"
	"library-backend/internal/models"
	"library-backend/pkg/database"
	"net/url"
	"strings"
)

type URLService struct {
	db *database.Database
}

func NewURLService(db *database.Database) *URLService {
	return &URLService{db: db}
}

func (s *URLService) ProcessURL(request *models.URLRequest, clientIP, userAgent string) (*models.URLResponse, error) {
	parsedURL, err := url.Parse(request.URL)
	if err != nil {
		return nil, fmt.Errorf("invalid URL format: %w", err)
	}

	var processedURL string

	switch request.Operation {
	case "canonical":
		processedURL = s.canonicalCleanup(parsedURL)
	case "redirection":
		processedURL = s.redirectionCleanup(parsedURL)
	case "all":
		canonicalURL, _ := url.Parse(s.canonicalCleanup(parsedURL))
		processedURL = s.redirectionCleanup(canonicalURL)
	default:
		return nil, fmt.Errorf("invalid operation: %s", request.Operation)
	}

	// Log the operation
	log := &models.URLProcessLog{
		OriginalURL:  request.URL,
		ProcessedURL: processedURL,
		Operation:    request.Operation,
		IPAddress:    clientIP,
		UserAgent:    userAgent,
	}

	if err := s.db.Create(log).Error; err != nil {
		// Log error but don't fail the request
		fmt.Printf("Failed to log URL processing: %v\n", err)
	}

	return &models.URLResponse{
		Success:      true,
		ProcessedURL: processedURL,
		Original:     request.URL,
		Operation:    request.Operation,
		LogID:        log.ID,
	}, nil
}

func (s *URLService) canonicalCleanup(parsedURL *url.URL) string {
	parsedURL.RawQuery = ""
	parsedURL.Path = strings.TrimSuffix(parsedURL.Path, "/")
	return parsedURL.String()
}

func (s *URLService) redirectionCleanup(parsedURL *url.URL) string {
	parsedURL.Host = "www.byfood.com"
	return strings.ToLower(parsedURL.String())
}

func (s *URLService) GetProcessingStats() (map[string]interface{}, error) {
	var stats struct {
		TotalRequests int64
		ByOperation   map[string]int64
	}

	// Total requests
	s.db.Model(&models.URLProcessLog{}).Count(&stats.TotalRequests)

	// By operation
	var operationStats []struct {
		Operation string
		Count     int64
	}

	s.db.Model(&models.URLProcessLog{}).
		Select("operation, count(*) as count").
		Group("operation").
		Scan(&operationStats)

	stats.ByOperation = make(map[string]int64)
	for _, stat := range operationStats {
		stats.ByOperation[stat.Operation] = stat.Count
	}

	return map[string]interface{}{
		"total_requests": stats.TotalRequests,
		"by_operation":   stats.ByOperation,
	}, nil
}
