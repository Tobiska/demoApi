package user

import (
	"demoApi/internal/app/adapters/handlers"
	"demoApi/internal/app/domain/user"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
		api.POST("/", h.CreateUser())
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
		id, err := strconv.Atoi(c.Params.ByName("id"))
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
		}

		u, err := h.service.GetByID(c.Request.Context(), id)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
		}

		c.JSON(http.StatusOK, u)
	}
}

func (h *Handler) CreateUser() gin.HandlerFunc {
	type request struct {
		Name     string `json:"name" binding:"required,min=2,max=10"`
		Email    string `json:"email" validate:"required, email"`
		Password string `json:"password" binding:"required,min=6,max=12"`
	}
	return func(c *gin.Context) {
		var rt request
		if err := c.ShouldBindJSON(&rt); err != nil {
			c.String(http.StatusUnprocessableEntity, err.Error())
		}

		dto := &user.CreateUserDto{
			Name:     rt.Name,
			Email:    rt.Email,
			Password: rt.Password,
		}

		if err := h.service.Create(c.Request.Context(), dto); err != nil {
			c.String(http.StatusNotFound, err.Error())
		}

		c.String(http.StatusCreated, "User was created successfully!")
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
