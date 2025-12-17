package handlers

import (
	"net/http"
	"time"

	"task4/config"
	"task4/jwt"
	"task4/models"
	"task4/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct {
	*BaseHandler
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Password string `json:"password" binding:"required,min=6"`
	Email    string `json:"email" binding:"required,email"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type LoginResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{BaseHandler: NewBaseHandler(db)}
}

func (h *UserHandler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ErrorResponse(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	var existingUser models.User
	if err := h.DB.Where("username = ? OR email = ?", req.Username, req.Email).First(&existingUser).Error; err == nil {
		ErrorResponse(c, http.StatusConflict, "Username or email already exists")
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		ErrorResponse(c, http.StatusInternalServerError, "Failed to process password")
		return
	}

	user := models.User{
		Username: req.Username,
		Password: hashedPassword,
		Email:    req.Email,
	}

	if err := h.DB.Create(&user).Error; err != nil {
		ErrorResponse(c, http.StatusInternalServerError, "Failed to create user")
		return
	}

	userResponse := UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	SuccessResponse(c, userResponse)
}

func (h *UserHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ErrorResponse(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	var user models.User
	if err := h.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		ErrorResponse(c, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	if err := utils.CheckPassword(req.Password, user.Password); err != nil {
		ErrorResponse(c, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	token, err := jwt.GenerateToken(&user, config.AppConfig.JWT.SecretKey, config.AppConfig.JWT.ExpireHours)
	if err != nil {
		ErrorResponse(c, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	userResponse := UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	SuccessResponse(c, LoginResponse{
		Token: token,
		User:  userResponse,
	})
}

func (h *UserHandler) GetProfile(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		ErrorResponse(c, http.StatusUnauthorized, "User not authenticated")
		return
	}

	var user models.User
	if err := h.DB.First(&user, userID).Error; err != nil {
		ErrorResponse(c, http.StatusNotFound, "User not found")
		return
	}

	userResponse := UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	SuccessResponse(c, userResponse)
}
