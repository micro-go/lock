package lock

import (
	"sync/atomic"
)

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
