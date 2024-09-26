package app

import "github.com/purisaurabh/database-connection/internal/repository"

type Service interface{}

type ServiceStruct struct {
	Repo repository.Repository
}

func NewService(repo repository.Repository) *ServiceStruct {
	return &ServiceStruct{
		Repo: repo,
	}
}
