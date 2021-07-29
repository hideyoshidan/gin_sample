package controllers

// 
type H struct {
    Message string `json:"message"`
    Data interface{} `json:"data"`
}

// Hの構造体を初期化してポインタを返す
func NewH(message string, data interface{}) *H {
    H := new(H)
    H.Message = message
    H.Data = data
    return H
}