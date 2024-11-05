package service

type DownloaderService interface{}

type downloaderService struct{}

func NewDownloaderService() DownloaderService {
	return &downloaderService{}
}
