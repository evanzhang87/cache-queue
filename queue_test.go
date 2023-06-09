package queue

import (
	"strconv"
	"testing"
	"time"
)

func TestQueue(t *testing.T) {
	queue := NewQueue(5)
	for i := 0; i < 10; i++ {
		e := strconv.Itoa(i)
		queue.Add(e)
		time.Sleep(time.Second)
	}

	for _, q := range queue.array {
		t.Log(q)
	}
}
