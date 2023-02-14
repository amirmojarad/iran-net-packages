package domain

import (
	"context"
	"gorm.io/gorm"
)

type Provider struct {
	gorm.Model
	Link     string        `json:"link"`
	Title    string        `json:"title"`
	Packages []*NetPackage `json:"packages" gorm:"foreignKey:ProviderRefer"`
}

type ProviderUseCase interface {
	CreateProvider(ctx context.Context, provider *Provider) (*Provider, error)
	GetProvider(ctx context.Context, providerID int) (*Provider, error)
	Fetch(ctx context.Context) error
	GetPackages(ctx context.Context) ([]*NetPackage, error)
}

type ProviderRepository interface {
	CreateProvider(ctx context.Context, provider *Provider) (*Provider, error)
	GetProvider(ctx context.Context, providerID int) (*Provider, error)
	DeleteProvider(ctx context.Context, providerID uint) error
}
