package main

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

// 原始的数据定义
type Meta struct {
	Name string
	Id   int64
}

var flag bool

// 生产者
func FetchMeta(cond *sync.Cond, meta *Meta) {
	cond.L.Lock()

	////模拟阻塞操作
	time.Sleep(time.Millisecond * 500) //调用rpc 、查询mysql
	meta.Name = "changed-by-fetch"
	meta.Id = 1234

	//看上去
	flag = true

	//试试换成 cond.Signal()
	cond.Broadcast()
	cond.L.Unlock()
}

func Consumer(cond *sync.Cond, meta *Meta) {
	cond.L.Lock()

	for !flag {
		cond.Wait()
	}

	metaJson, err := json.Marshal(meta)
	if err != nil {
		panic(err)
	}

	fmt.Println("res--->", string(metaJson))

	cond.L.Unlock()
}

func main() {

	mutex := &sync.Mutex{}
	cond := sync.NewCond(mutex)

	m := &Meta{
		Name: "init",
		Id:   0,
	}

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		FetchMeta(cond, m)
		wg.Done()
	}()

	go func() {
		Consumer(cond, m)
		wg.Done()
	}()

	wg.Wait()
}
