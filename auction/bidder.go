package auction

import (
	"github.com/garukun/goex/auction/types"
	"golang.org/x/net/context"
)

type Bidder interface {
	Bid(auctionContext context.Context, bidRequest *types.BidRequest) (<-chan *types.BidResponse, error)
}
