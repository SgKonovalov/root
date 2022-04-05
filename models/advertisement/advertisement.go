package advertisement

import "time"

//Понятие объявления.
type Advertisement struct {
	ID                  int       `json:"id"`
	NameOfAdvertisement string    `json:"adv_name"`
	Price               int       `json:"price"`
	PublishDate         time.Time `json:"pub_date"`
	Description         string    `json:"description"`
	MainPhoto           Photo     `json:"main_photo"`
	Photos              []Photo   `json:"all_photos"`
}

//Понятие фотографии.
type Photo struct {
	PhotosURL string `json:"photo"`
}

//Каждое объявление для списка
type OneAdvForList struct {
	NameOfAdvertisement string `json:"adv_name"`
	MainPhoto           Photo  `json:"main_photo"`
	Price               int    `json:"price"`
}

//Отдельное обявление (развёртнутая информация).
type ShortAdvertisement interface {
	GetAdvName() string
}

type ShortAdvOne struct {
	NameOfAdvertisement string  `json:"adv_name"`
	Price               int     `json:"price"`
	Description         string  `json:"description"`
	Photos              []Photo `json:"all_photos"`
}

func (sao *ShortAdvOne) GetAdvName() string {
	return sao.NameOfAdvertisement
}

type ShortAdvTwo struct {
	NameOfAdvertisement string `json:"adv_name"`
	Price               int    `json:"price"`
}

func (sat *ShortAdvTwo) GetAdvName() string {
	return sat.NameOfAdvertisement
}

//Создание нового объявления
type AddNewAdvertisement struct {
	NameOfAdvertisement string `json:"adv_name"`
	Price               int    `json:"price"`
	Description         string `json:"description"`
}
