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
			// User Services
			NewUserService,

			// Download Services
			NewAria2Service,
			NewBittorrentService,

			// Manga Services
			NewMangaService,

			// Music Services
			NewMusicService,

			// Novel Services

			// Search Services
			NewSearchService,

			// Video Services
			NewVideoService,
		),
	),
)

type StargazerService interface {
	User() UserService
	Aria2() Aria2Service
	Bittorrent() BittorrentService
	Manga() MangaService
	Music() MusicService
	Search() SearchService
	Video() VideoService
}

type stargazerService struct {
	user       UserService
	aria2      Aria2Service
	bittorrent BittorrentService
	manga      MangaService
	music      MusicService
	search     SearchService
	video      VideoService
}

func (ss *stargazerService) User() UserService {
	return ss.user
}

func (ss *stargazerService) Aria2() Aria2Service {
	return ss.aria2
}

func (ss *stargazerService) Bittorrent() BittorrentService {
	return ss.bittorrent
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
		user:       NewUserService(),
		aria2:      NewAria2Service(),
		bittorrent: NewBittorrentService(),
		manga:      NewMangaService(),
		music:      NewMusicService(),
		search:     NewSearchService(),
		video:      NewVideoService(),
	}
}
