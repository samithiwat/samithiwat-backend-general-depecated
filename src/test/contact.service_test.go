package test

import (
	"github.com/samithiwat/samithiwat-backend-general/src/proto"
	"github.com/samithiwat/samithiwat-backend-general/src/service"
	"github.com/samithiwat/samithiwat-backend-general/src/test/mock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestFindOneContact(t *testing.T) {
	mock.InitializeMockContact()

	var errors []string

	assert := assert.New(t)
	want := &proto.ContactResponse{
		Data:       RawToDtoContact(&mock.Cont1),
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	contService := service.NewContactService(&mock.ContactMockRepo{})
	contRes, err := contService.FindOne(mock.Context{}, &proto.FindOneContactRequest{Id: 1})
	if err != nil {
		t.Errorf("Got an error")
	}

	assert.Equal(want, contRes)
}

func TestFindOneErrNotFoundContact(t *testing.T) {
	mock.InitializeMockContact()

	errors := []string{"Not found contact"}

	assert := assert.New(t)
	want := &proto.ContactResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusNotFound,
	}

	contService := service.NewContactService(&mock.ContactMockErrRepo{})
	contRes, err := contService.FindOne(mock.Context{}, &proto.FindOneContactRequest{Id: 1})

	assert.True(err != nil, "Must got an error")
	assert.Equal(want, contRes)
}

func TestCreateContact(t *testing.T) {
	mock.InitializeMockContact()

	var errors []string

	assert := assert.New(t)
	want := &proto.ContactResponse{
		Data:       RawToDtoContact(&mock.Cont1),
		Errors:     errors,
		StatusCode: http.StatusCreated,
	}

	contService := service.NewContactService(&mock.ContactMockRepo{})
	contRes, err := contService.Create(mock.Context{}, &mock.CreateContactReqMock)
	if err != nil {
		t.Errorf("Got an error")
	}

	assert.Equal(want, contRes)
}

func TestUpdateContact(t *testing.T) {
	mock.InitializeMockContact()

	var errors []string

	assert := assert.New(t)
	want := &proto.ContactResponse{
		Data:       RawToDtoContact(&mock.Cont1),
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	contService := service.NewContactService(&mock.ContactMockRepo{})
	contRes, err := contService.Update(mock.Context{}, &mock.UpdateContactReqMock)
	if err != nil {
		t.Errorf("Got an error")
	}

	assert.Equal(want, contRes)
}

func TestUpdateErrNotFoundContact(t *testing.T) {
	mock.InitializeMockContact()

	errors := []string{"Not found contact"}

	assert := assert.New(t)

	want := &proto.ContactResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusNotFound,
	}

	contService := service.NewContactService(&mock.ContactMockErrRepo{})
	contRes, err := contService.Update(mock.Context{}, &mock.UpdateContactReqMock)

	assert.True(err != nil, "Must got an error")
	assert.Equal(want, contRes)
}

func TestDeleteContact(t *testing.T) {
	mock.InitializeMockContact()

	var errors []string

	assert := assert.New(t)
	want := &proto.ContactResponse{
		Data:       RawToDtoContact(&mock.Cont1),
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	contService := service.NewContactService(&mock.ContactMockRepo{})
	contRes, err := contService.Delete(mock.Context{}, &proto.DeleteContactRequest{Id: 1})
	if err != nil {
		t.Errorf("Got an error")
	}

	assert.Equal(want, contRes)
}

func TestDeleteErrNotFoundContact(t *testing.T) {
	mock.InitializeMockContact()

	errors := []string{"Not found contact"}

	assert := assert.New(t)

	want := &proto.ContactResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusNotFound,
	}

	contService := service.NewContactService(&mock.ContactMockErrRepo{})
	contRes, err := contService.Delete(mock.Context{}, &proto.DeleteContactRequest{Id: 1})

	assert.True(err != nil, "Must got an error")
	assert.Equal(want, contRes)
}
