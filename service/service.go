//go:generate moq -rm -pkg mock -out mock/user_mock.go . UserService
package service

import (
	"github.com/nuomizi-fw/stargazer/repository"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"service",
	fx.Options(
		fx.Provide(NewStargazerService),
		// Add new service below
		fx.Provide(
			NewAuthService,
			NewUserService,
		),
	),
)

type StargazerService struct {
	Auth       AuthService
	User       UserService
	Rss        RssService
	Downloader DownloaderService
	Search     SearchService
}

func NewStargazerService(
	repository repository.Repository,
) StargazerService {
	return StargazerService{
		Auth: NewAuthService(repository),
		User: NewUserService(repository),
	}
}
