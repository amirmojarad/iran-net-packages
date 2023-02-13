package usecase

import (
	"context"
	"errors"
	"internet_service/config"
	"internet_service/internal/domain"
	"internet_service/internal/repository"
)

type NetPackageUseCase struct {
	*config.Config
	*domain.NetPackage
	*repository.NetPackageRepository
}

func (np *NetPackageUseCase) CreateNetPackage(ctx context.Context, netPackage *domain.NetPackage) (*domain.NetPackage, error) {
	if netPackage.Price <= 0 {
		return nil, errors.New("price is less than one")
	}
	if netPackage.Amount <= 0 {
		return nil, errors.New("amount of package is less than zero, its impossible")
	}
	if netPackage.Duration <= 0 {
		return nil, errors.New("duration of package is less than zero")
	}
	return np.NetPackageRepository.CreateNetPackage(ctx, netPackage)
}
func (np *NetPackageUseCase) GetNetPackage(ctx context.Context, npID int) (*domain.NetPackage, error) {
	return np.NetPackageRepository.GetNetPackage(ctx, npID)
}
