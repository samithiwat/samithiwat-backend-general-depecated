package test

import (
	"github.com/samithiwat/samithiwat-backend-general/src/model"
	"github.com/samithiwat/samithiwat-backend-general/src/proto"
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
