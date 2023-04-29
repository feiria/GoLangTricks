package main

import (
	"fmt"
	"sync"
)

type Object struct {
	Name string
	// 创建对象的闭包
	New func() (any, error)
	// 是否每次都创建新的对象
	NewEveryTime bool
	refresher    refresher
	once         sync.Once
}

func (t *Object) Refresh() error {
	if t.NewEveryTime {
		return fmt.Errorf("error: '%s' is NewEverytime, unable to refresh", t.Name)
	}
	t.once = sync.Once{}
	t.refresher.on()
	return nil
}

type refresher struct {
	val bool
}

func (t *refresher) on() {
	t.val = true
}

func (t *refresher) off() {
	t.val = false
}

func (t *refresher) status() bool {
	return t.val
}
