package service

import (
	"context"
	"server/src/model"
	"server/src/repository"
)

type ClothService interface {
    ListCloth(ctx context.Context) ([]model.Cloth, error)
    GetCloth(ctx context.Context, id int) (*model.Cloth, error)
    CreateCloth(ctx context.Context, cloth *model.Cloth) (*model.Cloth, error)
    UpdateCloth(ctx context.Context, id int, updatedCloth *model.Cloth) (*model.Cloth, error)
    DeleteCloth(ctx context.Context, id int) error
}

type ClothServiceImpl struct {
    Repo repository.ClothRepository
}

func (s *ClothServiceImpl) ListCloth(ctx context.Context) ([]model.Cloth, error) {
    return s.Repo.ListCloth(ctx)
}

func (s *ClothServiceImpl) GetCloth(ctx context.Context, id int) (*model.Cloth, error) {
    return s.Repo.GetCloth(ctx, id)
}

func (s *ClothServiceImpl) CreateCloth(ctx context.Context, cloth *model.Cloth) (*model.Cloth, error) {
    return s.Repo.CreateCloth(ctx, cloth)
}

func (s *ClothServiceImpl) UpdateCloth(ctx context.Context, id int, updatedCloth *model.Cloth) (*model.Cloth, error) {
    return s.Repo.UpdateCloth(ctx, id, updatedCloth)
}

func (s *ClothServiceImpl) DeleteCloth(ctx context.Context, id int) error {
    return s.Repo.DeleteCloth(ctx, id)
}