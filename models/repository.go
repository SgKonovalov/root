package models

import (
	"mime/multipart"
	"root/models/advertisement"
)

//Общий интерфейс для Репозитория
type GeneralRepo interface {
	GetAdvsList() ([]advertisement.OneAdvForList, error)
	GetOneAdv(id int) (*advertisement.ShortAdvOne, error)
	GetTwoAdv(id int) (*advertisement.ShortAdvTwo, error)
	AddNewAdv(adv advertisement.AddNewAdvertisement) (int, error)
	AddNewPhoto(advID int, fileName string, photo multipart.File, isMain bool) error
}
