package auction

import (
	"io"
	"net/http"
)

type Result interface {
	IsCompleted() bool
	Price() float64
	Write(writer io.Writer)
	WriteResponse(responseWriter http.ResponseWriter)
}
