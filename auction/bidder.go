package auction
import (
"golang.org/x/net/context"
"github.com/garukun/goex/auction/types"
)

type Bidder interface {
	Bid(auctionContext context.Context, bidRequest *types.BidRequest) (<-chan *types.BidResponse, error)
}