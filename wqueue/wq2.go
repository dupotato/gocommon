package wqueue

import "time"

type Wq1 struct {
	maxlen int
	buffer chan interface{}
}

func NewWq1(maxlen int) *Wq1 {
	return &Wq1{
		maxlen: maxlen,
		buffer: make(chan interface{}, maxlen),
	}
}

func (n *Wq1) Put(data interface{}, timeout int) bool {
	timer := time.NewTimer(time.Duration(timeout) * time.Second)
	select {
	case n.buffer <- data:
		return true
	case <-timer.C:
		return false
	}
}

func (n *Wq1) Get(timeout int) (data interface{}, ok bool) {
	timer := time.NewTimer(time.Duration(timeout) * time.Second)
	select {
	case data <- n.buffer:
		return data, true
	case <-timer.C:
		return nil, false
	}
}
