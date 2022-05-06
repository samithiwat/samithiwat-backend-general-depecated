package service

import (
	"github.com/samithiwat/samithiwat-backend-general/src/model"
	"github.com/samithiwat/samithiwat-backend-general/src/proto"
	"gorm.io/gorm"
)

func RawToDtoContact(cont *model.Contact) *proto.Contact {
	return &proto.Contact{
		Id:        uint32(cont.ID),
		Facebook:  cont.Facebook,
		Instagram: cont.Instagram,
		Twitter:   cont.Twitter,
		Linkedin:  cont.Linkedin,
	}
}

func DtoToRawContact(cont *proto.Contact) *model.Contact {
	return &model.Contact{
		Model:     gorm.Model{ID: uint(cont.Id)},
		Facebook:  cont.Facebook,
		Instagram: cont.Instagram,
		Twitter:   cont.Twitter,
		Linkedin:  cont.Linkedin,
	}
}
