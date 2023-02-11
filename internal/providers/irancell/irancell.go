package irancell

import (
	"fmt"
	"github.com/gocolly/colly"
	"internet_service/config"
	"internet_service/internal/providers"
	"internet_service/internal/utils"
	"strings"
)

type Irancell struct {
	Link     string
	Packages []*providers.NetPackage
}

func New(link string) *Irancell {
	return &Irancell{
		Link:     link,
		Packages: make([]*providers.NetPackage, 0),
	}
}

func (instance *Irancell) GetOffers() []*providers.NetPackage {
	return nil
}

func (instance *Irancell) Fetch() {
	conf := config.GetConfToml()
	c := colly.NewCollector()
	c.OnError(func(response *colly.Response, err error) {
		fmt.Println(err)
	})
	result := make([]string, 0)
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
			netPackage := &providers.NetPackage{
				Duration: duration,
				Amount:   amount,
				Price:    price,
			}
			instance.Packages = append(instance.Packages, netPackage)
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
