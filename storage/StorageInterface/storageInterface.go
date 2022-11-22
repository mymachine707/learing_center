package storageinterface

import "mymachine707/models"

// StorageInterface ...
type StorageInterface interface {
	// For Person
	AddPerson(id string, entity models.CreatePersonModul) error
	GetPersonByID(id string) (models.Person, error)
	GetPersonList(offset, limit int, search string) (resp []models.Person, err error)
	UpdatePerson(person models.UpdatePersonModel) error
	DeletePerson(id string) error
}
