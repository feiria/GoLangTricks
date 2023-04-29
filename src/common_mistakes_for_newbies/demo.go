package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime"
	unsafe "unsafe"
)

// 注意参数展开的问题。
func m1() {
	var a = []interface{}{1, 2, 3}

	fmt.Println(a)
	fmt.Println(a...)
	/**
	1. [1 2 3]
	2. 1 2 3
	*/
}

// 数组是值传递，无法通过修改数组类型的参数返回结果。
func m2() {
	x := [3]int{1, 2, 3}
	func(arr [3]int) {
		arr[0] = 7
		fmt.Println(arr)
	}(x)

	fmt.Println(x)

	/**
	[1, 2, 3]
	*/
}

// map是一种hash表实现，每次遍历的顺序都可能不一样。
func m3() {
	m := map[string]string{
		"1": "1",
		"2": "2",
		"3": "3",
	}

	for k, v := range m {
		println(k, v)
	}
}

// recover必须在defer中执行
func m4() {
	//defer recover()
	//panic(1) error

	//defer func() {
	//	func() {
	//		recover()
	//	}()
	//}()
	//panic(1) error

	defer func() {
		recover()
	}()
	panic(1)
}

// 避免独占CPU Goroutine是协作式抢占调度，Goroutine本身不会主动放弃CPU
func m5() {
	// 1
	runtime.GOMAXPROCS(1)
	//go func() {
	//	for i := 0; i < 10; i++ {
	//		fmt.Println(i)
	//	}
	//}()
	//for {
	//	runtime.Gosched()
	//}

	// 2
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
		os.Exit(0)
	}()
	select {}
}

// 闭包错误引用同一个变量
func m6() {
	//for i := 0; i < 3; i++ {
	//	defer func() {
	//		println(i)
	//	}()
	//}
	// error
	// 3, 3, 3

	//for i := 0; i < 3; i++ {
	//	defer func(i int) {
	//		println(i)
	//	}(i)
	//}
	// correct
	// 0, 1, 2
}

// defer在函数退出时才能执行，在for执行defer会导致资源延迟释放
func m7() {
	for i := 0; i < 3; i++ {
		f, err := os.Open("path")
		if err != nil {
			log.Fatal(err)
		}
		defer func(f *os.File) {
			err := f.Close()
			if err != nil {

			}
		}(f)
	} //error

	for i := 0; i < 3; i++ {
		func() {
			f, err := os.Open("path")
			if err != nil {
				log.Fatal(err)
			}
			defer func(f *os.File) {
				err := f.Close()
				if err != nil {

				}
			}(f)
		}()
	} // correct
}

// 切片会导致整个底层数组被锁定，无法释放内存。
func m8() {
	headerMap := make(map[string][]byte)
	for i := 0; i < 3; i++ {
		name := "/path"
		data, err := os.ReadFile(name)
		if err != nil {
			log.Fatal(err)
		}
		// headerMap[name] = data[:1] error
		headerMap[name] = append([]byte{}, data[:1]...) // correct
	}
}

// 对象的地址可能发生变化，因此指针不能从其它非指针类型的值生成
func m9() {
	var x int = 32
	var p uintptr = uintptr(unsafe.Pointer(&x))
	runtime.GC()
	var px *int = (*int)(unsafe.Pointer(p))
	println(*px)
	// 当内存发送变化的时候，相关的指针会同步更新，但是非指针类型的uintptr不会做同步更新。
	// 同理CGO中也不能保存Go对象地址。
}

func m10() {
	//ch := func() <-chan int {
	//	ch := make(chan int)
	//	go func() { // 无法回收
	//		for i := 0; ; i++ {
	//			ch <- i
	//		}
	//	}()
	//	return ch
	//}()
	//
	//for v := range ch {
	//	fmt.Println(v)
	//	if v == 5 {
	//		break
	//	}
	//}

	ctx, cancel := context.WithCancel(context.Background())
	ch := func(ctx context.Context) <-chan int {
		ch := make(chan int)
		go func() {
			for i := 0; ; i++ {
				select {
				case <-ctx.Done():
					return
				case ch <- i:
				}
			}
		}()
		return ch
	}(ctx)

	for v := range ch {
		fmt.Println(v)
		if v == 5 {
			break
		}
	}
	cancel()
}

func main() {
	m1()
	m2()
	m3()
	m4()
	m5()
	m6()
	m7()
	m8()
	m9()
	m10()
}
