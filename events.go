package raychip

import (
	"fmt"
	"reflect"
)

type Topic struct {
	name string
	typ  reflect.Type
}

type EventBus struct {
	subscriptions map[Topic][]func(any)
}

func NewEventBus() EventBus {
	return EventBus{
		subscriptions: make(map[Topic][]func(any)),
	}
}

func (bus *EventBus) CreateSubscription(topicName string, msgType any, callback any) {
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

	bus.subscriptions[topic] = append(bus.subscriptions[topic], wrappedCallback)
}

func (bus *EventBus) Publish(topicName string, msg any) {
	msgType := reflect.TypeOf(msg)

	if msgType == nil {
		panic("msg cannot be nil")
	}

	for topic, cbks := range bus.subscriptions {
		if topicName == topic.name && topic.typ == msgType {
			for _, cbk := range cbks {
				cbk(msg)
			}
			return
		}
	}
	fmt.Printf("Topic %s not found for type %s\n", topicName, msgType)
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

type Publisher struct {
	bus   *EventBus
	topic Topic
}

func (p *Publisher) Publish(msg any) {
	p.bus.Publish(p.topic.name, msg)
}
