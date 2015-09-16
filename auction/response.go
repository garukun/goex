package auction
import "io"

type Response interface {
	Price()
	Write(writer io.Writer)
}