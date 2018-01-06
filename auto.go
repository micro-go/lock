package lock

import (
	"sync"
)

// Scoped-lock style conveniences, for one line lock/auto-unlock.

// sync.RWMutex examples:
//		defer lock.Read(&rwmutex).Unlock()
//		defer lock.Write(&rwmutex).Unlock()

// sync.Mutex examples:
//		defer lock.Mutex(&mutex).Unlock()


// struct ReadAuto manages an automatic RWMutex read lock.
type ReadAuto struct {
	rw *sync.RWMutex
}

func (l ReadAuto) Unlock() {
	l.rw.RUnlock()
}

func Read(rw *sync.RWMutex) ReadAuto {
	rw.RLock()
	return ReadAuto{rw}
}

// struct WriteAuto manages an automatic RWMutex write lock.
type WriteAuto struct {
	rw *sync.RWMutex
}

func (l WriteAuto) Unlock() {
	l.rw.Unlock()
}

func Write(rw *sync.RWMutex) WriteAuto {
	rw.Lock()
	return WriteAuto{rw}
}

// struct MutexAuto manages an automatic Mutex lock.
type MutexAuto struct {
	m *sync.Mutex
}

func (l MutexAuto) Unlock() {
	l.m.Unlock()
}

func Mutex(m *sync.Mutex) MutexAuto {
	m.Lock()
	return MutexAuto{m}
}
