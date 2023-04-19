package request

type CreateSocialMediaRequest struct {
	Name           string `json:"name" form:"name" validate:"required"`
	SocialMediaUrl string `json:"social_media_url" form:"social_media_url" validate:"required"`
}

type UpdateSocialMediaRequest struct {
	Name           *string `json:"name,omitempty" form:"name,omitempty"`
	SocialMediaUrl *string `json:"social_media_url,omitempty" form:"social_media_url,omitempty" validate:"omitempty"`
}
