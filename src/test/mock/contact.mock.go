package mock

import (
	"github.com/bxcodec/faker/v3"
	"github.com/pkg/errors"
	"github.com/samithiwat/samithiwat-backend-general/src/model"
	"github.com/samithiwat/samithiwat-backend-general/src/proto"
	"gorm.io/gorm"
)

var Cont1 model.Contact
var Cont2 model.Contact
var Cont3 model.Contact
var Cont4 model.Contact
var Conts []*model.Contact
var CreateContactReqMock proto.CreateContactRequest
var UpdateContactReqMock proto.UpdateContactRequest

type ContactMockRepo struct {
}

func (r *ContactMockRepo) FindOne(id int, cont *model.Contact) error {
	*cont = Cont1
	return nil
}

func (r *ContactMockRepo) Create(cont *model.Contact) error {
	*cont = Cont1
	return nil
}

func (r *ContactMockRepo) Update(id int, cont *model.Contact) error {
	*cont = Cont1
	return nil
}

func (r *ContactMockRepo) Delete(id int, cont *model.Contact) error {
	*cont = Cont1
	return nil
}

type ContactMockErrRepo struct {
}

func (r *ContactMockErrRepo) FindOne(id int, cont *model.Contact) error {
	return errors.New("Not found contact")
}

func (r *ContactMockErrRepo) Create(cont *model.Contact) error {
	return nil
}

func (r *ContactMockErrRepo) Update(id int, cont *model.Contact) error {
	return errors.New("Not found contact")
}

func (r *ContactMockErrRepo) Delete(id int, cont *model.Contact) error {
	return errors.New("Not found contact")
}

func InitializeMockContact() (err error) {
	Cont1 = model.Contact{
		Model:     gorm.Model{ID: 1},
		Facebook:  faker.URL(),
		Instagram: faker.URL(),
		Linkedin:  faker.URL(),
		Twitter:   faker.URL(),
	}

	Cont2 = model.Contact{
		Model:     gorm.Model{ID: 2},
		Facebook:  faker.URL(),
		Instagram: faker.URL(),
		Linkedin:  faker.URL(),
		Twitter:   faker.URL(),
	}

	Cont3 = model.Contact{
		Model:     gorm.Model{ID: 3},
		Facebook:  faker.URL(),
		Instagram: faker.URL(),
		Linkedin:  faker.URL(),
		Twitter:   faker.URL(),
	}

	Cont4 = model.Contact{
		Model:     gorm.Model{ID: 4},
		Facebook:  faker.URL(),
		Instagram: faker.URL(),
		Linkedin:  faker.URL(),
		Twitter:   faker.URL(),
	}

	CreateContactReqMock = proto.CreateContactRequest{
		Contact: &proto.Contact{
			Facebook:  faker.URL(),
			Instagram: faker.URL(),
			Linkedin:  faker.URL(),
			Twitter:   faker.URL(),
		},
	}
	if err != nil {
		panic("Error occur while mocking data")
	}

	UpdateContactReqMock = proto.UpdateContactRequest{
		Contact: &proto.Contact{
			Id:        uint32(Cont1.ID),
			Facebook:  faker.URL(),
			Instagram: faker.URL(),
			Linkedin:  faker.URL(),
			Twitter:   faker.URL(),
		},
	}
	if err != nil {
		panic("Error occur while mocking data")
	}

	Conts = append(Conts, &Cont1, &Cont2, &Cont3, &Cont4)

	return nil
}
