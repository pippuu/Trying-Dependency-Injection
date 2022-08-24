package service

import (
	"Week4/repository"
)

type Service struct {
	R repository.Repository
}

func InitService(repo repository.Repository) Service {
	return Service{
		R: repo,
	}
}
