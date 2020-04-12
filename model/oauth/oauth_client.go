package oauth

import "github.com/FlashFeiFei/my-gin/model"

type OauthClient struct {
	model.BaseModel
	ClientName   string `gorm:"cloumn:client_name"`
	ClientId     string `gorm:"cloumn:client_id"`
	ClientSecret string `gorm:"cloumn:client_secret"`
	RedirectUrl  string `gorm:"cloumn:redirect_url"`
}

func (OauthClient) TableName() string {
	return "oauth_client"
}
