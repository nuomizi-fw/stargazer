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

// @Summary 扫描本地动漫并自动刮削
// @Description 扫描指定目录中的所有动漫文件，并自动刮削相关信息
// @Tags Bangumi
// @Accept json
// @Produce json
// @Param directory body string true "Directory Path"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/bangumi/scan-local [post]
func (br BangumiRouter) ScanBangumi(c *fiber.Ctx) error {
	return nil
}

// @Summary 获取观看历史
// @Description 获取用户的观看历史
// @Tags WatchHistory
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/bangumi/watch-history [get]
func (br BangumiRouter) GetWatchHistory(c *fiber.Ctx) error {
	return nil
}

// @Summary 添加观看历史
// @Description 添加新的观看历史
// @Tags WatchHistory
// @Accept json
// @Produce json
// @Param history body string true "Watch History"
// @Success 200 {object} map[string]interface{}
// @Router /api/bangumi/watch-history [post]
func (br BangumiRouter) AddWatchHistory(c *fiber.Ctx) error {
	return nil
}

// @Summary 更新观看历史
// @Description 更新现有的观看历史
// @Tags WatchHistory
// @Accept json
// @Produce json
// @Param history body string true "Watch History"
// @Success 200 {object} map[string]interface{}
// @Router /api/bangumi/watch-history [put]
func (br BangumiRouter) UpdateWatchHistory(c *fiber.Ctx) error {
	return nil
}

// @Summary 删除观看历史
// @Description 删除指定的观看历史
// @Tags WatchHistory
// @Produce json
// @Param id path string true "Watch History ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/bangumi/watch-history/{id} [delete]
func (br BangumiRouter) DeleteWatchHistory(c *fiber.Ctx) error {
	return nil
}

// @Summary 获取收藏夹
// @Description 获取用户的收藏夹
// @Tags Favorites
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/bangumi/favorites [get]
func (br BangumiRouter) GetFavorites(c *fiber.Ctx) error {
	return nil
}

// @Summary 添加收藏
// @Description 添加新的收藏
// @Tags Favorites
// @Accept json
// @Produce json
// @Param favorite body string true "Favorite"
// @Success 200 {object} map[string]interface{}
// @Router /api/bangumi/favorites [post]
func (br BangumiRouter) AddFavorite(c *fiber.Ctx) error {
	return nil
}

// @Summary 删除收藏
// @Description 删除指定的收藏
// @Tags Favorites
// @Produce json
// @Param id path string true "Favorite ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/bangumi/favorites/{id} [delete]
func (br BangumiRouter) DeleteFavorite(c *fiber.Ctx) error {
	return nil
}

// @Summary 获取标签
// @Description 获取所有标签
// @Tags Tags
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/bangumi/tags [get]
func (br BangumiRouter) GetTags(c *fiber.Ctx) error {
	return nil
}

// @Summary 添加标签
// @Description 添加新的标签
// @Tags Tags
// @Accept json
// @Produce json
// @Param tag body string true "Tag"
// @Success 200 {object} map[string]interface{}
// @Router /api/bangumi/tags [post]
func (br BangumiRouter) AddTag(c *fiber.Ctx) error {
	return nil
}

// @Summary 删除标签
// @Description 删除指定的标签
// @Tags Tags
// @Produce json
// @Param id path string true "Tag ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/bangumi/tags/{id} [delete]
func (br BangumiRouter) DeleteTag(c *fiber.Ctx) error {
	return nil
}

// @Summary 获取集合
// @Description 获取所有集合
// @Tags Collections
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/bangumi/collections [get]
func (br BangumiRouter) GetCollections(c *fiber.Ctx) error {
	return nil
}

// @Summary 添加集合
// @Description 添加新的集合
// @Tags Collections
// @Accept json
// @Produce json
// @Param collection body string true "Collection"
// @Success 200 {object} map[string]interface{}
// @Router /api/bangumi/collections [post]
func (br BangumiRouter) AddCollection(c *fiber.Ctx) error {
	return nil
}

// @Summary 删除集合
// @Description 删除指定的集合
// @Tags Collections
// @Produce json
// @Param id path string true "Collection ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/bangumi/collections/{id} [delete]
func (br BangumiRouter) DeleteCollection(c *fiber.Ctx) error {
	return nil
}

// @Summary 获取番剧流媒体
// @Description 获取指定番剧的流媒体
// @Tags Feeds
// @Produce application/octet-stream
// @Param id path string true "Bangumi ID"
// @Success 200 {string} string "binary data"
// @Failure 404 {object} map[string]interface{}
// @Router /api/bangumi/feeds/{id} [get]
func (br BangumiRouter) BangumiFeeds(c *fiber.Ctx) error {
	return nil
}
