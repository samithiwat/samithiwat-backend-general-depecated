package test

import (
	"github.com/samithiwat/samithiwat-backend-general/src/proto"
	"github.com/samithiwat/samithiwat-backend-general/src/service"
	"github.com/samithiwat/samithiwat-backend-general/src/test/mock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestFindOneLocation(t *testing.T) {
	mock.InitializeMockLocation()

	var errors []string

	assert := assert.New(t)
	want := &proto.LocationResponse{
		Data:       RawToDtoLocation(&mock.Loc1),
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	locService := service.NewLocationService(&mock.LocationMockRepo{})
	locRes, err := locService.FindOne(mock.Context{}, &proto.FindOneLocationRequest{Id: 1})
	if err != nil {
		t.Errorf("Got an error")
	}

	assert.Equal(want, locRes)
}

func TestFindOneErrNotFoundLocation(t *testing.T) {
	mock.InitializeMockLocation()

	errors := []string{"Not found location"}

	assert := assert.New(t)
	want := &proto.LocationResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusNotFound,
	}

	locService := service.NewLocationService(&mock.LocationMockErrRepo{})
	locRes, err := locService.FindOne(mock.Context{}, &proto.FindOneLocationRequest{Id: 1})

	assert.True(err != nil, "Must got an error")
	assert.Equal(want, locRes)
}

func TestCreateLocation(t *testing.T) {
	mock.InitializeMockLocation()

	var errors []string

	assert := assert.New(t)
	want := &proto.LocationResponse{
		Data:       RawToDtoLocation(&mock.Loc1),
		Errors:     errors,
		StatusCode: http.StatusCreated,
	}

	locService := service.NewLocationService(&mock.LocationMockRepo{})
	locRes, err := locService.Create(mock.Context{}, &mock.CreateLocationReqMock)
	if err != nil {
		t.Errorf("Got an error")
	}

	assert.Equal(want, locRes)
}

func TestUpdateLocation(t *testing.T) {
	mock.InitializeMockLocation()

	var errors []string

	assert := assert.New(t)
	want := &proto.LocationResponse{
		Data:       RawToDtoLocation(&mock.Loc1),
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	locService := service.NewLocationService(&mock.LocationMockRepo{})
	locRes, err := locService.Update(mock.Context{}, &mock.UpdateLocationReqMock)
	if err != nil {
		t.Errorf("Got an error")
	}

	assert.Equal(want, locRes)
}

func TestUpdateErrNotFoundLocation(t *testing.T) {
	mock.InitializeMockLocation()

	errors := []string{"Not found location"}

	assert := assert.New(t)

	want := &proto.LocationResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusNotFound,
	}

	locService := service.NewLocationService(&mock.LocationMockErrRepo{})
	locRes, err := locService.Update(mock.Context{}, &mock.UpdateLocationReqMock)

	assert.True(err != nil, "Must got an error")
	assert.Equal(want, locRes)
}

func TestDeleteLocation(t *testing.T) {
	mock.InitializeMockLocation()

	var errors []string

	assert := assert.New(t)
	want := &proto.LocationResponse{
		Data:       RawToDtoLocation(&mock.Loc1),
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	locService := service.NewLocationService(&mock.LocationMockRepo{})
	locRes, err := locService.Delete(mock.Context{}, &proto.DeleteLocationRequest{Id: 1})
	if err != nil {
		t.Errorf("Got an error")
	}

	assert.Equal(want, locRes)
}

func TestDeleteErrNotFoundLocation(t *testing.T) {
	mock.InitializeMockLocation()

	errors := []string{"Not found location"}

	assert := assert.New(t)

	want := &proto.LocationResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusNotFound,
	}

	locService := service.NewLocationService(&mock.LocationMockErrRepo{})
	locRes, err := locService.Delete(mock.Context{}, &proto.DeleteLocationRequest{Id: 1})

	assert.True(err != nil, "Must got an error")
	assert.Equal(want, locRes)
}
