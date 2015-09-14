package exchange
import (
	"net/http"
	"encoding/json"
	"github.com/garukun/goex/auction"
	"io"
)

type SimpleExchange struct {

}

func (this *SimpleExchange) HandleRequests() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		bidRequest := decodeBidRequest(request.Body)

		auc := auction.Auction()
		result := <-auc.Run(bidRequest)

		if !result.HasResponse() {
			writer.WriteHeader(http.StatusNoContent)
		}

		writer.Header().Set("Content-Type", "application/json")
		writeBidResponse(writer, result.GetResponse())
	}
}

func decodeBidRequest(requestBody io.ReadCloser) *auction.BidRequest {
	decoder := json.NewDecoder(requestBody)
	var bidRequest auction.BidRequest
	decoder.Decode(&bidRequest)
	requestBody.Close()
	return &bidRequest
}

func writeBidResponse(writer io.Writer, bidResponse *auction.BidResponse) {
	encoder := json.NewEncoder(writer)
	encoder.Encode(bidResponse)
}
