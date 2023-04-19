package rastreamento

type Subscriber interface {
	Update(string)
	GetID() string
}
