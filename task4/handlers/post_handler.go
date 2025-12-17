package handlers

import (
	"net/http"
	"strconv"

	"task4/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PostHandler struct {
	*BaseHandler
}

type CreatePostRequest struct {
	Title   string `json:"title" binding:"required,min=1,max=200"`
	Content string `json:"content" binding:"required,min=1"`
}

type UpdatePostRequest struct {
	Title   string `json:"title" binding:"omitempty,min=1,max=200"`
	Content string `json:"content" binding:"omitempty,min=1"`
}

func NewPostHandler(db *gorm.DB) *PostHandler {
	return &PostHandler{BaseHandler: NewBaseHandler(db)}
}

func (h *PostHandler) CreatePost(c *gin.Context) {
	userID := c.GetUint("user_id")

	var req CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ErrorResponse(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	post := models.Post{
		Title:   req.Title,
		Content: req.Content,
		UserID:  userID,
	}

	if err := h.DB.Create(&post).Error; err != nil {
		ErrorResponse(c, http.StatusInternalServerError, "Failed to create post")
		return
	}

	SuccessResponse(c, post)
}

func (h *PostHandler) GetAllPosts(c *gin.Context) {
	var posts []models.Post

	query := h.DB.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, username, email")
	}).Order("created_at DESC")

	if err := query.Find(&posts).Error; err != nil {
		ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch posts")
		return
	}

	SuccessResponse(c, posts)
}

func (h *PostHandler) GetPost(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, "Invalid post ID")
		return
	}

	var post models.Post
	if err := h.DB.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, username, email")
	}).Preload("Comments.User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, username")
	}).First(&post, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ErrorResponse(c, http.StatusNotFound, "Post not found")
		} else {
			ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch post")
		}
		return
	}

	SuccessResponse(c, post)
}

func (h *PostHandler) UpdatePost(c *gin.Context) {
	userID := c.GetUint("user_id")
	postID, err := strconv.ParseUint(c.Param("id"), 10, 32)
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

	if post.UserID != userID {
		ErrorResponse(c, http.StatusForbidden, "You are not authorized to update this post")
		return
	}

	var req UpdatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ErrorResponse(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	updates := make(map[string]interface{})
	if req.Title != "" {
		updates["title"] = req.Title
	}
	if req.Content != "" {
		updates["content"] = req.Content
	}

	if err := h.DB.Model(&post).Updates(updates).Error; err != nil {
		ErrorResponse(c, http.StatusInternalServerError, "Failed to update post")
		return
	}

	SuccessResponse(c, post)
}

func (h *PostHandler) DeletePost(c *gin.Context) {
	userID := c.GetUint("user_id")
	postID, err := strconv.ParseUint(c.Param("id"), 10, 32)
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

	if post.UserID != userID {
		ErrorResponse(c, http.StatusForbidden, "You are not authorized to delete this post")
		return
	}

	if err := h.DB.Delete(&post).Error; err != nil {
		ErrorResponse(c, http.StatusInternalServerError, "Failed to delete post")
		return
	}

	SuccessResponse(c, gin.H{"message": "Post deleted successfully"})
}
