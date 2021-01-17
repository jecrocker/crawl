package utils

import "sync"

// Tracker tracks the system
type Tracker struct {
	mu     *sync.Mutex
	data   map[string]string
	queued int
}

// NewTracker returns a new tracker
func NewTracker() Tracker {
	return Tracker{
		mu:     &sync.Mutex{},
		data:   map[string]string{},
		queued: 0,
	}
}

// HasItem checks if an item already exists on the tracker
func (t *Tracker) HasItem(check string) bool {
	t.mu.Lock()
	_, ok := t.data[check]
	t.mu.Unlock()
	return ok
}

// AddURL to the data
func (t *Tracker) AddURL(url string) {
	t.mu.Lock()
	t.data[url] = url
	t.mu.Unlock()
}

// AddWork to the queue
func (t *Tracker) AddWork() {
	t.mu.Lock()
	t.queued++
	t.mu.Unlock()
}

// CompleteWork remove from the queue
func (t *Tracker) CompleteWork() {
	t.mu.Lock()
	t.queued--
	t.mu.Unlock()
}
