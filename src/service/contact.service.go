package service

import (
	"context"
	"github.com/samithiwat/samithiwat-backend-general/src/model"
	"github.com/samithiwat/samithiwat-backend-general/src/proto"
	"net/http"
)

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

func (s *ContactService) FindOne(_ context.Context, req *proto.FindOneContactRequest) (res *proto.ContactResponse, err error) {
	cont := model.Contact{}
	var errors []string

	res = &proto.ContactResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	err = s.repository.FindOne(int(req.Id), &cont)
	if err != nil {
		res.Errors = append(errors, err.Error())
		res.StatusCode = http.StatusNotFound
		return
	}

	result := RawToDtoContact(&cont)
	res.Data = result
	return
}

func (s *ContactService) Create(_ context.Context, req *proto.CreateContactRequest) (res *proto.ContactResponse, err error) {
	cont := DtoToRawContact(req.Contact)
	var errors []string

	res = &proto.ContactResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusCreated,
	}

	err = s.repository.Create(cont)
	if err != nil {
		res.Errors = append(errors, err.Error())
		res.StatusCode = http.StatusUnprocessableEntity
		return
	}

	result := RawToDtoContact(cont)
	res.Data = result

	return
}

func (s *ContactService) Update(_ context.Context, req *proto.UpdateContactRequest) (res *proto.ContactResponse, err error) {
	cont := DtoToRawContact(req.Contact)
	var errors []string

	res = &proto.ContactResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	err = s.repository.Update(int(cont.ID), cont)
	if err != nil {
		res.Errors = append(errors, err.Error())
		res.StatusCode = http.StatusNotFound
		return
	}

	result := RawToDtoContact(cont)
	res.Data = result

	return
}

func (s *ContactService) Delete(_ context.Context, req *proto.DeleteContactRequest) (res *proto.ContactResponse, err error) {
	cont := model.Contact{}
	var errors []string

	res = &proto.ContactResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	err = s.repository.Delete(int(req.Id), &cont)
	if err != nil {
		res.Errors = append(errors, err.Error())
		res.StatusCode = http.StatusNotFound
		return
	}

	result := RawToDtoContact(&cont)
	res.Data = result

	return
}
