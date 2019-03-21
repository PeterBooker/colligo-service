package auth

import (
	"sync"

	"github.com/mrjones/oauth"
)

// Tokens ...
type Tokens struct {
	List map[string]*oauth.RequestToken
	sync.RWMutex
}

var (
	tokens *Tokens
)

// Setup ...
func Setup() {
	tokens = &Tokens{
		List: make(map[string]*oauth.RequestToken),
	}
}

// Get ...
func (t *Tokens) Get(key string) (*oauth.RequestToken, bool) {
	t.RLock()
	rt, ok := t.List[key]
	t.RUnlock()

	return rt, ok
}

// Set ...
func (t *Tokens) Set(key string, token *oauth.RequestToken) {
	t.Lock()
	t.List[key] = token
	t.Unlock()
}

// Del ...
func (t *Tokens) Del(key string) {
	t.Lock()
	delete(t.List, key)
	t.Unlock()
}

// Exists ...
func (t *Tokens) Exists(key string) bool {
	t.RLock()
	_, ok := t.List[key]
	t.RUnlock()

	return ok
}
