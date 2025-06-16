package main

import (
	"library-backend/internal/api/handlers"
	"library-backend/internal/api/middleware"
	"library-backend/internal/config"
	"library-backend/internal/service"
	"library-backend/pkg/database"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	// Swagger imports
	_ "library-backend/docs" // This will be generated

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Library Management API
// @version         1.0
// @description     A library management system with URL processing capabilities
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.example.com/support
// @contact.email  support@example.com

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize logger
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	// Set Gin mode
	gin.SetMode(cfg.Server.Mode)

	log.Printf("🚀 Starting %s v%s", cfg.App.Name, cfg.App.Version)

	// Initialize database
	db, err := database.NewGormDB(&cfg.Database)
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}

	// Auto-migrate models
	if err := db.AutoMigrate(); err != nil {
		log.Fatalf("❌ Auto-migration failed: %v", err)
	}

	// Seed sample data
	if err := db.SeedData(); err != nil {
		logger.WithError(err).Warn("Failed to seed data")
	}

	// Initialize services
	bookService := service.NewBookService(db)
	urlService := service.NewURLService(db)

	// Initialize handlers
	bookHandler := handlers.NewBookHandler(bookService)
	urlHandler := handlers.NewURLHandler(urlService)

	// Setup router with middleware
	router := gin.New()

	// Global middleware
	router.Use(middleware.Logger(logger))
	router.Use(middleware.CORS())
	router.Use(middleware.ErrorHandler(logger))
	router.Use(gin.Recovery())

	// Swagger endpoint
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Health check
	router.GET("/health", func(c *gin.Context) {
		if err := db.Health(); err != nil {
			c.JSON(500, gin.H{
				"status":  "unhealthy",
				"error":   err.Error(),
				"app":     cfg.App.Name,
				"version": cfg.App.Version,
			})
			return
		}
		c.JSON(200, gin.H{
			"status":      "healthy",
			"app":         cfg.App.Name,
			"version":     cfg.App.Version,
			"environment": cfg.App.Environment,
			"swagger_url": "/swagger/index.html",
		})
	})

	// API routes
	api := router.Group("/api/v1")
	{
		// Books endpoints
		books := api.Group("/books")
		{
			books.GET("", bookHandler.GetBooks)
			books.POST("", bookHandler.CreateBook)
			books.GET("/search", bookHandler.SearchBooks)
			books.GET("/:id", bookHandler.GetBook)
			books.PUT("/:id", bookHandler.UpdateBook)
			books.DELETE("/:id", bookHandler.DeleteBook)
		}

		// URL processing endpoints
		api.POST("/process-url", urlHandler.ProcessURL)
		api.GET("/url-stats", urlHandler.GetStats)
	}

	log.Printf("🌟 Server running on port %s", cfg.Server.Port)
	log.Printf("📖 Swagger UI: http://localhost:%s/swagger/index.html", cfg.Server.Port)
	log.Printf("📊 Environment: %s", cfg.App.Environment)
	log.Printf("🗄️ Database: %s:%s/%s", cfg.Database.Host, cfg.Database.Port, cfg.Database.DBName)

	if err := router.Run(":" + cfg.Server.Port); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}
