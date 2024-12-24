package model

type TodolistRequest struct {
	Id   int64  `validate:"required"`
	Task string `validate:"required"`
}

type TodolistResponse struct {
	Id   int64  `json:"id"`
	Task string `json:"task"`
}

type GlobalResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
	Errors  any    `json:"errors"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Error   any    `json:"errors"`
}
