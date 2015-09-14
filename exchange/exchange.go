package exchange
import (
	"net/http"
//	"encoding/json"
)

type Exchange interface {
	HandleRequests(writer http.ResponseWriter, request *http.Request)
}


//func test(writer http.ResponseWriter, request *http.Request) {
//	a := auction.SimpleAuction{}
//	result := <- a.Run(bidRequest)
//
//	bidResponse := result.GetBidResponse()
//
//	js, err = json.Marshal(bidResponse)
//	if err != nil {
//		writer.WriteHeader(http.StatusBadRequest)
//		panic("Cannot create JSON payload from bid response.")
//	}
//	writer.Header().Set("Content-Type", "application/json")
//	writer.Write(js)
//}