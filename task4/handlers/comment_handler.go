package handlers

import (
	"net/http"
	"strconv"

	"task4/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CommentHandler struct {
	*BaseHandler
}

type CreateCommentRequest struct {
	Content string `json:"content" binding:"required,min=1"`
}

func NewCommentHandler(db *gorm.DB) *CommentHandler {
	return &CommentHandler{BaseHandler: NewBaseHandler(db)}
}

func (h *CommentHandler) CreateComment(c *gin.Context) {
	userID := c.GetUint("user_id")

	postID, err := strconv.ParseUint(c.Param("post_id"), 10, 32)
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, "Invalid post ID")
		return
	}

	var post models.Post
	if err := h.DB.First(&post, postID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ErrorResponse(c, http.StatusNotFound, "Post not found")
		} else {
			ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch post")
		}
		return
	}

	var req CreateCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ErrorResponse(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	comment := models.Comment{
		Content: req.Content,
		UserID:  userID,
		PostID:  uint(postID),
	}

	if err := h.DB.Create(&comment).Error; err != nil {
		ErrorResponse(c, http.StatusInternalServerError, "Failed to create comment")
		return
	}

	h.DB.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, username")
	}).First(&comment, comment.ID)

	SuccessResponse(c, comment)
}

func (h *CommentHandler) GetPostComments(c *gin.Context) {
	postID, err := strconv.ParseUint(c.Param("post_id"), 10, 32)
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, "Invalid post ID")
		return
	}

	var post models.Post
	if err := h.DB.First(&post, postID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ErrorResponse(c, http.StatusNotFound, "Post not found")
		} else {
			ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch post")
		}
		return
	}

	var comments []models.Comment
	if err := h.DB.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, username")
	}).Where("post_id = ?", postID).Order("created_at DESC").Find(&comments).Error; err != nil {
		ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch comments")
		return
	}

	SuccessResponse(c, comments)
}
