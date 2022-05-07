package mock

import (
	"github.com/bxcodec/faker/v3"
	"github.com/pkg/errors"
	"github.com/samithiwat/samithiwat-backend-general/src/model"
	"github.com/samithiwat/samithiwat-backend-general/src/proto"
	"gorm.io/gorm"
)

var Loc1 model.Location
var Loc2 model.Location
var Loc3 model.Location
var Loc4 model.Location
var Locs []*model.Location
var CreateLocationReqMock proto.CreateLocationRequest
var UpdateLocationReqMock proto.UpdateLocationRequest

type LocationMockRepo struct {
}

func (LocationMockRepo) FindOne(_ int, loc *model.Location) error {
	*loc = Loc1
	return nil
}

func (LocationMockRepo) Create(loc *model.Location) error {
	*loc = Loc1
	return nil
}

func (LocationMockRepo) Update(_ int, loc *model.Location) error {
	*loc = Loc1
	return nil
}

func (LocationMockRepo) Delete(_ int, loc *model.Location) error {
	*loc = Loc1
	return nil
}

type LocationMockErrRepo struct {
}

func (LocationMockErrRepo) FindOne(int, *model.Location) error {
	return errors.New("Not found location")
}

func (LocationMockErrRepo) Create(*model.Location) error {
	return nil
}

func (LocationMockErrRepo) Update(int, *model.Location) error {
	return errors.New("Not found location")
}

func (LocationMockErrRepo) Delete(int, *model.Location) error {
	return errors.New("Not found location")
}

func InitializeMockLocation() (err error) {
	Loc1 = model.Location{
		Model:    gorm.Model{ID: 1},
		Address:  faker.Word(),
		District: faker.Word(),
		Province: faker.Word(),
		Country:  faker.Word(),
		ZipCode:  faker.Word(),
	}

	Loc2 = model.Location{
		Model:    gorm.Model{ID: 2},
		Address:  faker.Word(),
		District: faker.Word(),
		Province: faker.Word(),
		Country:  faker.Word(),
		ZipCode:  faker.Word(),
	}

	Loc3 = model.Location{
		Model:    gorm.Model{ID: 3},
		Address:  faker.Word(),
		District: faker.Word(),
		Province: faker.Word(),
		Country:  faker.Word(),
		ZipCode:  faker.Word(),
	}

	Loc4 = model.Location{
		Model:    gorm.Model{ID: 4},
		Address:  faker.Word(),
		District: faker.Word(),
		Province: faker.Word(),
		Country:  faker.Word(),
		ZipCode:  faker.Word(),
	}

	CreateLocationReqMock = proto.CreateLocationRequest{
		Location: &proto.Location{
			Address:  faker.Word(),
			District: faker.Word(),
			Province: faker.Word(),
			Country:  faker.Word(),
			Zipcode:  faker.Word(),
		},
	}
	if err != nil {
		panic("Error occur while mocking data")
	}

	UpdateLocationReqMock = proto.UpdateLocationRequest{
		Location: &proto.Location{
			Id:       uint32(Loc1.ID),
			Address:  faker.Word(),
			District: faker.Word(),
			Province: faker.Word(),
			Country:  faker.Word(),
			Zipcode:  faker.Word(),
		},
	}
	if err != nil {
		panic("Error occur while mocking data")
	}

	Locs = append(Locs, &Loc1, &Loc2, &Loc3, &Loc4)

	return nil
}
