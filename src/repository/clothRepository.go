package repository

import (
	"context"
	"server/src/config"
	"server/src/model"
)

type ClothRepository interface {
	ListCloth(ctx context.Context) ([]model.Cloth, error)
	GetCloth(ctx context.Context, id int) (*model.Cloth, error)
	CreateCloth(ctx context.Context, cloth *model.Cloth) (*model.Cloth, error)
	UpdateCloth(ctx context.Context, id int, updatedCloth *model.Cloth) (*model.Cloth, error)
	DeleteCloth(ctx context.Context, id int) error
}

type ClothRepositoryImpl struct{}

func (r *ClothRepositoryImpl) ListCloth(ctx context.Context) ([]model.Cloth, error) {
    var cloths []model.Cloth
    result := config.DB.WithContext(ctx).Find(&cloths)
    return cloths, result.Error
}

func (r *ClothRepositoryImpl) GetCloth(ctx context.Context, id int) (*model.Cloth, error) {
    var cloth model.Cloth
    result := config.DB.WithContext(ctx).First(&cloth, id)
    if result.Error != nil {
        return nil, result.Error
    }
    return &cloth, nil
}

func (r *ClothRepositoryImpl) CreateCloth(ctx context.Context, cloth *model.Cloth) (*model.Cloth, error) {
    result := config.DB.WithContext(ctx).Create(cloth)
    return cloth, result.Error
}

func (r *ClothRepositoryImpl) UpdateCloth(ctx context.Context, id int, updatedCloth *model.Cloth) (*model.Cloth, error) {
    var cloth model.Cloth
    result := config.DB.WithContext(ctx).First(&cloth, id)
    if result.Error != nil {
        return nil, result.Error
    }

    result = config.DB.WithContext(ctx).Model(&cloth).Updates(updatedCloth)
    if result.Error != nil {
        return nil, result.Error
    }
    return &cloth, nil
}

func (r *ClothRepositoryImpl) DeleteCloth(ctx context.Context, id int) error {
    result := config.DB.WithContext(ctx).Delete(&model.Cloth{}, id)
    return result.Error
}