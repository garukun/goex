package auction

type Result interface {
	IsCompleted() bool
	GetResponse() Response
}