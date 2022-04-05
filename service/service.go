package service

import (
	"context"
	"mime/multipart"

	"root/helpers"
	"root/models"
	"root/models/advertisement"
	"root/models/responses"
)

//Реализация №1.
type Service struct {
	Repository models.GeneralRepo
	Logger     *models.Logger
	Context    context.Context
	Helper     *helpers.Helper
}

//Функции сервиса

//Функция показа всех объявлений.
func (so Service) GetAdvsList() (*responses.ListOfAdvsSuccessAPIresponse, error) {

	allAdvs, err := so.Repository.GetAdvsList()

	if err != nil {
		so.Logger.Error(ErrorOccurredAt, GetAllGetAdvsListFuncName, Reason, err)
		return nil, err
	}

	advsForResp, err := so.Helper.Converter.AdvsListFromGeneralToAPIresponse(allAdvs)

	if err != nil {
		so.Logger.Error(ErrorOccurredAt, GetAllGetAdvsListFuncName, Reason, err)
		return nil, err
	}

	return advsForResp, nil
}

//Функция показа одного объявления.
func (so Service) GetOneAdv(id, resStruct int) (*responses.OneAdvSuccessAPIresponse, error) {

	switch resStruct {
	case 1:
		adv, err := so.Repository.GetOneAdv(id)

		if err != nil {
			so.Logger.Error(ErrorOccurredAt, GetOneAdvFuncName, Reason, err)
			return nil, err
		}

		return &responses.OneAdvSuccessAPIresponse{
			Advertisement: adv,
		}, nil

	case 2:
		adv, err := so.Repository.GetTwoAdv(id)

		if err != nil {
			so.Logger.Error(ErrorOccurredAt, GetOneAdvFuncName, Reason, err)
			return nil, err
		}

		return &responses.OneAdvSuccessAPIresponse{
			Advertisement: adv,
		}, nil

	}
	return nil, nil
}

//Функция добавления одного объявления.
func (so Service) AddNewAdv(adv advertisement.AddNewAdvertisement) (*responses.AddNewAdvSuccessAPIresponse, error) {

	id, err := so.Repository.AddNewAdv(adv)

	if err != nil {
		so.Logger.Error(ErrorOccurredAt, AddNewAdvFuncName, Reason, err)
		return nil, err
	}

	return &responses.AddNewAdvSuccessAPIresponse{
		ID: id,
	}, nil

}

//Функция добавления фотографии в объявление.
func (so Service) AddNewPhoto(advID int, fileName string, photo multipart.File, isMain bool) error {

	if err := so.Repository.AddNewPhoto(advID, fileName, photo, isMain); err != nil {
		so.Logger.Error(ErrorOccurredAt, AddNewPhotoFuncName, Reason, err)
		return err
	}

	return nil
}
