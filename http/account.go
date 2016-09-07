package http

import (
	"standardgo/domain"

	"github.com/kataras/iris"
)

type AccountRestService struct {
	DomainService domain.AccountService
}

func NewAccountRestService(repo domain.AccountService) *AccountRestService {
	return &AccountRestService{repo}
}

func (a *AccountRestService) List(c *iris.Context) {
	res, err := a.DomainService.Find()
	if err != nil {
		c.JSON(iris.StatusOK, err)
	}
	c.JSON(iris.StatusOK, res)
}
