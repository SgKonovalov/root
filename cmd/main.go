package main

import (
	"context"

	"root/handlers"
	"root/helpers"
	"root/models"
	"root/repository"
	"root/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {

	configuration, err := models.NewMainConfig()

	if err != nil {
		return
	}

	mainContext := context.Background()
	cWcan, _ := context.WithCancel(mainContext)

	dataBase := models.NewPostgreDBbyPGX(configuration, cWcan)

	loggerApp := models.NewLogger(logrus.New(), logrus.New())

	repo := repository.Repository{
		DataBase: dataBase,
		Context:  cWcan,
	}

	helper := &helpers.Helper{
		Logger: loggerApp,
		Converter: &helpers.Converter{
			Logger: loggerApp,
		},
	}

	serv := service.Service{
		Repository: repo,
		Logger:     loggerApp,
		Context:    cWcan,
		Helper:     helper,
	}

	handl := handlers.Handler{
		Service: serv,
		Logger:  loggerApp,
	}

	server := models.HTTPserver{
		Engine:  gin.Default(),
		Handler: handl,
	}

	api := models.API{
		Cofiguration: configuration,
		HTTPserver:   &server,
	}

	api.HTTPserver.Engine.GET(handlers.HomeURL, api.HTTPserver.Handler.Home)
	api.HTTPserver.Engine.GET(handlers.GetAdvsListURL, api.HTTPserver.Handler.GetAdvsList)
	api.HTTPserver.Engine.GET(handlers.GetOneAdvURL, api.HTTPserver.Handler.GetOneAdv)
	api.HTTPserver.Engine.POST(handlers.AddNewAdvURL, api.HTTPserver.Handler.AddNewAdv)
	api.HTTPserver.Engine.PUT(handlers.AddPhotoAtAdvURL, api.HTTPserver.Handler.AddNewPhoto)

	if err := api.HTTPserver.Engine.Run(api.Cofiguration.HTTPserverAddr); err != nil {
		loggerApp.Error(err)
	}
}
