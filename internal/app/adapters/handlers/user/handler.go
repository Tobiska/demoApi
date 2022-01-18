package user

import (
	"demoApi/internal/app/adapters/handlers"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) handlers.Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Register(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.SignUpUser())
		auth.POST("/sign-in", h.SignInUser())
	}

	api := router.Group("/api/users")
	{
		api.GET("/", h.GetAllUsers())
		api.GET("/:id", h.GetUserById())
		api.POST("/:id", h.CreateUser())
		api.DELETE("/:id", h.DeleteUser())
	}
}

func (h *Handler) GetAllUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		//todo
	}
}

func (h *Handler) GetUserById() gin.HandlerFunc {
	return func(c *gin.Context) {
		//todo
	}
}

func (h *Handler) CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		//todo
	}
}

func (h *Handler) DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		//todo
	}
}

func (h *Handler) SignInUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		//todo
	}
}

func (h *Handler) SignUpUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		//todo
	}
}
