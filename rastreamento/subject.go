package rastreamento

type Subject interface {
	register(subscriber Subscriber)
	deregister(subscriber Subscriber)
	notifyAll()
}
