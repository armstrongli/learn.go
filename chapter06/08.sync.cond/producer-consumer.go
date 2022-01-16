package main

import (
	"fmt"
	"sync"
	"time"
)

type Store struct {
	DataCount int
	Max       int
	lock      sync.Mutex
	pCond     *sync.Cond
	cCond     *sync.Cond
}

type Producer struct{}

func (Producer) Produce(s *Store) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.DataCount == s.Max {
		fmt.Println("生产者在等仓库拉货")
		s.pCond.Wait()
	}
	fmt.Println("开始生产+1")
	s.DataCount++
	s.cCond.Signal()
}

type Consumer struct{}

func (Consumer) Consume(s *Store) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.DataCount == 0 {
		fmt.Println("消费者在等货")
		s.cCond.Wait()
	}
	fmt.Println("消费者消费-1")
	s.DataCount--
	s.pCond.Signal()
}

func main() {
	s := &Store{
		Max: 10,
	}
	s.pCond = sync.NewCond(&s.lock)
	s.cCond = sync.NewCond(&s.lock)

	pCount, cCount := 50, 50
	for i := 0; i < pCount; i++ {
		go func() {
			for {
				time.Sleep(100 * time.Millisecond)
				Producer{}.Produce(s)
			}
		}()
	}
	for i := 0; i < cCount; i++ {
		go func() {
			for {
				time.Sleep(100 * time.Millisecond)
				Consumer{}.Consume(s)
			}
		}()
	}
	time.Sleep(1 * time.Second)
	fmt.Println(s.DataCount)
}
