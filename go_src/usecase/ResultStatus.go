package usecase

// ResultStatusの構造体
type ResultStatus struct {
	StatusCode int
	Error error
}


func NewResultStatus(code int, er error) *ResultStatus {
	result := &ResultStatus{
		StatusCode: code,
		Error: er,
	}

	return result
}
