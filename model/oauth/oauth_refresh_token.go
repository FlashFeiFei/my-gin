package oauth

import "github.com/FlashFeiFei/my-gin/model"

type OauthRefreshToken struct {
	model.BaseModel
	OauthClient_id int    `gorm:"cloumn:oauth_client_id"`
	RefreshToken   string `gorm:"cloumn:refresh_token"`
}

func (OauthRefreshToken) TableName() string {
	return "oauth_refresh_token"
}
