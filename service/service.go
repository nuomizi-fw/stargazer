package service

type StargazerService struct {
	User       UserService
	Rss        RssService
	Downloader DownloaderService
	Search     SearchService
}

func NewStargazerService() StargazerService {
	return StargazerService{
		User:       NewUserService(),
		Rss:        NewRssService(),
		Downloader: NewDownloaderService(),
		Search:     NewSearchService(),
	}
}
