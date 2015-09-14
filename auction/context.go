package auction
import "golang.org/x/net/context"

type wrapper struct {
	context.Context
}

// TODO(stevej): Figure out what goes into an auction context
func NewContext(parent context.Context) context.Context {
	return &wrapper{parent}
}
