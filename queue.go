package queue

import (
	"context"
	"sync"
)

type Queue struct {
	maxSize int
	array   []interface{}
	head    int
	tail    int
	mut     *sync.Mutex
	full    bool

	ctx    context.Context
	cancel func()
}

func NewQueue(max int) Queue {
	ctx, cancel := context.WithCancel(context.TODO())
	return Queue{
		maxSize: max,
		array:   make([]interface{}, max),
		head:    0,
		tail:    0,
		mut:     &sync.Mutex{},
		full:    false,
		ctx:     ctx,
		cancel:  cancel,
	}
}

func (q *Queue) Add(val interface{}) {
	q.mut.Lock()
	defer q.mut.Unlock()
	q.array[q.tail] = val
	q.tail++
	if q.tail == q.maxSize {
		q.tail = 0
		q.full = true
	}

	if q.full {
		q.head++
		if q.head == q.maxSize {
			q.head = 0
		}
	}
	return
}

func (q *Queue) Get() interface{} {
	q.mut.Lock()
	defer q.mut.Unlock()
	val := q.array[q.head]
	q.array[q.head] = nil
	q.head++
	if q.head == q.maxSize {
		q.head = 0
	}
	return val
}

func (q *Queue) Len() int {
	return len(q.array)
}

func (q *Queue) ConsumeChan(cacheChan chan interface{}) {
	for {
		select {
		case event := <-cacheChan:
			q.Add(event)
		case <-q.ctx.Done():
			return
		}
	}
}

func (q *Queue) StopConsume() {
	q.cancel()
}
