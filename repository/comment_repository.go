package repository

import (
	"errors"

	"mygram-api/models"

	"gorm.io/gorm"
)

type CommentRepository interface {
	FindAll() (*[]models.Comment, error)
	FindById(id string) (*models.Comment, error)
	Save(comment models.Comment) (*models.Comment, error)
	Update(comment models.Comment) (*models.Comment, error)
	Delete(id string) error
}

type CommentRepositoryImpl struct {
	DB *gorm.DB
}

func NewCommentRepositoryImpl(DB *gorm.DB) CommentRepository {
	return &CommentRepositoryImpl{DB: DB}
}

func (c *CommentRepositoryImpl) FindAll() (*[]models.Comment, error) {
	var comments []models.Comment
	result := c.DB.Find(&comments)
	if result.Error != nil {
		return nil, result.Error
	}
	return &comments, nil
}

func (c *CommentRepositoryImpl) FindById(id string) (*models.Comment, error) {
	var comment models.Comment
	result := c.DB.Find(&comment, "ID = ?", id)
	if result.RowsAffected == 0 {
		return &models.Comment{}, errors.New("comment not found")
	}
	return &comment, nil
}

func (c *CommentRepositoryImpl) Save(comment models.Comment) (*models.Comment, error) {
	result := c.DB.Create(&comment)
	if result.Error != nil {
		return &models.Comment{}, result.Error
	}
	return &comment, nil
}

func (c *CommentRepositoryImpl) Update(comment models.Comment) (*models.Comment, error) {
	result := c.DB.Save(&comment)
	if result.Error != nil {
		return &models.Comment{}, result.Error
	}
	return &comment, nil
}

func (c *CommentRepositoryImpl) Delete(id string) error {
	var comment models.Comment
	result := c.DB.Delete(&comment, "ID = ?", id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
