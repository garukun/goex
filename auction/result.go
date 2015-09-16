package auction
import "io"

type Result interface {
	IsCompleted() bool
	Price() float64
	Write(writer io.Writer)
}