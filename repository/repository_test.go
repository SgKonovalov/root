package repository

import (
	"context"
	"root/models"
	"root/models/advertisement"
	"testing"
)

func TestGetAdvsList(t *testing.T) {

	configuration, err := models.NewMainConfig()

	if err != nil {
		t.Errorf("Configuration is not working. Reason is %v", err)
	}

	dataBase := models.NewPostgreDBbyPGX(configuration, context.Background())

	repo := Repository{
		DataBase: dataBase,
		Context:  context.Background(),
	}

	result, err := repo.GetAdvsList()

	if err != nil {
		t.Errorf("GetAdvsList is not working. Reason is %v", err)
	}

	if len(result) == 0 {
		t.Errorf("GetAdvsList is not working. Reason len(result) == %d", len(result))
	}

}

func TestGetAdv(t *testing.T) {

	configuration, err := models.NewMainConfig()

	if err != nil {
		t.Errorf("Configuration is not working. Reason is %v", err)
	}

	dataBase := models.NewPostgreDBbyPGX(configuration, context.Background())

	repo := Repository{
		DataBase: dataBase,
		Context:  context.Background(),
	}

	resultByLong, err := repo.GetOneAdv(1)

	if err != nil {
		t.Errorf("GetOneAdv is not working. Reason is %v", err)
	}

	if resultByLong.NameOfAdvertisement == "" {
		t.Errorf("GetOneAdv is not working. Reason result.NameOfAdvertisement == %s", resultByLong.NameOfAdvertisement)
	}

	resultByShort, err := repo.GetTwoAdv(1)

	if err != nil {
		t.Errorf("GetTwoAdv is not working. Reason is %v", err)
	}

	if resultByShort.NameOfAdvertisement == "" {
		t.Errorf("GetTwoAdv is not working. Reason result.NameOfAdvertisement == %s", resultByShort.NameOfAdvertisement)
	}

}

func TestAddNewAdv(t *testing.T) {

	configuration, err := models.NewMainConfig()

	if err != nil {
		t.Errorf("Configuration is not working. Reason is %v", err)
	}

	dataBase := models.NewPostgreDBbyPGX(configuration, context.Background())

	repo := Repository{
		DataBase: dataBase,
		Context:  context.Background(),
	}

	id, err := repo.AddNewAdv(advertisement.AddNewAdvertisement{
		NameOfAdvertisement: "test",
		Price:               100,
		Description:         "Test",
	})

	if err != nil {
		t.Errorf("AddNewAdv is not working. Reason is %v", err)
	}

	if id <= 0 {
		t.Errorf("AddNewAdv is not working. Reason is id == %d", id)
	}

}
