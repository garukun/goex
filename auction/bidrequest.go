package auction
import (
	"time"
	"github.com/garukun/goex/auction/types"
)

type BidRequest struct {
	Id string `json:"id"` // TODO(stevej): Embed Identifiable struct?
	Impressions []Impression `json:"imp"`
	// No site.
	Application *Application `json:"app"`
	Device *Device `json:"device"`
	User *User `json:"user"`
	AuctionType AuctionType `json:"at"`
	// No test.
	Timeout time.Duration `json:"tmax"`
	// No wseat.
	// No allimps.
	Currency Currency `json:"cur"`
	// No bcat.
	// No badv.
	// No regs.
}

type Impression struct {
	Id string `json:"id"`
	// No banner.
	Video *Video `json:"video"`
	// No native.
	DisplayManager string `json:"displaymanager"`
	DisplayManagerVersion string `json:"displaymangerver"`
	IsInterstitial types.NumericBool `json:"instl"`
	// No tagid.
	BidFloorPrice float64 `json:"bidfloor"`
	BidFloorCurrency Currency `json:"bidfloorcur"`
	// No secure.
	// No iframebuster.
	// No pmp.
}

type Application struct {

}

type Device struct {

}

type User struct {

}

type AuctionType int

const (
	FIRST_PRICE AuctionType = 1
	SECOND_PRICE AuctionType = 2
)

type Currency string

const (
	USD Currency = "USD"
	CNY Currency = "CNY"
	EUR Currency = "EUR"
)

type Video struct {

}