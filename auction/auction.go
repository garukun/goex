package auction

type Auction interface {
	Invite(bidder Bidder)
	InviteAll(bidders []Bidder)
	Run(bidRequest *BidRequest) <-chan Result
}
