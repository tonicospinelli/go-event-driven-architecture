package inmemory

import (
	"context"
	"sync"

	"github.com/google/uuid"

	"eda-in-golang/internal/am"
)

type subscription struct {
	id      string
	topic   string
	handler am.RawMessageHandler
	stream  *stream
}

func (s *subscription) Unsubscribe() error {
	return s.stream.removeSubscription(s)
}

type stream struct {
	mu            sync.RWMutex
	subscriptions map[string]map[string]am.RawMessageHandler
}

var _ am.RawMessageStream = (*stream)(nil)

func NewStream() am.RawMessageStream {
	return &stream{
		subscriptions: make(map[string]map[string]am.RawMessageHandler),
	}
}

func (t *stream) Publish(ctx context.Context, topicName string, v am.RawMessage) error {
	t.mu.RLock()
	handlers, exists := t.subscriptions[topicName]
	if !exists {
		t.mu.RUnlock()
		return nil
	}

	// Copy handlers to avoid holding the lock during HandleMessage
	handlersCopy := make([]am.RawMessageHandler, 0, len(handlers))
	for _, handler := range handlers {
		handlersCopy = append(handlersCopy, handler)
	}
	t.mu.RUnlock()

	for _, handler := range handlersCopy {
		err := handler.HandleMessage(ctx, &rawMessage{v})
		if err != nil {
			return err
		}
	}
	return nil
}

func (t *stream) Subscribe(topicName string, handler am.RawMessageHandler, options ...am.SubscriberOption) (am.Subscription, error) {
	cfg := am.NewSubscriberConfig(options)

	var filters map[string]struct{}
	if len(cfg.MessageFilters()) > 0 {
		filters = make(map[string]struct{})
		for _, key := range cfg.MessageFilters() {
			filters[key] = struct{}{}
		}
	}

	fn := am.RawMessageHandlerFunc(func(ctx context.Context, msg am.IncomingRawMessage) error {
		if filters != nil {
			if _, exists := filters[msg.MessageName()]; !exists {
				return nil
			}
		}

		return handler.HandleMessage(ctx, msg.(*rawMessage))
	})

	t.mu.Lock()
	defer t.mu.Unlock()

	if _, exists := t.subscriptions[topicName]; !exists {
		t.subscriptions[topicName] = make(map[string]am.RawMessageHandler)
	}

	subID := uuid.New().String()
	t.subscriptions[topicName][subID] = fn

	return &subscription{
		id:      subID,
		topic:   topicName,
		handler: fn,
		stream:  t,
	}, nil
}

func (t *stream) Unsubscribe() error {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.subscriptions = make(map[string]map[string]am.RawMessageHandler)

	return nil
}

func (t *stream) removeSubscription(sub *subscription) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	if handlers, exists := t.subscriptions[sub.topic]; exists {
		delete(handlers, sub.id)
		if len(handlers) == 0 {
			delete(t.subscriptions, sub.topic)
		}
	}

	return nil
}
