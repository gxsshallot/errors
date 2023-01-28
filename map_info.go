package errors

import "sync"

// map for code and message
var GlobalCodes SyncCodeMessageMap

// map for subCode and message
var GlobalSubCodes SyncCodeMessageMap

type SyncCodeMessageMap struct {
	m    map[int]string
	lock sync.RWMutex
}

func (s *SyncCodeMessageMap) Add(code int, message string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.m == nil {
		s.m = make(map[int]string)
	}
	s.m[code] = message
}

func (s *SyncCodeMessageMap) Del(code int) {
	s.lock.Lock()
	defer s.lock.Unlock()
	delete(s.m, code)
}

func (s *SyncCodeMessageMap) Get(code int) string {
	s.lock.RLock()
	defer s.lock.RUnlock()
	message, ok := s.m[code]
	if !ok {
		return ""
	}
	return message
}
