package domain

import (
	"context"
	"gorm.io/gorm"
)

const NetPackageTableName = "netPackage"

type NetPackage struct {
	gorm.Model
	Duration      int  `json:"duration"` // as days
	Amount        int  `json:"amount"`   // By Byte
	Price         int  `json:"price"`    // Toman
	ProviderRefer uint `json:"provider_id"`
}

type NetPackageUseCase interface {
	CreateNetPackage(ctx context.Context, netPackage *NetPackage) (*NetPackage, error)
	GetNetPackage(ctx context.Context, npID int) (*NetPackage, error)
}

type NetPackageRepository interface {
	CreateNetPackage(ctx context.Context, netPackage *NetPackage) (*NetPackage, error)
	GetNetPackage(ctx context.Context, npID int) (*NetPackage, error)
}
