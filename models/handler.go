package models

import "github.com/gin-gonic/gin"

//Общий интерфейс для всех хендеров приложения
type GeneralHandler interface {
	Home(c *gin.Context)
	GetAdvsList(c *gin.Context)
	GetOneAdv(c *gin.Context)
	AddNewAdv(c *gin.Context)
	AddNewPhoto(c *gin.Context)
}
