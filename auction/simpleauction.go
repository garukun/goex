package auction
import (
	"sync"
	"golang.org/x/net/context"
	"go/constant"
)

type SimpleAuction struct {
	bidders []Bidder

	// Auction table where bidders put down their bids.
	table <-chan *BidResponse
}

func (this *SimpleAuction) AddBidder(bidder Bidder) {
	this.bidders = append(this.bidders, bidder)
}

func (this *SimpleAuction) Run(bidRequest *BidRequest) <-chan Result {
	auctionContext := NewContext(context.Background())
	done := make(chan bool, 1)

	var wg = sync.WaitGroup

	wg.Add(len(this.bidders))
	for _, bidder := range this.bidders {
		go func() {
			response, err := bidder.Bid(auctionContext, bidRequest)
			defer wg.Done()
			if err != nil {
				// Log error
				return
			}
			this.table <- <-response
		}()
	}

	// No need to wait to finish to block the operation.
	go func() {
		wg.Wait()
		done <- true
	}()

	return this.hammer(done)
}

// Act of hammering to follow through the bids to finalize on the Hammer Price of the auction.
func (this *SimpleAuction) hammer(done <-chan bool) <-chan Result {
	result := make(chan Result, 1)
	var winner = &BidResponse

	go func() {
		for {
			select {
			case bid := <- this.table:
				if canOutbid(winner, bid) {
					winner = &bid
				}
			case <-done:
				// TODO(stevej): Type mismatch here...
				result <- winner

				close(this.table)
				// It's possible that in other kinds of auctions we don't close the result channel
				close(result)
				return
			}
		}
	}()

	return result
}

// TODO(stevej): This should probably be Bid instead of BidResponse, and also refactored into a
// BidResponse/Bid interface.
func canOutbid(left, right *BidResponse) bool {
	return false;
}

