package inmemadapter

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/segmentio/kafka-go"

	"ModerationService/internal/config"
	"ModerationService/internal/entity"
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
		wg        sync.WaitGroup
	}

	CachedMessage struct {
		Message   *kafka.Message
		Timestamp time.Time
	}
)

var (
	errLimitExceeded       = fmt.Errorf("limit exceeded")
	errDurationBelowZero   = fmt.Errorf("time.Duration must be > 0")
	errMessageAlreadyInUse = fmt.Errorf("message already in use")
)

func NewInMemModerationRequestAdapter(cfg *config.Config) (*Adapter, error) {
	if cfg.InMem.TTL <= 0 || cfg.InMem.CleanupInterval <= 0 || cfg.InMem.Limit <= 0 {

		return nil, errDurationBelowZero
	}

	a := &Adapter{
		messagesByID:    make(map[string]CachedMessage),
		mu:              sync.RWMutex{},
		ttl:             cfg.InMem.TTL,
		cleanupInterval: cfg.InMem.CleanupInterval,
		limit:           cfg.InMem.Limit,
	}

	return a, nil
}

func (a *Adapter) Put(ctx context.Context, id string, message entity.KafkaMessageEnvelope) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	rawMessage := entity.KafkaEnvelopeToMessage(message)

	if a.limit > 0 && len(a.messagesByID) >= a.limit {
		return errLimitExceeded
	}

	if _, ok := a.messagesByID[id]; ok {
		if !shouldPut(a.messagesByID[id], rawMessage) {
			return errMessageAlreadyInUse
		}
	}

	t := time.Now()
	a.messagesByID[id] = CachedMessage{
		Message:   &rawMessage,
		Timestamp: t,
	}

	return nil
}

func (a *Adapter) Get(ctx context.Context, id string) (entity.KafkaMessageEnvelope, bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()

	if v, ok := a.messagesByID[id]; ok {
		return entity.KafkaMessageToEnvelope(*v.Message), true
	}

	return entity.KafkaMessageEnvelope{}, false
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

		a.wg.Add(1)
		go a.runCleanupLoop(childCtx)
	})
}

func (a *Adapter) Stop() {
	a.stopOnce.Do(func() {
		if a.cancel != nil {
			a.cancel()
			a.wg.Wait()
		}
	})
}

func shouldPut(stored CachedMessage, new kafka.Message) bool {
	if stored.Message == nil {
		return true
	}
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
	defer a.wg.Done()

	timer := time.NewTimer(a.cleanupInterval)
	defer timer.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-timer.C:
			a.cleanup()
			timer.Reset(a.cleanupInterval)
		}
	}
}
