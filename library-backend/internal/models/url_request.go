package models

import (
	"time"

	"gorm.io/gorm"
)

// URLProcessLog - Optional: Log URL processing for analytics
type URLProcessLog struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	OriginalURL  string         `json:"original_url" gorm:"type:text;not null"`
	ProcessedURL string         `json:"processed_url" gorm:"type:text;not null"`
	Operation    string         `json:"operation" gorm:"type:varchar(20);not null"`
	IPAddress    string         `json:"ip_address,omitempty" gorm:"type:varchar(45)"`
	UserAgent    string         `json:"user_agent,omitempty" gorm:"type:text"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`
}

func (URLProcessLog) TableName() string {
	return "url_process_logs"
}

// Request/Response DTOs
type URLRequest struct {
	URL       string `json:"url" validate:"required,url"`
	Operation string `json:"operation" validate:"required,oneof=redirection canonical all"`
}

type URLResponse struct {
	Success      bool   `json:"success"`
	ProcessedURL string `json:"processed_url"`
	Original     string `json:"original_url"`
	Operation    string `json:"operation"`
	LogID        uint   `json:"log_id,omitempty"`
}
