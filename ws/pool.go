package ws

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"log"
	"sync"
	"time"
)

type wsPools struct {
	Pool *ants.Pool
	wg   sync.WaitGroup
}

var wp *wsPools
var once sync.Once

func GetInstance() *wsPools {
	once.Do(func() {
		wp = new(wsPools)
	})
	return wp
}

func (wp *wsPools) InitPool(size int, options ...ants.Option) {
	defer ants.Release()
	wp.Pool, _ = ants.NewPool(size, options...)
}

func (wp *wsPools) NewTask(c *Client, task func()) {
	wp.wg.Add(1)
	err := wp.Pool.Submit(task)
	if err != nil {
		c.Message <- []byte("请求过于频繁,请稍后再试")
		wp.wg.Done()
	}
}

func (wp *wsPools) Close() {
	wp.Pool.Release()
	wp.wg.Wait()
}

type Message struct {
	Client  *Client
	Message []byte
}

func (wp *wsPools) Process(msg Message) func() {
	return func() {
		fmt.Println(msg.Message)
		time.Sleep(time.Second * 5)
		log.Println("ok")
		wp.wg.Done()
	}
}
