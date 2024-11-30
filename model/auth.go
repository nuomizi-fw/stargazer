package model

type RefreshToken struct {
	AccessToken  string
	RefreshToken string
	ExpiresIn    int
}
