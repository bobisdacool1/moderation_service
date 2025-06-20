package inmemadapter

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/segmentio/kafka-go"
)

type (
	Adapter struct {
		messagesByID map[string]CachedMessage

		mu              sync.RWMutex
		ttl             time.Duration
		cleanupInterval time.Duration
		limit           int

		startOnce sync.Once
		stopOnce  sync.Once
		cancel    context.CancelFunc
	}

	CachedMessage struct {
		Message   kafka.Message
		Timestamp time.Time
	}
)

var (
	errLimitExceeded       = errors.New("limit exceeded")
	errDurationBelowZero   = errors.New("time.Duration must be > 0")
	errMessageAlreadyInUse = errors.New("message already in use")
)

func NewModerationRequestAdapter(ttl time.Duration, cleanupInterval time.Duration) (*Adapter, error) {
	if ttl <= 0 || cleanupInterval <= 0 {

		return nil, errDurationBelowZero
	}

	a := &Adapter{
		messagesByID:    make(map[string]CachedMessage),
		mu:              sync.RWMutex{},
		ttl:             ttl,
		cleanupInterval: cleanupInterval,
	}

	return a, nil
}

func (a *Adapter) Put(ctx context.Context, id string, message kafka.Message) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.limit > 0 && len(a.messagesByID) >= a.limit {
		return errLimitExceeded
	}

	if _, ok := a.messagesByID[id]; ok {
		if !shouldPut(a.messagesByID[id], message) {
			return errMessageAlreadyInUse
		}
	}

	t := time.Now()
	a.messagesByID[id] = CachedMessage{
		Message:   message,
		Timestamp: t,
	}

	return nil
}

func (a *Adapter) Get(ctx context.Context, id string) (kafka.Message, bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()

	if v, ok := a.messagesByID[id]; ok {
		return v.Message, true
	}

	return kafka.Message{}, false
}

func (a *Adapter) Delete(ctx context.Context, id string) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if _, ok := a.messagesByID[id]; ok {
		delete(a.messagesByID, id)
	}
}

func (a *Adapter) Start(ctx context.Context) {
	a.startOnce.Do(func() {
		childCtx, cancel := context.WithCancel(ctx)
		a.cancel = cancel

		go a.runCleanupLoop(childCtx)
	})
}

func (a *Adapter) Stop() {
	a.stopOnce.Do(func() {
		if a.cancel != nil {
			a.cancel()
		}
	})
}

func shouldPut(stored CachedMessage, new kafka.Message) bool {
	return stored.Message.Offset <= new.Offset &&
		stored.Message.Partition == new.Partition &&
		stored.Message.Topic == new.Topic
}

func (a *Adapter) cleanup() {
	expiredKeys := make([]string, 0, len(a.messagesByID)/4)
	t := time.Now()
	a.mu.RLock()
	for k, v := range a.messagesByID {
		if v.Timestamp.Add(a.ttl).Before(t) {
			expiredKeys = append(expiredKeys, k)
		}
	}
	a.mu.RUnlock()

	a.mu.Lock()
	for _, v := range expiredKeys {
		delete(a.messagesByID, v)
	}
	a.mu.Unlock()
}

func (a *Adapter) runCleanupLoop(ctx context.Context) {
	var loop func()
	loop = func() {
		if ctx.Err() != nil {
			return
		}

		a.cleanup()
		time.AfterFunc(a.cleanupInterval, loop)
	}

	loop()
}
