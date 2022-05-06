package service

import "github.com/samithiwat/samithiwat-backend-general/src/model"

type ContactService struct {
	repository ContactRepository
}

type ContactRepository interface {
	FindOne(int, *model.Contact) error
	Create(*model.Contact) error
	Update(int, *model.Contact) error
	Delete(int, *model.Contact) error
}

func NewContactService(repository ContactRepository) *ContactService {
	return &ContactService{
		repository: repository,
	}
}
