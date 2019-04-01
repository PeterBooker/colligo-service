package store

import (
	"sync"
	"time"

	"github.com/mrjones/oauth"
)

const (
	ttl = time.Minute * 15
)

// Store stores visitor data
// Used to process oAuth authentication
type Store struct {
	items map[string]*Visitor
	sync.RWMutex
}

// Visitor represents an API visitor
// Eventually redirected back to Callback URL
type Visitor struct {
	Token    *oauth.RequestToken
	Callback string
	created  int64
}

// New returns a new Store
func New() *Store {
	s := &Store{
		items: make(map[string]*Visitor, 10),
	}
	go s.Purge()

	return s
}

// Purge remove old visitor data
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

// Get fetches visitor data
func (s *Store) Get(key string) (*Visitor, bool) {
	s.RLock()
	v, ok := s.items[key]
	s.RUnlock()

	return v, ok
}

// Set stores visitor data
func (s *Store) Set(key string, v *Visitor) {
	v.created = time.Now().Unix()
	s.Lock()
	s.items[key] = v
	s.Unlock()
}

// Del removes visitor data
func (s *Store) Del(key string) {
	s.Lock()
	delete(s.items, key)
	s.Unlock()
}

// Exists checks for existence of visitor data
func (s *Store) Exists(key string) bool {
	s.RLock()
	_, ok := s.items[key]
	s.RUnlock()

	return ok
}
