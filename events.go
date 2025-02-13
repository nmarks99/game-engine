package raychip

import (
	"fmt"
	"reflect"
)

type Topic struct {
	name string
	typ  reflect.Type
}

type Subscription struct {
	function func(any)
	id       int
}

type EventBus struct {
	subscriptions map[Topic][]Subscription
}

func NewEventBus() EventBus {
	return EventBus{
		subscriptions: make(map[Topic][]Subscription),
	}
}

func (bus *EventBus) CreateSubscription(topicName string, msgType any, callback any) int {
	typ := reflect.TypeOf(msgType)

	if typ == nil || typ.Kind() == reflect.Pointer {
		panic("Subscription message type cannot be nil or pointer")
	}

	callbackValue := reflect.ValueOf(callback)
	callbackType := callbackValue.Type()
	if callbackType.Kind() != reflect.Func {
		panic("callback must be a function")
	}
	if callbackType.NumIn() != 1 || callbackType.In(0) != typ {
		panic(fmt.Sprintf("callback must accept one argument of type %v", typ))
	}

	topic := Topic{name: topicName, typ: typ}

	// wrap the callback in a type assertion
	wrappedCallback := func(msg any) {
		callbackValue.Call([]reflect.Value{reflect.ValueOf(msg)})
	}

	id := len(bus.subscriptions[topic])
	sub := Subscription{
		function: wrappedCallback,
		id:       len(bus.subscriptions[topic]),
	}
	bus.subscriptions[topic] = append(bus.subscriptions[topic], sub)

	return id
}

func (bus *EventBus) RemoveSubscription(topicName string, id int) {
    for topic, subs := range bus.subscriptions {
        if topic.name == topicName {
            if len(subs) > id {
                newSubs := append(subs[:id], subs[id+1:]...)
                bus.subscriptions[topic] = newSubs
                break
            }
        }
    }
}

func (bus *EventBus) ClearSubscriptions(topicName string) {
    for topic, subs := range bus.subscriptions {
        if topic.name == topicName {
            bus.subscriptions[topic] = subs[:0]
        }
    }
}

func (bus *EventBus) Publish(topicName string, msg any) {
	msgType := reflect.TypeOf(msg)

	if msgType == nil {
		panic("msg cannot be nil")
	}

	for topic, subs := range bus.subscriptions {
		if topicName == topic.name && topic.typ == msgType {
			for _, cbk := range subs {
				cbk.function(msg)
			}
			return
		}
	}
	fmt.Printf("Topic %s not found for type %s\n", topicName, msgType)
}

type Publisher struct {
	bus   *EventBus
	topic Topic
}

func (p *Publisher) Publish(msg any) {
	p.bus.Publish(p.topic.name, msg)
}

func (bus *EventBus) CreatePublisher(topicName string, msgType any) *Publisher {
	typ := reflect.TypeOf(msgType)
	if typ == nil {
		panic("msgType cannot be nil")
	}
	if typ.Kind() == reflect.Pointer {
		panic("msgType cannot be a pointer")
	}
	return &Publisher{
		bus:   bus,
		topic: Topic{name: topicName, typ: typ},
	}
}
