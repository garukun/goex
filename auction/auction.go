package auction
import "github.com/garukun/goex/auction/types"

type Auction interface {
	Invite(bidder Bidder)
	InviteAll(bidders []Bidder)
	Run(bidRequest *types.BidRequest) <-chan Result
}


func NewAuction() Auction {
	return &SimpleAuction{}
}