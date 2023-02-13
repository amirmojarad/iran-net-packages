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
	repository := NetPackageRepository{
		handle,
	}
	np := &domain.NetPackage{
		Duration:      15,
		Amount:        15000,
		Price:         15000,
		ProviderRefer: 0,
	}
	res, err := repository.CreateNetPackage(context.Background(), np)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
