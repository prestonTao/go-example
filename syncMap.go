package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	simple1()
	simple2()
}

func simple2() {
	t1 := time.Now()
	mstore := MapPri{
		lock:  new(sync.RWMutex),
		store: make(map[string]int),
	}
	mstore.add("tao", 1)
	for i := 0; i < 10000000; i++ {
		value, _ := mstore.find("tao")
		mstore.add("tao", value+1)
		// func(mstore *MapPri) {
		// 	mstore.add(key, value)
		// }(&mstore)
	}
	t2 := time.Now()
	fmt.Println("第二个测试结果", t2.Sub(t1).Nanoseconds()/10000000)

}

type MapPri struct {
	lock  *sync.RWMutex
	store map[string]int
}

func (this MapPri) add(key string, value int) {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.store[key] = value
}
func (this MapPri) find(key string) (int, bool) {
	this.lock.Lock()
	defer this.lock.Unlock()
	value, ok := this.store[key]
	return value, ok
}

//--------------------------------------
func simple1() {
	t1 := time.Now()
	m := New()
	m.Insert("tao", 1)
	for i := 0; i < 10000000; i++ {
		func() {
			m.Update("tao", func(value interface{}, ok bool) interface{} {
				result := value.(int) + 1
				return result
			})
		}()
	}

	result, _ := m.Find("tao")
	t2 := time.Now()
	fmt.Println("第二个测试结果", t2.Sub(t1).Nanoseconds()/10000000)
	fmt.Println(result)
	// time.Sleep(time.Second * 2)
}

//-------------------------------------------------
//-------------------------------------------------
//-------------------------------------------------

type UpdateFunc func(interface{}, bool) interface{}

type SafeMap interface {
	Insert(string, interface{})
	Delete(string)
	Find(string) (interface{}, bool)
	Len() int
	Update(string, UpdateFunc)
	Close() map[string]interface{}
}

type commandAction int

const (
	remove commandAction = iota
	end
	find
	insert
	length
	update
)

type commandData struct {
	action  commandAction
	key     string
	value   interface{}
	result  chan<- interface{}
	data    chan<- map[string]interface{}
	updater UpdateFunc
}

type safeMap chan commandData

func (this safeMap) Insert(key string, value interface{}) {
	this <- commandData{action: insert, key: key, value: value}
}
func (this safeMap) Delete(key string) {
	this <- commandData{action: remove, key: key}
}

type findResult struct {
	value interface{}
	found bool
}

func (this safeMap) Find(key string) (value interface{}, found bool) {
	reply := make(chan interface{})
	this <- commandData{action: find, key: key, result: reply}
	result := (<-reply).(findResult)
	return result.value, result.found
}

func (this safeMap) Len() int {
	reply := make(chan interface{})
	this <- commandData{action: length, result: reply}
	return (<-reply).(int)
}

func (this safeMap) Update(key string, updater UpdateFunc) {
	this <- commandData{action: update, key: key, updater: updater}
}

func (this safeMap) Close() map[string]interface{} {
	reply := make(chan map[string]interface{})
	this <- commandData{action: end, data: reply}
	return <-reply
}

func (this safeMap) run() {
	store := make(map[string]interface{})
	for command := range this {
		switch command.action {
		case insert:
			store[command.key] = command.value
		case remove:
			delete(store, command.key)
		case find:
			value, found := store[command.key]
			command.result <- findResult{value, found}
		case length:
			command.result <- len(store)
		case update:
			value, found := store[command.key]
			store[command.key] = command.updater(value, found)
		case end:
			close(this)
			command.data <- store
		}
	}
}

func New() SafeMap {
	sm := make(safeMap)
	go sm.run()
	return sm
}
