package repositories

import (
	"sv_be_posts/models"
)

func CreatePost(post *models.Post) error {
	return models.Create(post)
}

func GetPosts() ([]models.Post, error) {
	return models.GetAll()
}

func GetPost(id int) (*models.Post, error) {
	return models.GetByID(id)
}

func UpdatePost(post *models.Post) error {
	return models.Update(post)
}

func DeletePost(id int) error {
	return models.Delete(id)
}
