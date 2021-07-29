package usecase

import (
    "github.com/hideyoshidan/go_bbs/domain"
)

type SampleInteractor struct {
    DB DBRepository
    Sample SampleRepository
}

func (interactor *SampleInteractor) Get(id int) (sample domain.SamplesForGet, resultStatus *ResultStatus) {
    db := interactor.DB.Connect()
    // Users の取得
    foundSample, err := interactor.Sample.FindByID(db, id)
    if err != nil {
        return domain.SamplesForGet{}, NewResultStatus(404, err)
    }
    sample = foundSample.BuildForGet()
    return sample, NewResultStatus(200, nil)
}