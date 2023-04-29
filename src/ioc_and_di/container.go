package main

import (
	"fmt"
	"reflect"
	"sync"
)

var iContainer *container

// container
type container struct {
	Objects     []*Object
	tidyObjects sync.Map
	instances   sync.Map
}

func init() {
	iContainer = New()
}

func New() *container {
	return &container{}
}

func Container() *container {
	return iContainer
}

func Provide(objects ...*Object) error {
	return iContainer.Provide(objects...)
}

func Populate(name string, ptr interface{}) error {
	return iContainer.Populate(name, ptr)
}

func (t *container) Provide(objects ...*Object) error {
	for _, o := range objects {
		// 检查map中是否已经加载该对象
		if _, ok := t.tidyObjects.Load(o.Name); ok {
			return fmt.Errorf("error: object '%s' existing", o.Name)
		}
		// 写入对象
		t.tidyObjects.Store(o.Name, o)
	}
	return nil
}

func (t *container) Object(name string) (*Object, error) {
	v, ok := t.tidyObjects.Load(name)
	if !ok {
		return nil, fmt.Errorf("error: object '%s' not found", name)
	}
	obj := v.(*Object)
	return obj, nil
}

func (t *container) Populate(name string, ptr interface{}) error {
	obj, err := t.Object(name)
	if err != nil {
		return err
	}
	ptrCopy := func(to, from interface{}) {
		reflect.ValueOf(to).Elem().Set(reflect.ValueOf(from))
	}
	if !obj.NewEveryTime {
		refresher := &obj.refresher
		if p, ok := t.instances.Load(name); ok && !refresher.status() {
			ptrCopy(ptr, p)
			return nil
		}
		// 处理并发穿透
		var e error
		obj.once.Do(func() {
			v, err := obj.New()
			if err != nil {
				e = err
				return
			}
			t.instances.Store(name, v)
			refresher.off()
		})
		if e != nil {
			obj.once = sync.Once{}
			return e
		}
		p, _ := t.instances.Load(name)
		ptrCopy(ptr, p)
		return e
	} else {
		v, err := obj.New()
		if err != nil {
			return err
		}
		ptrCopy(ptr, v)
	}
	return nil
}
