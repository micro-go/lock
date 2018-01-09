package lock

import (
	"sync/atomic"
)

type AtomicBool interface {
	SetTo(b bool)
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

func (a *atomicBool) IsSet() bool {
	return atomic.LoadInt32(&a.val) != 0
}
