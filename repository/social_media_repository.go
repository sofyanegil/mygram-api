package repository

import (
	"mygram-api/models"

	"gorm.io/gorm"
)

type SocialMediaRepository interface {
	FindAll() (*[]models.SocialMedia, error)
	FindById(socialMediaId string) (*models.SocialMedia, error)
	Save(socialMedia models.SocialMedia) (*models.SocialMedia, error)
	Update(socialMedia models.SocialMedia) (*models.SocialMedia, error)
	Delete(socialMediaId string) error
}

type SocialMediaRepositoryImpl struct {
	DB *gorm.DB
}

func NewSocialMediaRepositoryImpl(DB *gorm.DB) SocialMediaRepository {
	return &SocialMediaRepositoryImpl{DB: DB}
}

func (sm *SocialMediaRepositoryImpl) Save(socialMedia models.SocialMedia) (*models.SocialMedia, error) {
	err := sm.DB.Create(&socialMedia).Error
	if err != nil {
		return nil, err
	}
	return &socialMedia, nil
}

func (sm *SocialMediaRepositoryImpl) FindAll() (*[]models.SocialMedia, error) {
	var socialMedias []models.SocialMedia
	err := sm.DB.Find(&socialMedias).Error
	if err != nil {
		return nil, err
	}
	return &socialMedias, nil
}

func (sm *SocialMediaRepositoryImpl) FindById(socialMediaId string) (*models.SocialMedia, error) {
	var socialMedia models.SocialMedia
	err := sm.DB.Where("id = ?", socialMediaId).First(&socialMedia).Error
	if err != nil {
		return nil, err
	}
	return &socialMedia, nil
}

func (sm *SocialMediaRepositoryImpl) Update(socialMedia models.SocialMedia) (*models.SocialMedia, error) {
	err := sm.DB.Save(socialMedia).Error
	if err != nil {
		return nil, err
	}
	return &socialMedia, nil
}

func (sm *SocialMediaRepositoryImpl) Delete(socialMediaId string) error {
	var socialMedia models.SocialMedia
	err := sm.DB.Where("id = ?", socialMediaId).Delete(&socialMedia).Error
	return err
}
