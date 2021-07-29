package controllers

import (
	"strconv"
	"github.com/hideyoshidan/go_bbs/interfaces/database"
	"github.com/hideyoshidan/go_bbs/usecase"
)

// SampleControllerの構造体
type SampleController struct {
	Interactor usecase.SampleInteractor
}

// SampleControllerを初期化してポインタを返す
func NewSampleController(db database.DB) *SampleController {
	return &SampleController {
		Interactor: usecase.SampleInteractor{
			DB: &database.DBRepository{ DB: db },
            Sample: &database.SampleRepository{},
		},
	}
}

// ControllerのAction
func (controller *SampleController) Get(c Context) {

    id, _ := strconv.Atoi(c.Param("id"))

    sample, res := controller.Interactor.Get(id)
    if res.Error != nil {
        c.JSON(res.StatusCode, NewH(res.Error.Error(), nil))
        return
    }
    c.JSON(res.StatusCode, NewH("success", sample))
}