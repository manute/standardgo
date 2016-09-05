package main

import (
	"mytests/api/db"
	"mytests/api/http"

	"github.com/iris-contrib/logger"
	mLogger "github.com/iris-contrib/middleware/logger"
	"github.com/kataras/iris"
)

func main() {
	db, err := db.NewPSQL()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	accountRestService := http.NewAccountRestService(db)

	theLogger := logger.New(logger.DefaultConfig())

	iris.Use(mLogger.New(theLogger))

	iris.Get("/accounts", accountRestService.GetAll)

	iris.Listen(":8080")
}
