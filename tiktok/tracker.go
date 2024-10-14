package tiktok

import (
	"sync"
	"time"
)

type request struct {
	tikTokURL string
	addedAt   time.Time
	username  string
}

type requestsTracker struct {
	inProgress map[string]request
	mutex      sync.RWMutex
}

func newTracker() *requestsTracker {
	return &requestsTracker{
		inProgress: make(map[string]request, 16),
	}
}

func (t *requestsTracker) track(url string, username string) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.inProgress[url] = request{
		tikTokURL: url,
		addedAt:   time.Now(),
		username:  username,
	}
}

func (t *requestsTracker) untrack(url string) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	delete(t.inProgress, url)
}

func (t *requestsTracker) inProgressCount() int {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	return len(t.inProgress)
}

func (t *requestsTracker) getInProgress() []request {
	t.mutex.RLock()
	defer t.mutex.RUnlock()

	result := make([]request, 0, len(t.inProgress))
	for _, value := range t.inProgress {
		result = append(result, value)
	}

	return result
}
