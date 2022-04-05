package responses

import "root/models/advertisement"

//Стрктура для ответа от API (успешный ответ поле Data). Список всех объявлений.
type ListOfAdvsSuccessAPIresponse struct {
	AdvertisementList []advertisement.OneAdvForList `json:"all_advs"`
}

//Стрктура для ответа от API (успешный ответ поле Data). 1 объвление - развёрнутая структура.
type OneAdvSuccessAPIresponse struct {
	Advertisement advertisement.ShortAdvertisement `json:"adv_long"`
}

//Стрктура для ответа от API (успешный ответ поле Data). Добавление нового объявления.
type AddNewAdvSuccessAPIresponse struct {
	ID int `json:"adv_id"`
}
