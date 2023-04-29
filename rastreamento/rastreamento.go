package rastreamento

import (
	"fmt"
	"logistica/gerenciamento"
)

type Publisher struct {
	subscriberList []Subscriber
	code           int
	status         string
	shipping       gerenciamento.Transporte
}

func NewPublisher(code int, status string, shipping gerenciamento.Transporte) *Publisher {
	return &Publisher{
		code:     code,
		status:   status,
		shipping: shipping,
	}
}

func (p *Publisher) UpdateEvent() {
	fmt.Printf("Item de ID -> %v, está sendo enviado por: %v, status: %v\n", p.code, p.shipping.TranportMethodName(), p.status)
	p.status = "a caminho"
	p.NotifyAll()
}

func (p *Publisher) NotifyAll() {
	for _, subscriber := range p.subscriberList {
		subscriber.Update(string(p.code))
	}
}

func (p *Publisher) Register(s Subscriber) {
	p.subscriberList = append(p.subscriberList, s)
}

func (p *Publisher) Deregister(s Subscriber) {
	p.subscriberList = removeFromSlice(p.subscriberList, s)
}

func removeFromSlice(subscriberList []Subscriber, subscriberToRemove Subscriber) []Subscriber {
	subscriberListLength := len(subscriberList)
	for i, subscriber := range subscriberList {
		if subscriberToRemove.GetID() == subscriber.GetID() {
			subscriberList[subscriberListLength-1], subscriberList[i] = subscriberList[i], subscriberList[subscriberListLength-1]
			return subscriberList[:subscriberListLength-1]
		}
	}
	return subscriberList
}
