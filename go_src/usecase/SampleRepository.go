package usecase

import (
    "github.com/jinzhu/gorm"
    "github.com/hideyoshidan/go_bbs/domain"
)

type SampleRepository interface {
    FindByID(db *gorm.DB, id int) (sample domain.Samples, err error)
}