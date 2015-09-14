package auction

type Result interface {
	HasResponse() bool
	GetResponse() *BidResponse
}