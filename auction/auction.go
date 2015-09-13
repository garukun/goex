package auction

type Auction interface {
	Run() <-chan Result
}
