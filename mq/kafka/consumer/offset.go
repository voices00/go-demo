package main

import (
	"fmt"
	"sync"
)

var consumeOffset = &consumeOffsetManager{
	mu:        sync.Mutex{},
	offsetMap: map[string]int64{},
}

type consumeOffsetManager struct {
	mu        sync.Mutex
	offsetMap map[string]int64
}

func (w *consumeOffsetManager) set(topic string, partition int32, val int64) {
	w.mu.Lock()
	defer w.mu.Unlock()

	w.offsetMap[fmt.Sprintf("%s_%d", topic, partition)] = val
	return
}

func (w *consumeOffsetManager) get(topic string, partition int32) int64 {
	w.mu.Lock()
	defer w.mu.Unlock()
	if offset, exist := w.offsetMap[fmt.Sprintf("%s_%d", topic, partition)]; exist {
		return offset
	}
	return 0
}
