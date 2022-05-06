package repository

import (
	"github.com/samithiwat/samithiwat-backend-general/src/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ContactRepository struct {
	db gorm.DB
}

func (r *ContactRepository) FindOne(id int, perm *model.Contact) error {
	return r.db.Preload(clause.Associations).First(&perm, id).Error
}

func (r *ContactRepository) Create(perm *model.Contact) error {
	return r.db.Create(&perm).Error
}

func (r *ContactRepository) Update(id int, perm *model.Contact) error {
	return r.db.Where(id).Updates(&perm).First(&perm).Error
}

func (r *ContactRepository) Delete(id int, perm *model.Contact) error {
	return r.db.First(&perm, id).Delete(&model.Contact{}).Error
}
