package models

import (
	"mime/multipart"
	"root/models/advertisement"
	"root/models/responses"
)

//Общий интерфейс сервиса (обрабатывает результаты, полученные от репозитория)
type GeneralService interface {
	GetAdvsList() (*responses.ListOfAdvsSuccessAPIresponse, error)
	GetOneAdv(id, resStruct int) (*responses.OneAdvSuccessAPIresponse, error)
	AddNewAdv(adv advertisement.AddNewAdvertisement) (*responses.AddNewAdvSuccessAPIresponse, error)
	AddNewPhoto(advID int, fileName string, photo multipart.File, isMain bool) error
}
