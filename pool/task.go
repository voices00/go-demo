package pool

import (
	"fmt"
	"sync"
)

type Task func() (interface{}, error)

type TaskResult struct {
	Value interface{}
	Err   error
}

type ConcurrentTaskPool struct {
	wg        sync.WaitGroup
	functions []Task
	results   []TaskResult
	ch        chan struct{}
}

func NewDefaultConcurrentTaskPool() *ConcurrentTaskPool {
	return NewConcurrentTaskPool(10)
}

func NewConcurrentTaskPool(count int) *ConcurrentTaskPool {
	if count < 0 {
		count = 10
	}
	return &ConcurrentTaskPool{
		ch: make(chan struct{}, count),
	}
}

func (p *ConcurrentTaskPool) Add(tasks ...Task) {
	p.functions = append(p.functions, tasks...)
	p.wg.Add(len(tasks))
}

func (p *ConcurrentTaskPool) Wait() {
	p.wg.Wait()
}

func (p *ConcurrentTaskPool) Run() {
	p.results = make([]TaskResult, len(p.functions))
	for i, f := range p.functions {
		p.ch <- struct{}{}
		go func(index int, t Task) {
			defer func() {
				if err := recover(); err != nil {
					p.results[index].Err = fmt.Errorf("task run err:%v", err)
				}
				<-p.ch
			}()
			defer p.wg.Done()

			p.results[index].Value, p.results[index].Err = t()
		}(i, f)
	}
}

func (p *ConcurrentTaskPool) Results() []TaskResult {
	return p.results
}

func (p *ConcurrentTaskPool) Reset() {
	p.wg = sync.WaitGroup{}
	p.functions = []Task{}
	p.ch = make(chan struct{}, 10)
}
