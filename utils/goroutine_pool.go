package utils

import (
	"fmt"
	"sync"

	"github.com/panjf2000/ants"
)

type PoolInfo struct {
	Pool          *ants.Pool
	MaxWorkers    int
	TaskWaitGroup sync.WaitGroup
}

func (pool *PoolInfo) NewPool(num int) {
	p, err := ants.NewPool(num)
	if err != nil {
		fmt.Println("Create Pool Error:", err)
		return
	}
	pool.Pool = p
	pool.MaxWorkers = num
}
func (pool *PoolInfo) AddTask(fun func()) {
	pool.Pool.Submit(fun)
}
