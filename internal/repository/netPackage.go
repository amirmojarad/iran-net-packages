package repository

import (
	"context"
	"gorm.io/gorm"
	"internet_service/internal/domain"
)

type NetPackageRepository struct {
	*gorm.DB
}

func (npRepository *NetPackageRepository) CreateNetPackage(ctx context.Context, netPackage *domain.NetPackage) (*domain.NetPackage, error) {
	result := npRepository.WithContext(ctx).Create(netPackage)
	if result.Error != nil {
		return nil, result.Error
	}
	return netPackage, nil
}

func (npRepository *NetPackageRepository) GetNetPackage(ctx context.Context, npID int) (*domain.NetPackage, error) {
	netPackage := &domain.NetPackage{}
	result := npRepository.WithContext(ctx).First(&netPackage, npID)
	if result.Error != nil {
		return nil, result.Error
	}
	return netPackage, nil
}

func (npRepository *NetPackageRepository) DeleteNetPackage(ctx context.Context, npID uint) error {
	result := npRepository.Delete(&domain.NetPackage{
		Model: gorm.Model{
			ID: npID,
		},
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
