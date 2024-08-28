package service

import (
	"go.uber.org/fx"
)

var Module = fx.Module(
	"service",
	fx.Options(
		fx.Provide(NewService),
		// Add new service below
		fx.Provide(
			NewAuthService,
			NewUserService,
			NewAria2Service,
			NewBittorrentService,
			NewMangaService,
			NewMusicService,
			NewNovelService,
			NewSearchService,
			NewVideoService,
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
	Video() VideoService
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
	video        VideoService
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

func (ss *stargazerService) Video() VideoService {
	return ss.video
}

func NewService() StargazerService {
	return &stargazerService{
		auth:       NewAuthService(),
		user:       NewUserService(),
		aria2:      NewAria2Service(),
		bittorrent: NewBittorrentService(),
		manga:      NewMangaService(),
		music:      NewMusicService(),
		search:     NewSearchService(),
		video:      NewVideoService(),
	}
}
