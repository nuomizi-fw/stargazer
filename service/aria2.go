package service

type Aria2Service interface{}

type aria2Service struct{}

func NewAria2Service() Aria2Service {
	return &aria2Service{}
}
