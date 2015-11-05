package auction

import "github.com/garukun/goex/auction/types"

type Auction interface {
	Invite(bidders []Bidder)
	Run(bidRequest *types.BidRequest) Result
}

func NewAuction() Auction {
	return &SimpleAuction{}
}
