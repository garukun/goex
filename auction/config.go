package auction
import (
	"time"
	"github.com/garukun/goex/auction/types"
)

type Config struct {
	Timeout time.Duration
	AuctionType types.AuctionType
}