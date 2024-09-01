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

type StargazerService interface {
	Auth() AuthService
	User() UserService
	Rss() RssService
	Aria2() Aria2Service
	Bittorrent() BittorrentService
	Transmission() TransmissionService
	Manga() MangaService
	Music() MusicService
	Search() SearchService
	Bangumi() BangumiService
}

type stargazerService struct {
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

func (ss *stargazerService) Auth() AuthService {
	return ss.auth
}

func (ss *stargazerService) User() UserService {
	return ss.user
}

func (ss *stargazerService) Rss() RssService {
	return ss.rss
}

func (ss *stargazerService) Aria2() Aria2Service {
	return ss.aria2
}

func (ss *stargazerService) Bittorrent() BittorrentService {
	return ss.bittorrent
}

func (ss *stargazerService) Transmission() TransmissionService {
	return ss.transmission
}

func (ss *stargazerService) Manga() MangaService {
	return ss.manga
}

func (ss *stargazerService) Music() MusicService {
	return ss.music
}

func (ss *stargazerService) Search() SearchService {
	return ss.search
}

func (ss *stargazerService) Bangumi() BangumiService {
	return ss.bangumi
}

func NewStargazerService(
	db core.StargazerDB,
	logger core.StargazerLogger,
) StargazerService {
	return &stargazerService{
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
