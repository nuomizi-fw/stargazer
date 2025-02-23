package service

import (
	"github.com/nuomizi-fw/stargazer/core"
	"github.com/nuomizi-fw/stargazer/pkg/jwt"
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
	logger core.StargazerLogger,
	repository repository.Repository,
) StargazerService {
	// Generate key pair for JWT signing
	privateKey, publicKey, err := jwt.GenerateKeyPair()
	if err != nil {
		logger.Fatalf("Failed to generate key pair: %s", err)
	}

	return StargazerService{
		Auth:       NewAuthService(repository, privateKey, publicKey),
		User:       NewUserService(logger, repository),
		Rss:        NewRssService(),
		Downloader: NewDownloaderService(),
		Search:     NewSearchService(),
	}
}
