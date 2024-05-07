package simplemessagerouter

import (
	"slices"
	"testing"
)

func TestSimpleMessageRouter(t *testing.T) {
	t.Run("Happy Path: registered subscribers receive all messages", func(t *testing.T) {
		// [A]rrange
		router, err := NewSimpleMessageRouter()
		if err != nil {
			t.Errorf("Error creating router: %v", err)
		}
		key := RouteKey{"topic1", "string"}
		received1 := []string{}
		subscriber1 := func(m Message) error {
			t.Logf("Received message: %v", m)
			received1 = append(received1, m.(string))
			return nil
		}

		received2 := []string{}
		subscriber2 := func(m Message) error {
			t.Logf("Received message: %v", m)
			received2 = append(received2, m.(string))
			return nil
		}
		router.Subscribe(key, subscriber1)
		router.Subscribe(key, subscriber2)

		// [A]ct
		router.Publish(key, "message1")
		router.Publish(key, "message2")

		// [A]ssert
		have := received1
		want := []string{"message1", "message2"}
		if !slices.Equal(have, want) {
			t.Errorf("have %+v, want %+v", have, want)
		}

		have = received2
		want = []string{"message1", "message2"}
		if !slices.Equal(have, want) {
			t.Errorf("have %+v, want %+v", have, want)
		}
	})

	t.Run("subscriber under diff key not called", func(t *testing.T) {
		// [A]rrange
		router, err := NewSimpleMessageRouter()
		if err != nil {
			t.Errorf("Error creating router: %v", err)
		}

		key1 := RouteKey{"topic1", "string"}
		wasCalled1 := false
		subscriber1 := func(m Message) error {
			t.Logf("Received message: %v", m)
			wasCalled1 = true
			return nil
		}
		router.Subscribe(key1, subscriber1)

		key2 := RouteKey{"topic2", "string"}
		wasCalled2 := false
		subscriber2 := func(m Message) error {
			t.Logf("Received message: %v", m)
			wasCalled2 = true
			return nil
		}
		router.Subscribe(key2, subscriber2)

		// [A]ct
		router.Publish(key1, "message1")

		// [A]ssert
		if wasCalled1 == false {
			t.Errorf("Expected subscriber to be called")
		}
		if wasCalled2 == true {
			t.Errorf("Expected subscriber called")
		}
	})

	t.Run("no calls after unsub", func(t *testing.T) {
		// [A]rrange
		router, err := NewSimpleMessageRouter()
		if err != nil {
			t.Errorf("Error creating router: %v", err)
		}

		key := RouteKey{"topic1", "string"}
		received := []string{}
		subscriber := func(m Message) error {
			t.Logf("Received message: %v", m)
			received = append(received, m.(string))
			return nil
		}
		subid, _ := router.Subscribe(key, subscriber)

		// [A]ct
		router.Publish(key, "message1")
		router.Unsubscribe(key, subid)
		router.Publish(key, "message2")

		// [A]ssert
		have := received
		want := []string{"message1"}
		if !slices.Equal(have, want) {
			t.Errorf("have %+v, want %+v", have, want)
		}
	})
}
