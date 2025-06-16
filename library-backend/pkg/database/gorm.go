package database

import (
	"fmt"
	"library-backend/internal/config"
	"library-backend/internal/models"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	*gorm.DB
}

func NewGormDB(cfg *config.DatabaseConfig) (*Database, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port, cfg.SSLMode, cfg.TimeZone)

	// GORM configuration
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	}

	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Configure connection pool
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database instance: %w", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("✅ Database connected successfully")

	return &Database{db}, nil
}

// AutoMigrate creates/updates tables based on models
func (db *Database) AutoMigrate() error {
	log.Println("🔄 Running auto-migration...")

	err := db.DB.AutoMigrate(
		&models.Book{},
		&models.URLProcessLog{},
	)

	if err != nil {
		return fmt.Errorf("auto-migration failed: %w", err)
	}

	log.Println("✅ Auto-migration completed successfully")
	return nil
}

// SeedData inserts sample data
func (db *Database) SeedData() error {
	log.Println("🌱 Seeding sample data...")

	// Check if books already exist
	var count int64
	db.Model(&models.Book{}).Count(&count)
	if count > 0 {
		log.Println("📚 Sample data already exists, skipping seed")
		return nil
	}

	sampleBooks := []models.Book{
		{
			Title:       "The Go Programming Language",
			Author:      "Alan Donovan",
			Year:        2015,
			ISBN:        "9780134190440",
			Description: "A comprehensive guide to Go programming",
		},
		{
			Title:       "Clean Code",
			Author:      "Robert C. Martin",
			Year:        2008,
			ISBN:        "9780132350884",
			Description: "A handbook of agile software craftsmanship",
		},
		{
			Title:       "Design Patterns",
			Author:      "Gang of Four",
			Year:        1994,
			ISBN:        "9780201633610",
			Description: "Elements of Reusable Object-Oriented Software",
		},
		{
			Title:       "Microservices Patterns",
			Author:      "Chris Richardson",
			Year:        2018,
			Description: "With examples in Java",
		},
	}

	result := db.Create(&sampleBooks)
	if result.Error != nil {
		return fmt.Errorf("failed to seed data: %w", result.Error)
	}

	log.Printf("✅ Seeded %d sample books", len(sampleBooks))
	return nil
}

// Health check
func (db *Database) Health() error {
	sqlDB, err := db.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Ping()
}
