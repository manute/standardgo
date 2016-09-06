package http

import (
	"standardgo/domain"

	"github.com/kataras/iris"
)

type AccountRestService interface {
	List(c *iris.Context)
}

type accountService struct {
	DomainService domain.AccountService
}

func NewAccountRestService(repo domain.AccountService) AccountRestService {
	return &accountService{repo}
}

func (a *accountService) List(c *iris.Context) {
	res, err := a.DomainService.Find()
	if err != nil {
		c.JSON(iris.StatusOK, err)
	}
	c.JSON(iris.StatusOK, res)
}
