package http

import (
	"mytests/api/db"
	"mytests/api/domain"
	"mytests/api/repository"

	"github.com/kataras/iris"
)

type AccountRestService struct {
	DomainService domain.AccountService
}

func NewAccountRestService(db *db.DB) *AccountRestService {
	return &AccountRestService{
		&repository.AccountRepository{Db: db},
	}
}

func (a *AccountRestService) GetAll(c *iris.Context) {
	res, err := a.DomainService.Find()
	if err != nil {
		c.JSON(iris.StatusOK, err)
	}
	c.JSON(iris.StatusOK, res)
}
