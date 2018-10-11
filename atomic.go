package lock

import (
	"sync/atomic"
)

// --------------------------------
// ATOMIC-BOOL

type AtomicBool interface {
	SetTo(b bool)
	IsTrue() bool

	// I think deprecate this -- I tend to get confused on the meaning in use.
	IsSet() bool
}

func NewAtomicBool() AtomicBool {
	return &atomicBool{val: 0}
}

type atomicBool struct {
	val int32
}

func (a *atomicBool) SetTo(b bool) {
	var newVal int32 = 1
	if !b {
		newVal = 0
	}
	atomic.StoreInt32(&a.val, newVal)
}

func (a *atomicBool) IsTrue() bool {
	return atomic.LoadInt32(&a.val) != 0
}

// I think deprecate this -- I tend to get confused on the meaning in use.
func (a *atomicBool) IsSet() bool {
	return atomic.LoadInt32(&a.val) != 0
}

// --------------------------------
// ATOMIC-INT32

type AtomicInt32 interface {
	// Get() the current value.
	Get() int32
	// SetTo() the new value.
	SetTo(newval int32)
	// Add() adds the delta, returning the new value.
	Add(delta int32) int32
	// TrySetTo() sets to the new value only if the current value is compareval. Answer true if the set succeeded.
	TrySetTo(newval, compareval int32) bool
}

func NewAtomicInt32() AtomicInt32 {
	return &AtomicInt32_t{val: 0}
}

type AtomicInt32_t struct {
	val int32
}

func (a *AtomicInt32_t) Get() int32 {
	return atomic.LoadInt32(&a.val)
}

func (a *AtomicInt32_t) SetTo(newval int32) {
	atomic.StoreInt32(&a.val, newval)
}

func (a *AtomicInt32_t) Add(delta int32) int32 {
	return atomic.AddInt32(&a.val, delta)
}

func (a *AtomicInt32_t) TrySetTo(newval, compareval int32) bool {
	return atomic.CompareAndSwapInt32(&a.val, compareval, newval)
}
