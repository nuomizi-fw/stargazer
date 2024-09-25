//go:generate moq -rm -pkg mock -out mock/aria2_mock.go . Aria2Service
//go:generate moq -rm -pkg mock -out mock/auth_mock.go . AuthService
//go:generate moq -rm -pkg mock -out mock/bangumi_mock.go . BangumiService
//go:generate moq -rm -pkg mock -out mock/bittorrent_mock.go . BittorrentService
//go:generate moq -rm -pkg mock -out mock/manga_mock.go . MangaService
//go:generate moq -rm -pkg mock -out mock/music_mock.go . MusicService
//go:generate moq -rm -pkg mock -out mock/novel_mock.go . NovelService
//go:generate moq -rm -pkg mock -out mock/rss_mock.go . RssService
//go:generate moq -rm -pkg mock -out mock/search_mock.go . SearchService
//go:generate moq -rm -pkg mock -out mock/transmission_mock.go . TransmissionService
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
