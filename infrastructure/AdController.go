package infrastructure

import (
	"github.com/ferminhg/golang-gin-poc/domain"
	"github.com/ferminhg/golang-gin-poc/service"
	"github.com/gin-gonic/gin"
)

type AdController interface {
	FindAll() []domain.Ad
	Save(ctx *gin.Context) domain.Ad
}

type controller struct {
	service service.AdService
}

func New(service service.AdService) AdController {
	return &controller{service: service}
}

func (c *controller) FindAll() []domain.Ad {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) domain.Ad {
	var ad domain.Ad
	ctx.BindJSON(&ad)
	c.service.Save(ad)
	return ad
}
