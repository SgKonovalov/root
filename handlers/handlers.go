package handlers

import (
	"net/http"
	"strconv"

	"root/models"
	"root/models/advertisement"
	"root/models/responses"
	"root/service"

	"github.com/gin-gonic/gin"
)

//Хендлеры приложения
type Handler struct {
	Service models.GeneralService
	Logger  *models.Logger
}

//Функция установки заголовков для ответов json.
func (h Handler) InitHeaders(c *gin.Context) {
	c.Header(ContentType, ApplicationJson)
}

//Функция домашней страницы.
func (h Handler) Home(c *gin.Context) {
	h.InitHeaders(c)
}

//Функция получения всех объявлений.
func (h Handler) GetAdvsList(c *gin.Context) {

	h.InitHeaders(c)

	if c.Request.Method != http.MethodGet {
		h.Logger.Error(ErrorOccurredAt, GetAdvsListFuncName, WrongMethodTypeGET)
		errReq := responses.NewResponse(Error,
			responses.NewMetaError(http.StatusMethodNotAllowed, WrongMethodTypeGET),
			nil)
		c.IndentedJSON(http.StatusMethodNotAllowed, errReq)
		return
	}

	data, err := h.Service.GetAdvsList()

	if err != nil {
		errRes := responses.NewResponse(Error,
			responses.NewMetaError(http.StatusInternalServerError, err.Error()),
			nil)
		c.IndentedJSON(http.StatusInternalServerError, errRes)
		h.Logger.Error(service.ErrorOccurredAt, GetAdvsListFuncName, Reason, err)
		return
	}

	successRes := responses.NewResponse(Success, responses.MetaSuccess{}, data)
	c.IndentedJSON(http.StatusOK, successRes)
	h.Logger.Info(GetAdvsListFuncName, LoggerSuccessOperation)
}

//Функция получения всех объявлений.
func (h Handler) GetOneAdv(c *gin.Context) {

	h.InitHeaders(c)

	if c.Request.Method != http.MethodGet {
		h.Logger.Error(ErrorOccurredAt, GetOneAdvFuncName, WrongMethodTypeGET)
		errReq := responses.NewResponse(Error,
			responses.NewMetaError(http.StatusMethodNotAllowed, WrongMethodTypeGET),
			nil)
		c.IndentedJSON(http.StatusMethodNotAllowed, errReq)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		errRes := responses.NewResponse(Error,
			responses.NewMetaError(http.StatusInternalServerError, err.Error()),
			nil)
		c.IndentedJSON(http.StatusInternalServerError, errRes)
		h.Logger.Error(service.ErrorOccurredAt, GetOneAdvFuncName, Reason, err)
		return
	}

	fPar := c.Request.Header.Get("fields")

	var field int

	if fPar != "" {
		field = 1
	} else {
		field = 2
	}

	data, err := h.Service.GetOneAdv(id, field)

	if err != nil {
		errRes := responses.NewResponse(Error,
			responses.NewMetaError(http.StatusInternalServerError, err.Error()),
			nil)
		c.IndentedJSON(http.StatusInternalServerError, errRes)
		h.Logger.Error(service.ErrorOccurredAt, GetOneAdvFuncName, Reason, err)
		return
	}

	successRes := responses.NewResponse(Success, responses.MetaSuccess{}, data)
	c.IndentedJSON(http.StatusOK, successRes)
	h.Logger.Info(GetOneAdvFuncName, LoggerSuccessOperation)
}

//Функция добавления нового объявления.
func (h Handler) AddNewAdv(c *gin.Context) {

	h.InitHeaders(c)

	if c.Request.Method != http.MethodPost {
		h.Logger.Error(ErrorOccurredAt, AddNewAdvFuncName, WrongMethodTypePOST)
		errReq := responses.NewResponse(Error,
			responses.NewMetaError(http.StatusMethodNotAllowed, WrongMethodTypePOST),
			nil)
		c.IndentedJSON(http.StatusMethodNotAllowed, errReq)
		return
	}

	addThis := new(advertisement.AddNewAdvertisement)

	if err := c.BindJSON(&addThis); err != nil {
		errRes := responses.NewResponse(Error,
			responses.NewMetaError(http.StatusInternalServerError, err.Error()),
			nil)
		c.IndentedJSON(http.StatusInternalServerError, errRes)
		h.Logger.Error(service.ErrorOccurredAt, AddNewAdvFuncName, Reason, err)
		return
	}

	data, err := h.Service.AddNewAdv(*addThis)

	if err != nil {
		errRes := responses.NewResponse(Error,
			responses.NewMetaError(http.StatusInternalServerError, err.Error()),
			nil)
		c.IndentedJSON(http.StatusInternalServerError, errRes)
		h.Logger.Error(service.ErrorOccurredAt, AddNewAdvFuncName, Reason, err)
		return
	}

	successRes := responses.NewResponse(Success, responses.MetaSuccess{}, data)
	c.IndentedJSON(http.StatusOK, successRes)
	h.Logger.Info(AddNewAdvFuncName, LoggerSuccessOperation)
}

//Функция добавления фотографии.
func (h Handler) AddNewPhoto(c *gin.Context) {

	h.InitHeaders(c)

	if c.Request.Method != http.MethodPut {
		h.Logger.Error(ErrorOccurredAt, AddNewPhotoFuncName, WrongMethodTypePUT)
		errReq := responses.NewResponse(Error,
			responses.NewMetaError(http.StatusMethodNotAllowed, WrongMethodTypePUT),
			nil)
		c.IndentedJSON(http.StatusMethodNotAllowed, errReq)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		errRes := responses.NewResponse(Error,
			responses.NewMetaError(http.StatusInternalServerError, err.Error()),
			nil)
		c.IndentedJSON(http.StatusInternalServerError, errRes)
		h.Logger.Error(service.ErrorOccurredAt, AddNewPhotoFuncName, Reason, err)
		return
	}

	file, handler, err := c.Request.FormFile(PhotoHeader)

	if err != nil {
		errRes := responses.NewResponse(Error,
			responses.NewMetaError(http.StatusInternalServerError, err.Error()),
			nil)
		c.IndentedJSON(http.StatusInternalServerError, errRes)
		h.Logger.Error(service.ErrorOccurredAt, AddNewPhotoFuncName, Reason, err)
		return
	}

	defer file.Close()

	isMain := c.Request.Header.Get(IsMainPhoto)

	var mainType bool

	if isMain == "true" {
		mainType = true
	} else {
		mainType = false
	}

	if err := h.Service.AddNewPhoto(id, handler.Filename, file, mainType); err != nil {
		errRes := responses.NewResponse(Error,
			responses.NewMetaError(http.StatusInternalServerError, err.Error()),
			nil)
		c.IndentedJSON(http.StatusInternalServerError, errRes)
		h.Logger.Error(service.ErrorOccurredAt, AddNewPhotoFuncName, Reason, err)
		return
	}

	c.Writer.WriteHeader(http.StatusOK)
	h.Logger.Info(AddNewPhotoFuncName, LoggerSuccessOperation)
}
