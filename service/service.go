package service

import (
	"github.com/nuomizi-fw/stargazer/core"
	"github.com/nuomizi-fw/stargazer/pkg/keystore"
	"github.com/nuomizi-fw/stargazer/repository"
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
	logger core.StargazerLogger,
	repository repository.Repository,
	ks *keystore.KeyStore,
) StargazerService {
	return StargazerService{
		User:       NewUserService(logger, repository, ks),
		Rss:        NewRssService(),
		Downloader: NewDownloaderService(),
		Search:     NewSearchService(),
	}
}
