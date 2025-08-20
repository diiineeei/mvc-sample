package handlers

import (
	"net/http"

	"github.com/diiineeei/mvc-sample/model"
	"github.com/diiineeei/mvc-sample/providers"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserHandler struct {
	provider *providers.UserProvider
}

func NewUserHandler(provider *providers.UserProvider) *UserHandler {
	return &UserHandler{provider: provider}
}
func (h *UserHandler) Webhook(c *gin.Context) {
	logger := c.MustGet("logger").(*zap.Logger)

	// Loga headers, URL e query parameters
	headers := c.Request.Header
	queryParams := c.Request.URL.Query()
	logger.Info("Webhook received",
		zap.String("method", c.Request.Method),
		zap.Any("headers", headers),
		zap.String("url", c.Request.URL.String()),
		zap.Any("query_params", queryParams),
	)

	// Tenta ler e logar o corpo JSON da requisição
	var payload interface{}
	if err := c.ShouldBindJSON(&payload); err != nil {
		// Se não houver JSON válido, loga que o corpo está vazio ou inválido
		logger.Warn("No valid JSON payload",
			zap.Error(err),
			zap.String("method", c.Request.Method),
		)
	} else {
		// Loga o payload JSON, se presente
		logger.Info("Webhook payload",
			zap.Any("payload", payload),
			zap.String("method", c.Request.Method),
		)
	}

	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

func (h *UserHandler) GetUser(c *gin.Context) {
	logger := c.MustGet("logger").(*zap.Logger)
	id := c.Param("id")
	user, err := h.provider.GetUser(id)
	if err != nil {
		logger.Error("Failed to get user", zap.String("id", id), zap.Error(err))
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	logger.Info("User retrieved", zap.String("id", id))
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	logger := c.MustGet("logger").(*zap.Logger)
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		logger.Error("Failed to bind JSON", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.provider.CreateUser(user); err != nil {
		logger.Error("Failed to create user", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	logger.Info("User created", zap.Any("user", user))
	c.JSON(http.StatusCreated, user)
}
