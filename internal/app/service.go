package app

import (
	"context"
	"fmt"
	"time"

	specs "github.com/purisaurabh/database-connection/internal/pkg"
	"github.com/purisaurabh/database-connection/internal/repository"
)

type Service interface {
	PostProfileData(context.Context, specs.PostProfile) error
	ListProfileData(context.Context) ([]specs.ListProfileResponse, error)
	UpdateProfileData(context.Context, int, specs.UpdateProfile) error
	DeleteProfileData(context.Context, int) error
}

type ServiceStruct struct {
	Repo repository.Repository
}

func NewService(repo repository.Repository) *ServiceStruct {
	return &ServiceStruct{
		Repo: repo,
	}
}

func (s *ServiceStruct) PostProfileData(ctx context.Context, req specs.PostProfile) error {
	now := time.Now().Unix()
	postProfileRepo := repository.PostProfileRepo{
		Name:       req.Profiles.Name,
		Email:      req.Profiles.Email,
		Mobile:     req.Profiles.Mobile,
		Created_At: now,
		Updated_At: now,
	}

	err := s.Repo.PostProfileData(ctx, postProfileRepo)
	if err != nil {
		fmt.Println("Error in calling repository:", err)
		return err
	}
	return nil
}

func (s *ServiceStruct) ListProfileData(ctx context.Context) ([]specs.ListProfileResponse, error) {
	resp, err := s.Repo.ListProfileData(ctx)
	if err != nil {
		fmt.Println("Error in calling repository:", err)
		return []specs.ListProfileResponse{}, err
	}

	var response []specs.ListProfileResponse
	for _, profile := range resp {
		response = append(response, specs.ListProfileResponse{
			ID:         profile.ID,
			Name:       profile.Name,
			Email:      profile.Email,
			Mobile:     profile.Mobile,
			Created_At: profile.Created_At,
			Updated_At: profile.Updated_At,
		})
	}

	return response, nil
}

func (s *ServiceStruct) UpdateProfileData(ctx context.Context, id int, req specs.UpdateProfile) error {
	now := time.Now().Unix()
	updateProfileData := repository.UpdateProfileData{
		ID:         id,
		Name:       req.Profiles.Name,
		Email:      req.Profiles.Email,
		Mobile:     req.Profiles.Mobile,
		Updated_At: now,
	}

	err := s.Repo.UpdateProfileData(ctx, updateProfileData)
	if err != nil {
		fmt.Println("Error in calling repository:", err)
		return err
	}
	return nil
}

func (s *ServiceStruct) DeleteProfileData(ctx context.Context, id int) error {
	err := s.Repo.DeleteProfileData(ctx, id)
	if err != nil {
		fmt.Println("Error in calling repository:", err)
		return err
	}
	return nil
}
