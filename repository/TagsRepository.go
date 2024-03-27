package repository

import (
	"gorestgorm/models"
)

type TagsRepository interface {
	Save(tags models.Tags)
	Update(tags models.Tags)
	Delete(tagsId int)
	FindById(tagsId int) (tags models.Tags, err error)
	FindAll() []models.Tags
}