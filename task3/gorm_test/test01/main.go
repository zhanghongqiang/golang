package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type User struct {
	gorm.Model
	ID        uint
	Usename   string
	Email     string
	Password  string
	PostCount int `gorm:"default:0"`
	Posts     []Post
}

type Post struct {
	gorm.Model
	ID            uint
	Title         string
	UserID        uint
	CommentCount  int
	CommentStatus string
	User          User
	Comments      []Comment
}

type Comment struct {
	gorm.Model
	ID      uint
	Content string
	UserID  uint
	PostID  uint
	User    User
	Post    Post
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&User{}, &Post{}, &Comment{})
	if err != nil {
		fmt.Println("创建表失败", err)
	}
}

// 询某个用户发布的所有文章及其对应的评论信息
func GetUserPostsWithComments(db *gorm.DB, userID uint) (posts []Post, err error) {
	var user User
	err = db.Preload("Posts.Comments.User").Preload("Posts.User").Preload(clause.Associations).First(&user, userID).Error
	return user.Posts, err
}

// 查询评论数量最多的文章信息
func GetMostCommentedPost(db *gorm.DB) (Post, error) {
	var post Post
	result := db.Preload("User").
		Preload("Comments").
		Order(clause.OrderByColumn{Column: clause.Column{Name: "comment_count"}, Desc: true}).
		Limit(1).
		First(&post)

	if result.Error != nil {
		return post, result.Error
	}
	return post, nil
}

// 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段
func (p *Post) AfterCreate(db *gorm.DB) error {
	db.Model(&User{}).Where("id = ?", p.UserID).UpdateColumn("post_count",
		db.Select("count(*)").Model(&Post{}).Where("userId = ?", p.UserID))
	return nil
}

// 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"
func (c *Comment) AfterDelete(db *gorm.DB) error {

	var commentCount int64
	db.Model(&Comment{}).Where("postId = ?", c.PostID).Count(&commentCount)
	fmt.Printf("删除的评论数文章:%d的评论数量: %d\n", c.PostID, commentCount)
	if commentCount == 0 {
		db.Model(&Post{}).Where("id = ?", c.PostID).Update("comment_status", "有评论")
		fmt.Printf("文章:%d更新为无评论", c.PostID)
	}

	db.Model(&Post{}).Where("id = ?", c.PostID).Update("comment_count", commentCount)

	return nil
}

func (c *Comment) AfterCreate(db *gorm.DB) error {
	var commentCount int64
	db.Model(&Comment{}).Where("postId = ?", c.PostID).Count(&commentCount)
	fmt.Printf("文章:%d的评论数量: %d\n", c.PostID, commentCount)

	if commentCount > 0 {
		db.Model(&Post{}).Where("id = ?", c.PostID).Update("comment_status", "有评论")
		fmt.Printf("文章:%d更新为有评论", c.PostID)
	}

	db.Model(&Post{}).Where("id = ?", c.PostID).Update("comment_count", commentCount)

	return nil
}
