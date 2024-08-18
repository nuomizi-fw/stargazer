package service

type SearchService interface{}

type searchService struct{}

func NewSearchService() SearchService {
	return &searchService{}
}
