package models

import (
	"sv_be_posts/config"
	"time"
)

type Post struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	Category    string    `json:"category"`
	Status      string    `json:"status"`
	CreatedDate time.Time `json:"created_date" gorm:"autoCreateTime"`
	UpdatedDate time.Time `json:"updated_date" gorm:"autoUpdateTime"`
}

func Create(post *Post) error {
	return config.DB.Create(post).Error
}

func GetByID(id int) (*Post, error) {
	var post Post
	err := config.DB.First(&post, id).Error
	return &post, err
}

func GetPaginated(limit, offset int,status string) ([]Post, error) {
	var posts []Post
	query := config.DB
	if(status!=""){
		query = query.Where("status",status)	
	}
	err := query.Limit(limit).Offset(offset).Find(&posts).Error
	
	return posts, err
}

func Update(post *Post) error {
	return config.DB.Model(&Post{}).Where("id = ?", post.ID).Updates(post).Error
}

func Delete(id int) error {
	return config.DB.Delete(&Post{}, id).Error
}
