package service

type BangumiService interface{}

type bangumiService struct{}

func NewBangumiService() BangumiService {
	return &bangumiService{}
}
