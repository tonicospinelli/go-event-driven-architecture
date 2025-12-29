package inmemory

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"eda-in-golang/internal/am"
)

type mockRawMessage struct {
	id   string
	name string
}

func (m mockRawMessage) ID() string          { return m.id }
func (m mockRawMessage) Subject() string     { return "" }
func (m mockRawMessage) MessageName() string { return m.name }
func (m mockRawMessage) Data() []byte        { return nil }

func TestStream_PublishSubscribe(t *testing.T) {
	s := NewStream()
	topic := "test-topic"
	handler := &am.MockRawMessageHandler[am.IncomingRawMessage]{}

	_, err := s.Subscribe(topic, handler)
	assert.NoError(t, err)

	ctx := context.Background()
	msg := mockRawMessage{id: "1", name: "test"}

	handler.On("HandleMessage", mock.Anything, mock.Anything).Return(nil)

	err = s.Publish(ctx, topic, msg)
	assert.NoError(t, err)

	handler.AssertExpectations(t)
}

func TestStream_MessageFiltering(t *testing.T) {
	s := NewStream()
	topic := "test-topic"
	handler := &am.MockRawMessageHandler[am.IncomingRawMessage]{}

	_, err := s.Subscribe(topic, handler, am.MessageFilter{"wanted"})
	assert.NoError(t, err)

	ctx := context.Background()

	// Should be filtered out
	msg1 := mockRawMessage{id: "1", name: "unwanted"}
	err = s.Publish(ctx, topic, msg1)
	assert.NoError(t, err)

	// Should be handled
	msg2 := mockRawMessage{id: "2", name: "wanted"}
	handler.On("HandleMessage", mock.Anything, mock.MatchedBy(func(m am.IncomingRawMessage) bool {
		return m.ID() == "2"
	})).Return(nil)

	err = s.Publish(ctx, topic, msg2)
	assert.NoError(t, err)

	handler.AssertExpectations(t)
	handler.AssertNumberOfCalls(t, "HandleMessage", 1)
}

func TestStream_MultipleTopics(t *testing.T) {
	s := NewStream()
	topic1 := "topic-1"
	topic2 := "topic-2"

	handler1 := &am.MockRawMessageHandler[am.IncomingRawMessage]{}
	handler2 := &am.MockRawMessageHandler[am.IncomingRawMessage]{}

	_, err := s.Subscribe(topic1, handler1)
	assert.NoError(t, err)
	_, err = s.Subscribe(topic2, handler2)
	assert.NoError(t, err)

	ctx := context.Background()
	msg := mockRawMessage{id: "1", name: "test"}

	handler1.On("HandleMessage", mock.Anything, mock.Anything).Return(nil)

	err = s.Publish(ctx, topic1, msg)
	assert.NoError(t, err)

	handler1.AssertExpectations(t)
	handler2.AssertNotCalled(t, "HandleMessage", mock.Anything, mock.Anything)
}

func TestStream_UnsubscribeAll(t *testing.T) {
	s := NewStream()
	topic := "test-topic"

	handler := &am.MockRawMessageHandler[am.IncomingRawMessage]{}

	_, err := s.Subscribe(topic, handler)
	assert.NoError(t, err)

	err = s.Unsubscribe()
	assert.NoError(t, err)

	ctx := context.Background()
	msg := mockRawMessage{
		id:   "1",
		name: "test",
	}

	err = s.Publish(ctx, topic, msg)
	assert.NoError(t, err)

	handler.AssertNotCalled(t, "HandleMessage", mock.Anything, mock.Anything)
}

func TestStream_Unsubscribe(t *testing.T) {
	s := NewStream()
	topic := "test-topic"
	ctx := context.Background()
	msg := mockRawMessage{id: "1", name: "test"}

	h1 := &am.MockRawMessageHandler[am.IncomingRawMessage]{}
	h2 := &am.MockRawMessageHandler[am.IncomingRawMessage]{}
	h1.On("HandleMessage", mock.Anything, mock.Anything).Return(nil)
	h2.On("HandleMessage", mock.Anything, mock.Anything).Return(nil)

	sub1, _ := s.Subscribe(topic, h1)
	_, _ = s.Subscribe(topic, h2)

	_ = s.Publish(ctx, topic, msg)
	_ = sub1.Unsubscribe()
	_ = s.Publish(ctx, topic, msg)

	h1.AssertNumberOfCalls(t, "HandleMessage", 1)
	h2.AssertNumberOfCalls(t, "HandleMessage", 2)
}
