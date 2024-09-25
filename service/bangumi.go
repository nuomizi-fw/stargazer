package service

import entModel "github.com/nuomizi-fw/stargazer/ent"

//go:generate moq -pkg service_test -out bangumi_test.go . BangumiService
type BangumiService interface {
	// Bangumi operations
	GetBangumi(id int) (*entModel.Bangumi, error)
	ListBangumis(page, pageSize int) ([]*entModel.Bangumi, error)
	CreateBangumi(bangumi *entModel.Bangumi) error
	UpdateBangumi(id int, bangumi *entModel.Bangumi) error
	DeleteBangumi(id int) error
	SearchBangumis(query string) ([]*entModel.Bangumi, error)

	// Season operations
	GetSeason(id int) (*entModel.Season, error)
	ListSeasons(bangumiID int) ([]*entModel.Season, error)
	CreateSeason(season *entModel.Season) error
	UpdateSeason(id int, season *entModel.Season) error
	DeleteSeason(id int) error

	// Episode operations
	GetEpisode(id int) (*entModel.Episode, error)
	ListEpisodes(seasonID int) ([]*entModel.Episode, error)
	CreateEpisode(episode *entModel.Episode) error
	UpdateEpisode(id int, episode *entModel.Episode) error
	DeleteEpisode(id int) error

	// CastMember operations
	GetCastMember(id int) (*entModel.CastMember, error)
	ListCastMembers(seasonID int) ([]*entModel.CastMember, error)
	CreateCastMember(castMember *entModel.CastMember) error
	UpdateCastMember(id int, castMember *entModel.CastMember) error
	DeleteCastMember(id int) error

	// Scraper operations
	ScrapeAndUpdateBangumi(id int) error
	ScrapeAndUpdateSeason(bangumiID, seasonID int) error
	ScrapeAndUpdateEpisode(seasonID, episodeID int) error
}

type bangumiService struct{}

func NewBangumiService() BangumiService {
	return &bangumiService{}
}

// Bangumi operations
func (bs *bangumiService) GetBangumi(id int) (*entModel.Bangumi, error) {
	return nil, nil
}

func (bs *bangumiService) ListBangumis(page, pageSize int) ([]*entModel.Bangumi, error) {
	return nil, nil
}

func (bs *bangumiService) CreateBangumi(bangumi *entModel.Bangumi) error {
	return nil
}

func (bs *bangumiService) UpdateBangumi(id int, bangumi *entModel.Bangumi) error {
	return nil
}

func (bs *bangumiService) DeleteBangumi(id int) error {
	return nil
}

func (bs *bangumiService) SearchBangumis(query string) ([]*entModel.Bangumi, error) {
	return nil, nil
}

// Season operations
func (bs *bangumiService) GetSeason(id int) (*entModel.Season, error) {
	return nil, nil
}

func (bs *bangumiService) ListSeasons(bangumiID int) ([]*entModel.Season, error) {
	return nil, nil
}

func (bs *bangumiService) CreateSeason(season *entModel.Season) error {
	return nil
}

func (bs *bangumiService) UpdateSeason(id int, season *entModel.Season) error {
	return nil
}

func (bs *bangumiService) DeleteSeason(id int) error {
	return nil
}

// Episode operations
func (bs *bangumiService) GetEpisode(id int) (*entModel.Episode, error) {
	return nil, nil
}

func (bs *bangumiService) ListEpisodes(seasonID int) ([]*entModel.Episode, error) {
	return nil, nil
}

func (bs *bangumiService) CreateEpisode(episode *entModel.Episode) error {
	return nil
}

func (bs *bangumiService) UpdateEpisode(id int, episode *entModel.Episode) error {
	return nil
}

func (bs *bangumiService) DeleteEpisode(id int) error {
	return nil
}

// CastMember operations
func (bs *bangumiService) GetCastMember(id int) (*entModel.CastMember, error) {
	return nil, nil
}

func (bs *bangumiService) ListCastMembers(seasonID int) ([]*entModel.CastMember, error) {
	return nil, nil
}

func (bs *bangumiService) CreateCastMember(castMember *entModel.CastMember) error {
	return nil
}

func (bs *bangumiService) UpdateCastMember(id int, castMember *entModel.CastMember) error {
	return nil
}

func (bs *bangumiService) DeleteCastMember(id int) error {
	return nil
}

// Scraper operations
func (bs *bangumiService) ScrapeAndUpdateBangumi(id int) error {
	return nil
}

func (bs *bangumiService) ScrapeAndUpdateSeason(bangumiID, seasonID int) error {
	return nil
}

func (bs *bangumiService) ScrapeAndUpdateEpisode(seasonID, episodeID int) error {
	return nil
}
