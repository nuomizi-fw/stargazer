//go:generate moq -rm -pkg mock -out mock/user_mock.go . UserService
package service

import (
	"github.com/nuomizi-fw/stargazer/core"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"service",
	fx.Options(
		fx.Provide(NewStargazerService),
		// Add new service below
		fx.Provide(
			NewUserService,
		),
	),
)

type StargazerService struct {
	User       UserService
	Rss        RssService
	Downloader DownloaderService
	Search     SearchService
}

func NewStargazerService(
	db core.StargazerDB,
) StargazerService {
	return StargazerService{
		User: NewUserService(db),
	}
}
