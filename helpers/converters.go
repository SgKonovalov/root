package helpers

import (
	"errors"

	"root/models"
	"root/models/advertisement"
	"root/models/responses"
)

//Converter - для перевода данных 1-го типа к другому
type Converter struct {
	Logger *models.Logger
}

//Переводит все объявления в тип, необходимый для ответа в API
func (con *Converter) AdvsListFromGeneralToAPIresponse(source []advertisement.OneAdvForList) (*responses.ListOfAdvsSuccessAPIresponse, error) {

	allAdvs := new(responses.ListOfAdvsSuccessAPIresponse)

	if len(source) == 0 {
		con.Logger.Error(ErrorOccurredAt, AdvsListFromGeneralToAPIresponseFuncName, Reason, EmptySourse)
		return allAdvs, errors.New(EmptySourse)
	}

	allAdvs.AdvertisementList = source

	return allAdvs, nil
}
