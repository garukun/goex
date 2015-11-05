package auction

import (
	"flag"
	"github.com/garukun/goex/auction/types"
	"time"
)

const DEFAULT_AUCTION_TIMEOUT = 50 * time.Millisecond

var timeout *time.Duration
var auctionType *types.AuctionType

func init() {
	timeout = flag.Duration("auction_timeout", DEFAULT_AUCTION_TIMEOUT, "Duration for how long the auction should run, eg. 200ms.")
	auctionType = flag.Uint("auction_type", 2, "Auction type, eg. first or second priced auction.")

	validateFlags()
}

func validateFlags() {
	if auctionType != uint(types.FIRST_PRICE) || auctionType != uint(types.SECOND_PRICE) {
		panic("Usage: auction_type indicate either first priced auction (1) or second priced (2).")
	}
}

func Tiemout() time.Duration {
	return *timeout
}

func Type() types.AuctionType {
	return types.AuctionType(*auctionType)
}
