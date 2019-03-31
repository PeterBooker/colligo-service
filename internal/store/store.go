package store

import (
	"sync"
	"time"

	"github.com/mrjones/oauth"
)

const (
	ttl = time.Minute * 15
)

// Store ...
type Store struct {
	items map[string]*Visitor
	sync.RWMutex
}

// Visitor ...
type Visitor struct {
	Token     *oauth.RequestToken
	OriginURL string
	created   int64
}

// New returns a new Store
func New() *Store {
	s := &Store{
		items: make(map[string]*Visitor, 10),
	}
	go s.Purge()

	return s
}

// Purge ...
func (s *Store) Purge() {
	for now := range time.Tick(time.Minute) {
		s.Lock()
		for k, v := range s.items {
			if now.Unix()-v.created > int64(ttl) {
				s.Del(k)
			}
		}
		s.Unlock()
	}
}

// Get ...
func (s *Store) Get(key string) (*Visitor, bool) {
	s.RLock()
	v, ok := s.items[key]
	s.RUnlock()

	return v, ok
}

// Set ...
func (s *Store) Set(key string, v *Visitor) {
	s.Lock()
	s.items[key] = v
	s.Unlock()
}

// Del ...
func (s *Store) Del(key string) {
	s.Lock()
	delete(s.items, key)
	s.Unlock()
}

// Exists ...
func (s *Store) Exists(key string) bool {
	s.RLock()
	_, ok := s.items[key]
	s.RUnlock()

	return ok
}
