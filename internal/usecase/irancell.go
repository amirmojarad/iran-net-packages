package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/gocolly/colly"
	"internet_service/config"
	"internet_service/internal/domain"
	"internet_service/internal/repository"
	"internet_service/internal/utils"
	"strings"
)

type IrancellUseCase struct {
	*domain.Provider
	*repository.ProviderRepository
	*config.Config
}

func fetch() {
	conf := config.GetConfToml()
	c := colly.NewCollector()
	c.OnError(func(response *colly.Response, err error) {
		fmt.Println(err)
	})
	result := make([]*domain.NetPackage, 0)
	c.OnHTML("div", func(element *colly.HTMLElement) {
		if element.Attr("class") == "product-box" {
			res := strings.Split(element.Text, "\n")
			amount, err := utils.ParseAmount(res[4])
			if err != nil {
				// TODO fix these errors
				panic(err)
			}
			duration, err := utils.ParseDuration(res[5])
			if err != nil {
				// TODO fix these errors
				panic(err)
			}
			price, err := utils.ParsePrice(res[14])
			if err != nil {
				// TODO fix these errors
				panic(err)
			}
			netPackage := &domain.NetPackage{
				Duration: duration,
				Amount:   amount,
				Price:    price,
			}
			result = append(result, netPackage)
		}
	})
	c.OnScraped(func(response *colly.Response) {
		fmt.Printf("on Scrapped %+v\n", result)
	})
	err := c.Visit(conf.Providers.Irancell)
	if err != nil {
		return
	}
}

func (irancell *IrancellUseCase) CreateProvider(ctx context.Context, provider *domain.Provider) (*domain.Provider, error) {
	if provider.Title == "" {
		return nil, errors.New("title is empty")
	}
	if provider.Link == "" {
		return nil, errors.New("link is empty")
	}
	return irancell.ProviderRepository.CreateProvider(ctx, provider)
}

func (irancell *IrancellUseCase) GetProvider(ctx context.Context, providerID int) (*domain.Provider, error) {
	return nil, nil

}
func (irancell *IrancellUseCase) Fetch(ctx context.Context) error {
	c := colly.NewCollector()
	c.OnError(func(response *colly.Response, err error) {
		fmt.Println(err)
	})
	result := make([]*domain.NetPackage, 0)
	c.OnHTML("div", func(element *colly.HTMLElement) {
		if element.Attr("class") == "product-box" {
			res := strings.Split(element.Text, "\n")
			amount, err := utils.ParseAmount(res[4])
			if err != nil {
				// TODO fix these errors
				panic(err)
			}
			duration, err := utils.ParseDuration(res[5])
			if err != nil {
				// TODO fix these errors
				panic(err)
			}
			price, err := utils.ParsePrice(res[14])
			if err != nil {
				// TODO fix these errors
				panic(err)
			}
			netPackage := &domain.NetPackage{
				Duration: duration,
				Amount:   amount,
				Price:    price,
			}
			result = append(result, netPackage)
		}
	})
	c.OnScraped(func(response *colly.Response) {
		fmt.Printf("on Scrapped %+v\n", result)
		provider := &domain.Provider{
			Link:     irancell.Config.Providers.Irancell,
			Title:    "Irancell",
			Packages: result,
		}
		_, err := irancell.CreateProvider(ctx, provider)
		if err != nil {
			return
		}
	})
	err := c.Visit(irancell.Config.Providers.Irancell)
	if err != nil {
		return err
	}
	return nil
}
func (irancell *IrancellUseCase) GetPackages(ctx context.Context) ([]*domain.NetPackage, error) {
	return nil, nil

}
