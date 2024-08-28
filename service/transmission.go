package service

type TransmissionService interface{}

type transmissionService struct{}

func NewTransmissionService() TransmissionService {
	return &transmissionService{}
}
