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

	// For Lesson
	AddLesson(id string, entity models.CreateLessonModels) error
	GetLessonByID(id string) (models.Lesson, error)
	GetListLesson(offset, limit int, search string) (resp []models.Lesson, err error)
	UpdateLessonByID(person models.UpdateLessonModels) error
	DeleteLessonByID(id string) error
}
