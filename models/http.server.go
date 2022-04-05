package models

import "github.com/gin-gonic/gin"

//HTTP сервер.
type HTTPserver struct {
	Engine  *gin.Engine
	Handler GeneralHandler
}
