package handlers

import (
	"library-backend/internal/models"
	"library-backend/internal/service"
	"library-backend/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type URLHandler struct {
	service   *service.URLService
	validator *validator.Validate
}

func NewURLHandler(service *service.URLService) *URLHandler {
	return &URLHandler{
		service:   service,
		validator: validator.New(),
	}
}

// ProcessURL processes a URL based on the specified operation
// @Summary      Process URL
// @Description  Process a URL based on the specified operation (canonical, redirection, or all)
// @Tags         url-processing
// @Accept       json
// @Produce      json
// @Param        request  body      models.URLRequest  true  "URL processing request"
// @Success      200      {object}  models.URLResponse
// @Failure      400      {object}  models.ValidationErrorResponse
// @Failure      500      {object}  models.ErrorResponse
// @Router       /process-url [post]
func (h *URLHandler) ProcessURL(c *gin.Context) {
	var req models.URLRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid request format", "INVALID_REQUEST", err.Error())
		return
	}

	if err := h.validator.Struct(&req); err != nil {
		utils.SendValidationError(c, err)
		return
	}

	// Get client info
	clientIP := c.ClientIP()
	userAgent := c.GetHeader("User-Agent")

	response, err := h.service.ProcessURL(&req, clientIP, userAgent)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to process URL", "PROCESSING_ERROR", err.Error())
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetStats retrieves URL processing statistics
// @Summary      Get URL processing statistics
// @Description  Get statistics about URL processing operations
// @Tags         url-processing
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.SuccessResponse
// @Failure      500  {object}  models.ErrorResponse
// @Router       /url-stats [get]
func (h *URLHandler) GetStats(c *gin.Context) {
	stats, err := h.service.GetProcessingStats()
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to get statistics", "STATS_ERROR", err.Error())
		return
	}

	utils.SendSuccess(c, http.StatusOK, "Statistics retrieved successfully", stats)
}
