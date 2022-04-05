package repository

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"time"

	"root/models"
	"root/models/advertisement"
)

const (
	PhotosPlaceHolder = "./photos"
)

//Реализация №1.
type Repository struct {
	DataBase models.GeneralDB
	Context  context.Context
}

//Функции репозитория

//Функция показа всех объявлений.
func (ro Repository) GetAdvsList() ([]advertisement.OneAdvForList, error) {

	var advList []advertisement.OneAdvForList

	dbPool, err := ro.DataBase.GetConnection()
	if err != nil {
		return advList, err
	}

	defer dbPool.Close()

	fromDB, err := dbPool.Query(ro.Context, GetAdvsList)

	if err != nil {
		return advList, err
	}

	var allAdvs []advertisement.OneAdvForList
	var adv advertisement.OneAdvForList

	for fromDB.Next() {

		if err := fromDB.Scan(&adv.NameOfAdvertisement,
			&adv.Price,
			&adv.MainPhoto.PhotosURL); err != nil {
			return advList, err
		}

		allAdvs = append(allAdvs, adv)

	}

	advList = allAdvs

	return advList, nil
}

//Функция показа одного объявления (длинная версия).
func (ro Repository) GetOneAdv(id int) (*advertisement.ShortAdvOne, error) {

	adv := new(advertisement.ShortAdvOne)
	var photo advertisement.Photo
	var allPhotos []advertisement.Photo

	dbPool, err := ro.DataBase.GetConnection()
	if err != nil {
		return adv, err
	}

	defer dbPool.Close()

	fromDB, err := dbPool.Query(ro.Context, LongGetOnePhoto, id)

	if err != nil {
		return adv, err
	}

	for fromDB.Next() {

		if err := fromDB.Scan(&adv.NameOfAdvertisement,
			&adv.Price,
			&adv.Description,
			&photo.PhotosURL,
		); err != nil {
			return adv, err
		}

		allPhotos = append(allPhotos, photo)
	}

	return &advertisement.ShortAdvOne{
		NameOfAdvertisement: adv.NameOfAdvertisement,
		Price:               adv.Price,
		Description:         adv.Description,
		Photos:              allPhotos,
	}, nil
}

//Функция показа одного объявления (короткая версия).
func (ro Repository) GetTwoAdv(id int) (*advertisement.ShortAdvTwo, error) {

	adv := new(advertisement.ShortAdvTwo)

	dbPool, err := ro.DataBase.GetConnection()
	if err != nil {
		return adv, err
	}

	defer dbPool.Close()

	fromDB, err := dbPool.Query(ro.Context, ShortGetOnePhoto, id)

	if err != nil {
		return adv, err
	}

	for fromDB.Next() {

		if err := fromDB.Scan(&adv.NameOfAdvertisement, &adv.Price); err != nil {
			return adv, err
		}
	}

	return &advertisement.ShortAdvTwo{
		NameOfAdvertisement: adv.NameOfAdvertisement,
		Price:               adv.Price,
	}, nil
}

//Функция добавления нового объявления.
func (ro Repository) AddNewAdv(adv advertisement.AddNewAdvertisement) (int, error) {

	var id int

	dbPool, err := ro.DataBase.GetConnection()
	if err != nil {
		return 0, err
	}

	defer dbPool.Close()

	fromDB := dbPool.QueryRow(ro.Context, AddNewAdv, adv.NameOfAdvertisement,
		adv.Description, adv.Price, time.Now())

	if err := fromDB.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

//Функция добавления фотографии в объявление.
func (ro Repository) AddNewPhoto(advID int, fileName string, photo multipart.File, isMain bool) error {

	filePath := fmt.Sprintf("%s/%s", PhotosPlaceHolder, fmt.Sprint(advID, fileName))

	photoFromReq, err := os.Create(filePath)

	if err != nil {
		return err
	}

	defer photoFromReq.Close()

	if _, err := io.Copy(photoFromReq, photo); err != nil {
		return err
	}

	dbPool, err := ro.DataBase.GetConnection()
	if err != nil {
		return err
	}

	defer dbPool.Close()

	if _, err := dbPool.Exec(ro.Context, PhotoAdd, advID, filePath, isMain); err != nil {
		return err
	}

	return nil
}
