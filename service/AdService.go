package service

import "github.com/ferminhg/golang-gin-poc/domain"

type AdService interface {
	FindAll() []domain.Ad
	Save(ad domain.Ad) domain.Ad
}

type adService struct {
	ads []domain.Ad
}

func New() AdService {
	return &adService{
		ads: []domain.Ad{},
	}
}

func (service *adService) FindAll() []domain.Ad {
	return service.ads
}

func (service *adService) Save(ad domain.Ad) domain.Ad {
	service.ads = append(service.ads, ad)
	return ad
}
