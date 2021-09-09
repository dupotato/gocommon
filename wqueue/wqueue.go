package wqueue

import (
	"runtime"
	"sync/atomic"
)

type Wqueue struct {
	capNum uint32
	capMod uint32
	getPos uint32
	putPos uint32
	cache  []wqcache
}

type wqcache struct {
	value interface{}
}

func NewQueue(capnum uint32) *Wqueue {
	q := &Wqueue{
		capNum: capnum,
		capMod: capnum - 1,
		cache:  make([]wqcache, capnum),
	}
	return q
}

func (w *Wqueue) Get() (val interface{}, ok bool) {
	var putPos, getPos, getPosNew, posCnt uint32

	for {
		putPos = w.putPos
		getPos = w.getPos

		if putPos >= getPos {
			posCnt = putPos - getPos
		} else {
			posCnt = w.capMod + putPos - getPos
		}

		if posCnt < 1 {
			runtime.Gosched()
			return nil, false
		}

		getPosNew = getPos + 1
		if atomic.CompareAndSwapUint32(&w.getPos, getPos, getPosNew) {
			return w.cache[getPosNew], true
		} else {
			runtime.Gosched()
		}
	}
}
