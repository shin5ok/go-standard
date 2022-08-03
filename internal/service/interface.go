package service

import "context"

type Service interface {
	GetInfo(context.Context) params
	PutInfo(context.Context, params) error
}
