package oauth

type Client struct {
	RedirectUrl string `form:"redirect_url" binding:"required"`
	ClientName  string `form:"client_name" binding:"required"`
}
