package database

import (
    "errors"
    "github.com/jinzhu/gorm"
    "github.com/hideyoshidan/go_bbs/domain"
    "fmt"
)

type  SampleRepository struct {}

func (repo *SampleRepository) FindByID(db *gorm.DB, id int) (sample domain.Samples, err error) {
    sample = domain.Samples{}
    db.First(&sample, id)
    if sample.ID <= 0 {
        fmt.Println(sample.ID)
        return domain.Samples{}, errors.New("user is not found")
    }
    return sample, nil
}