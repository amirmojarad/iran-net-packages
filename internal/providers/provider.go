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
