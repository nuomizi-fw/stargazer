package model

import "time"

// Bangumi 动画
type Bangumi struct {
	ID            uint     `json:"id" gorm:"primaryKey;autoIncrement"`  // ID
	Title         string   `json:"title"`                               // 标题
	EnglishTitle  string   `json:"english_title"`                       // 英文标题
	JapaneseTitle string   `json:"japanese_title"`                      // 日文标题
	Type          string   `json:"type"`                                // 类型(如:TV, OVA, Movie 等)
	Status        string   `json:"status"`                              // 状态（如：正在播放, 完结等）
	Score         float64  `json:"score"`                               // 评分
	Tags          string   `json:"tags"`                                // 标签
	Synopsis      string   `json:"synopsis"`                            // 简介
	CoverImage    string   `json:"cover_image"`                         // 封面图 URL
	TrailerURL    string   `json:"trailer_url"`                         // 预告片 URL
	Seasons       []Season `json:"seasons" gorm:"foreignKey:BangumiID"` // 季度
}

// Season 季度
type Season struct {
	ID           uint         `json:"id" gorm:"primaryKey;autoIncrement"`
	BangumiID    uint         `json:"bangumi_id"`
	SeasonNumber int          `json:"season_number"`                   // 季数
	StartDate    time.Time    `json:"start_date"`                      // 开始日期
	EndDate      time.Time    `json:"end_date"`                        // 结束日期
	Episodes     int          `json:"episodes"`                        // 集数
	Cast         []CastMember `json:"cast" gorm:"foreignKey:SeasonID"` // 演员阵容
}

// CastMember 演员
type CastMember struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	SeasonID uint   `json:"season_id"`
	Name     string `json:"name"`      // 演员姓名
	Role     string `json:"role"`      // 角色名
	ImageURL string `json:"image_url"` // 演员图像 URL
}
