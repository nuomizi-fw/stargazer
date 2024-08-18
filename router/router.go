package router

import (
	"github.com/nuomizi-fw/stargazer/router/download"
	"github.com/nuomizi-fw/stargazer/router/manga"
	"github.com/nuomizi-fw/stargazer/router/music"
	"github.com/nuomizi-fw/stargazer/router/novel"
	"github.com/nuomizi-fw/stargazer/router/search"
	"github.com/nuomizi-fw/stargazer/router/video"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"router",
	fx.Provide(
		NewStargazerRouter,
	),
	// Add new router below
	fx.Provide(
		// User routers
		NewUserRouter,

		// Download routers
		download.NewAria2Router,
		download.NewBittorrentRouter,

		// Manga routers
		manga.NewMangaRouter,

		// Music routers
		music.NewMusicRouter,

		// Novel routers
		novel.NewNovelRouter,

		// Search routers
		search.NewSearchRouter,

		// Video routers
		video.NewVideoRouter,
	),
)

type StargazerRouter interface {
	InitRouter()
}

type StargazerRouters []StargazerRouter

func (sr StargazerRouters) InitRouter() {
	for _, router := range sr {
		router.InitRouter()
	}
}

func NewStargazerRouter(
	// User routers
	userRouter UserRouter,
	// Download routers
	aria2Router download.Aria2Router,
	bittorrentRouter download.BittorrentRouter,
	// Manga routers
	mangaRouter manga.MangaRouter,
	// Music routers
	musicRouter music.MusicRouter,
	// Novel routers
	novelRouter novel.NovelRouter,
	// Search routers
	searchRouter search.SearchRouter,
	// Video routers
	videoRouter video.VideoRouter,
) StargazerRouters {
	return StargazerRouters{
		// User routers
		userRouter,
		// Download routers
		aria2Router,
		bittorrentRouter,
		// Manga routers
		mangaRouter,
		// Music routers
		musicRouter,
		// Novel routers
		novelRouter,
		// Search routers
		searchRouter,
		// Video routers
		videoRouter,
	}
}
