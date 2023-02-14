package repository

import (
	"context"
	"gorm.io/gorm"
	"internet_service/internal/domain"
)

type ProviderRepository struct {
	*gorm.DB
}

func (pr *ProviderRepository) CreateProvider(ctx context.Context, provider *domain.Provider) (*domain.Provider, error) {
	result := pr.WithContext(ctx).Create(provider)
	if result.Error != nil {
		return nil, result.Error
	}
	return provider, nil
}
func (pr *ProviderRepository) GetProvider(ctx context.Context, providerID int) (*domain.Provider, error) {
	provider := &domain.Provider{}
	result := pr.WithContext(ctx).First(provider, providerID)
	if result.Error != nil {
		return nil, result.Error
	}
	return provider, nil
}

func (pr *ProviderRepository) DeleteProvider(ctx context.Context, providerID uint) error {
	result := pr.Delete(&domain.Provider{
		Model: gorm.Model{
			ID: providerID,
		},
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
