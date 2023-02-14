package repository

import (
	"context"
	"github.com/stretchr/testify/assert"
	"internet_service/infrastructure/database"
	"internet_service/internal/domain"
	"testing"
)

func TestNetPackageRepository_CreateNetPackage(t *testing.T) {
	handle := database.Connect()
	npRepository := NetPackageRepository{
		handle,
	}
	providerRepository := ProviderRepository{
		handle,
	}
	ctx := context.Background()
	createdProvider, err := providerRepository.CreateProvider(ctx, &domain.Provider{
		Link:     "mocklink",
		Title:    "mci",
		Packages: nil,
	})
	assert.Nil(t, err)
	assert.NotNil(t, createdProvider)
	t.Log(createdProvider)
	createdNP, err := npRepository.CreateNetPackage(ctx, &domain.NetPackage{
		Duration:      15,
		Amount:        15000,
		Price:         15000,
		ProviderRefer: createdProvider.ID,
	})
	assert.Nil(t, err)
	assert.NotNil(t, createdNP)

	err = npRepository.DeleteNetPackage(ctx, createdNP.ID)
	assert.Nil(t, err)
	err = providerRepository.DeleteProvider(ctx, createdProvider.ID)
	assert.Nil(t, err)
}
