package pool

import (
	"fmt"
	"testing"
	"time"
)

func TestConcurrentTaskPool(t *testing.T) {
	pool := NewDefaultConcurrentTaskPool()
	for i := 0; i < 100; i++ {
		j := i
		pool.Add(
			func() (value interface{}, err error) {
				time.Sleep(time.Second)
				fmt.Println(j)
				return j, nil
			},
		)
	}

	pool.Run()
	pool.Wait()

	pool.Reset()
	for i := 0; i < 100; i++ {
		j := i
		pool.Add(
			func() (value interface{}, err error) {
				time.Sleep(time.Second)
				fmt.Println(j)
				return j, nil
			},
		)
	}
	pool.Run()
	pool.Wait()

	fmt.Println("end..")
}
