package auction
import "golang.org/x/net/context"

type Bidder interface {
	Bid(auctionContext context.Context, bidRequest *BidRequest) (<-chan *BidResponse, error)
}