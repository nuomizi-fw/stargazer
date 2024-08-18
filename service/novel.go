package service

type NovelService interface{}

type novelService struct{}

func NewNovelService() NovelService {
	return novelService{}
}
