package command

import (
	"net/http"
	"sync"
)

type (
	MetaStorage struct {
		sync.Mutex
		hash    string
		storage map[string]MetaHolder
	}

	MetaHolder struct {
		Cookies []*http.Cookie
		Header  http.Header
	}
)

var Storage = MetaStorage{storage: map[string]MetaHolder{}}

func (s *MetaStorage) SetHash(hash string) {
	s.Lock()
	defer s.Unlock()

	s.hash = hash
}

func (s *MetaStorage) Add(
	cookies []*http.Cookie,
	header http.Header,
) {
	s.Lock()
	defer s.Unlock()

	s.storage[s.hash] = MetaHolder{cookies, header}
}

func (s *MetaStorage) Get(hash string) MetaHolder {
	return s.storage[hash]
}
