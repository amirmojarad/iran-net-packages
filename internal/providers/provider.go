//package providers
//
//import (
//	"github.com/gocolly/colly/v2"
//	"log"
//)
//
//const (
//	IRANCELL_URL = "https://irancell.ir/o/1001/%D8%A7%DB%8C%D8%B1%D8%A7%D9%86%D8%B3%D9%84-%D8%A8%D8%B3%D8%AA%D9%87-%D9%87%D8%A7%DB%8C-%D8%A7%DB%8C%D9%86%D8%AA%D8%B1%D9%86%D8%AA-%D9%87%D9%85%D8%B1%D8%A7%D9%87"
//)
//
//type Providers struct {
//	Irancells   []*NetPackage
//	Rightels    []*NetPackage
//	HamrahAvals []*NetPackage
//}
//
//func NewProviders() *Providers {
//	p := &Providers{}
//	p.Rightels = make([]*NetPackage, 0)
//	p.Irancells = make([]*NetPackage, 0)
//	p.HamrahAvals = make([]*NetPackage, 0)
//	return p
//}
//
//type NetPackage struct {
//	Duration int `json:"duration"` // as days
//	Amount   int `json:"amount"`   // By Byte
//	Price    int `json:"price"`    // Toman
//}
//
//type Provider interface {
//	GetIrancell()
//	GetRightel()
//	GetHamrahAval()
//}
//
//func (np *Providers) GetIrancell() {
//	//
//	//c := colly.NewCollector()
//	//c.OnError(onError)
//	//c.OnHTML("div", func(element *colly.HTMLElement) {
//	//	if element.Attr("class") == "product-box" {
//	//		res := strings.Split(element.Text, "\n")
//	//		amountAsBytes, _ :=utils.ParseAmount(res[4])
//	//		netPackage := NetPackage{
//	//			Duration: ,
//	//			Amount:   0,
//	//			Price:    0,
//	//		}
//	//		np.Irancells = append(np.Irancells, res[4])
//	//		result = append(result, res[5])
//	//		result = append(result, res[14])
//	//
//	//		fmt.Printf("%+v\n", res)
//	//
//	//	}
//	//})
//	//if err := c.Visit(IRANCELL_URL); err != nil {
//	//	log.Println(err)
//	//}
//
//}
//
//func (np *Providers) GetRightel() {
//
//}
//
//func (np *Providers) GetHamrahAval() {
//
//}
//
//// Callback functions
//func onError(response *colly.Response, err error) {
//	log.Println(err)
//}

package providers

type Provider interface {
	Fetch()
	GetOffers() []*NetPackage
}

type NetPackage struct {
	Duration int `json:"duration"` // as days
	Amount   int `json:"amount"`   // By Byte
	Price    int `json:"price"`    // Toman
}
