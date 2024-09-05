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
			NewAuthService,
			NewUserService,
			NewAria2Service,
			NewBittorrentService,
			NewTransmissionService,
			NewMangaService,
			NewMusicService,
			NewNovelService,
			NewSearchService,
			NewBangumiService,
		),
	),
)

type StargazerService struct {
	auth         AuthService
	user         UserService
	rss          RssService
	aria2        Aria2Service
	bittorrent   BittorrentService
	transmission TransmissionService
	manga        MangaService
	music        MusicService
	search       SearchService
	bangumi      BangumiService
}

func NewStargazerService(
	db core.StargazerDB,
	logger core.StargazerLogger,
) StargazerService {
	return StargazerService{
		auth:       NewAuthService(db, logger),
		user:       NewUserService(db, logger),
		aria2:      NewAria2Service(),
		bittorrent: NewBittorrentService(),
		manga:      NewMangaService(),
		music:      NewMusicService(),
		search:     NewSearchService(),
		bangumi:    NewBangumiService(),
	}
}
