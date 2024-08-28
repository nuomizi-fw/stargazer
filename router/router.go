package router

import (
	"go.uber.org/fx"
)

var Module = fx.Module(
	"router",
	fx.Provide(
		NewStargazerRouter,
	),
	// Add new router below
	fx.Provide(
		NewAuthRouter,
		NewUserRouter,
		NewAria2Router,
		NewBittorrentRouter,
		NewMangaRouter,
		NewMusicRouter,
		NewNovelRouter,
		NewSearchRouter,
		NewVideoRouter,
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
	authRouter AuthRouter,
	userRouter UserRouter,
	aria2Router Aria2Router,
	bittorrentRouter BittorrentRouter,
	mangaRouter MangaRouter,
	musicRouter MusicRouter,
	novelRouter NovelRouter,
	searchRouter SearchRouter,
	videoRouter VideoRouter,
) StargazerRouters {
	return StargazerRouters{
		authRouter,
		userRouter,
		aria2Router,
		bittorrentRouter,
		mangaRouter,
		musicRouter,
		novelRouter,
		searchRouter,
		videoRouter,
	}
}
