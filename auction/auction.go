package auction

type Auction interface {
	Run(bidRequest *BidRequest) <-chan Result
}
