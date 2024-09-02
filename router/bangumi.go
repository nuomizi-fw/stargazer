package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nuomizi-fw/stargazer/core"
)

type BangumiRouter struct {
	stargazer core.StargazerServer
	logger    core.StargazerLogger
}

func NewBangumiRouter(stargazer core.StargazerServer, logger core.StargazerLogger) BangumiRouter {
	return BangumiRouter{
		stargazer: stargazer,
		logger:    logger,
	}
}

func (br BangumiRouter) InitRouter() {
	br.logger.Info("Initializing bangumi router")

	bangumi := br.stargazer.Api.Group("/bangumi")
	{
		bangumi.Post("/scan", br.ScanBangumi)

		// History
		bangumi.Get("/history", br.GetWatchHistory)
		bangumi.Post("/history", br.AddWatchHistory)
		bangumi.Put("/history/:id", br.UpdateWatchHistory)
		bangumi.Delete("/history/:id", br.DeleteWatchHistory)

		// Favorites
		bangumi.Get("/favorites", br.GetFavorites)
		bangumi.Post("/favorites", br.AddFavorite)
		bangumi.Delete("/favorites/:id", br.DeleteFavorite)

		// Tags
		bangumi.Get("/tags", br.GetTags)
		bangumi.Post("/tags", br.AddTag)
		bangumi.Delete("/tags/:id", br.DeleteTag)

		// Collections
		bangumi.Get("/collections", br.GetCollections)
		bangumi.Post("/collections", br.AddCollection)
		bangumi.Delete("/collections/:id", br.DeleteCollection)

		// Video Player
		bangumi.Get("/bangumi-feeds/:id", br.BangumiFeeds)
	}
}

// 扫描指定目录中的所有动漫文件，并自动刮削相关信息
func (br BangumiRouter) ScanBangumi(c *fiber.Ctx) error {
	return nil
}

// 获取用户的观看历史
func (br BangumiRouter) GetWatchHistory(c *fiber.Ctx) error {
	return nil
}

// 添加新的观看历史
func (br BangumiRouter) AddWatchHistory(c *fiber.Ctx) error {
	return nil
}

// 更新现有的观看历史
func (br BangumiRouter) UpdateWatchHistory(c *fiber.Ctx) error {
	return nil
}

// 删除指定的观看历史
func (br BangumiRouter) DeleteWatchHistory(c *fiber.Ctx) error {
	return nil
}

// 获取用户的收藏夹
func (br BangumiRouter) GetFavorites(c *fiber.Ctx) error {
	return nil
}

// 添加新的收藏
func (br BangumiRouter) AddFavorite(c *fiber.Ctx) error {
	return nil
}

// 删除指定的收藏
func (br BangumiRouter) DeleteFavorite(c *fiber.Ctx) error {
	return nil
}

// 获取所有标签
func (br BangumiRouter) GetTags(c *fiber.Ctx) error {
	return nil
}

// 添加新的标签
func (br BangumiRouter) AddTag(c *fiber.Ctx) error {
	return nil
}

// 删除指定的标签
func (br BangumiRouter) DeleteTag(c *fiber.Ctx) error {
	return nil
}

// 获取所有集合
func (br BangumiRouter) GetCollections(c *fiber.Ctx) error {
	return nil
}

// 添加新的集合
func (br BangumiRouter) AddCollection(c *fiber.Ctx) error {
	return nil
}

// 删除指定的集合
func (br BangumiRouter) DeleteCollection(c *fiber.Ctx) error {
	return nil
}

// 获取指定番剧的流媒体
func (br BangumiRouter) BangumiFeeds(c *fiber.Ctx) error {
	return nil
}
