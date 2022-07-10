package service

type OAuthService interface {
}

type oauthService struct {
}

func NewOAuthService() *oauthService {
	return &oauthService{}
}
