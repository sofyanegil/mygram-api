package repository

import (
	"errors"

	"mygram-api/helpers"
	"mygram-api/models"

	"gorm.io/gorm"
)

type PhotoRepository interface {
	FindAll() (*[]models.Photo, error)
	FindById(photoId string) (*models.Photo, error)
	Save(photo models.Photo) (*models.Photo, error)
	Update(photo models.Photo) (*models.Photo, error)
	Delete(photoId string) error
}

type PhotoRepositoryImpl struct {
	DB *gorm.DB
}

func NewPhotoRepositoryImpl(DB *gorm.DB) PhotoRepository {
	return &PhotoRepositoryImpl{DB: DB}
}

func (p *PhotoRepositoryImpl) FindAll() (*[]models.Photo, error) {
	var photos []models.Photo
	results := p.DB.Preload("Comments").Find(&photos)
	helpers.ErrorPanic(results.Error)
	return &photos, results.Error
}

func (p *PhotoRepositoryImpl) FindById(photoId string) (*models.Photo, error) {
	var photo models.Photo
	result := p.DB.Preload("Comments").Find(&photo, "ID = ?", photoId)
	if result.RowsAffected == 0 {
		return &models.Photo{}, errors.New("photo not found")
	}
	return &photo, nil
}

func (p *PhotoRepositoryImpl) Save(photo models.Photo) (*models.Photo, error) {
	result := p.DB.Create(&photo)
	return &photo, result.Error
}

func (p *PhotoRepositoryImpl) Update(photo models.Photo) (*models.Photo, error) {
	result := p.DB.Save(&photo)
	return &photo, result.Error
}

func (p *PhotoRepositoryImpl) Delete(photoId string) error {
	result := p.DB.Delete(&models.Photo{}, "ID = ?", photoId)
	return result.Error
}
