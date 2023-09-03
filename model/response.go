package model

type Response struct {
	IsSuccess bool   `json:"isSuccess"`
	Message   string `json:"message"`
}
