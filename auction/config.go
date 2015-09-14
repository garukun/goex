package auction
import "time"

type Config struct {
	Timeout time.Duration
	AuctionType AuctionType
}