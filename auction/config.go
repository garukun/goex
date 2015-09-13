package auction
import "time"

type Config struct {
	Timeout time.Duration
	AuctionType AuctionType
}

type AuctionType int

const (
	FIRST_PRICE AuctionType = 1
	SECOND_PRICE AuctionType = 2
)
