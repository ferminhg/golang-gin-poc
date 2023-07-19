package infrastructure

import (
	"github.com/ferminhg/golang-gin-poc/domain"
	"github.com/ferminhg/golang-gin-poc/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AdController interface {
	FindAll() []domain.Ad
	Save(ctx *gin.Context) error
}

type controller struct {
	service service.AdService
}

var validate *validator.Validate

func New(service service.AdService) AdController {
	validate = validator.New()
	validate.RegisterValidation("is-cool", ValidateCoolName)

	return &controller{service: service}
}

func (c *controller) FindAll() []domain.Ad {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) error {
	var ad domain.Ad
	err := ctx.ShouldBind(&ad)
	if err != nil {
		return err
	}

	err = validate.Struct(ad)
	if err != nil {
		return err
	}

	c.service.Save(ad)
	return nil
}
