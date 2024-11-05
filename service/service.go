//go:generate moq -rm -pkg mock -out mock/user_mock.go . UserService
//go:generate moq -rm -pkg mock -out mock/downloader_mock.go . DownloaderService
//go:generate moq -rm -pkg mock -out mock/rss_mock.go . RssService
//go:generate moq -rm -pkg mock -out mock/manga_mock.go . MangaService
//go:generate moq -rm -pkg mock -out mock/music_mock.go . MusicService
//go:generate moq -rm -pkg mock -out mock/novel_mock.go . NovelService
//go:generate moq -rm -pkg mock -out mock/bangumi_mock.go . BangumiService
//go:generate moq -rm -pkg mock -out mock/search_mock.go . SearchService
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
			NewDownloaderService,
			NewMangaService,
			NewMusicService,
			NewNovelService,
			NewBangumiService,
			NewSearchService,
		),
	),
)

type StargazerService struct {
	User       UserService
	Rss        RssService
	Downloader DownloaderService
	Manga      MangaService
	Music      MusicService
	Novel      NovelService
	Bangumi    BangumiService
	Search     SearchService
}

func NewStargazerService(
	db core.StargazerDB,
	logger core.StargazerLogger,
) StargazerService {
	return StargazerService{
		User:       NewUserService(db, logger),
		Downloader: NewDownloaderService(),
		Manga:      NewMangaService(),
		Music:      NewMusicService(),
		Novel:      NewNovelService(),
		Bangumi:    NewBangumiService(),
		Search:     NewSearchService(),
	}
}
