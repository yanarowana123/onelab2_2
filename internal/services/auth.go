package services

type IAuthService interface {
	IsGranted(key string) bool
}
type AuthService struct {
	key string
}

func NewAuthService(key string) *AuthService {
	return &AuthService{key: key}
}

func (s *AuthService) IsGranted(key string) bool {
	return s.key == key
}
