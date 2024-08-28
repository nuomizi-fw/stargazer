package service

type RssService interface{}

type rssService struct{}

func NewRssService() RssService {
	return &rssService{}
}
