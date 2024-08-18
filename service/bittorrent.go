package service

type BittorrentService interface{}

type bittorrentService struct{}

func NewBittorrentService() BittorrentService {
	return &bittorrentService{}
}
