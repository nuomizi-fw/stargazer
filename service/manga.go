package service

type MangaService interface{}

type mangaService struct{}

func NewMangaService() MangaService {
	return mangaService{}
}
