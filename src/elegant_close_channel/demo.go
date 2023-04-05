package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

// Don't close a channel from the receiver side and don't close a channel if the channel has multiple concurrent senders
// Don't close (or send values to) closed Channels
/**
	操作		|	nil channel		|	closed chanel	|	not nil, not closed channel
-------------------------------------------------------------------------------------------------------------------
	close	|		panic		|		panic		|				正常关闭
	<- ch	|		block		|  读取到对应类型的零值	|	阻塞或正常读取数据。缓冲型channel为空或非缓缓冲型channel没有等待的发送者时会block
	ch <-	|		block		|		panic		|	阻塞或正常读取数据。非缓缓冲型channel没有等待的接收者或缓冲型channel buf满时时会block
*/

// 普通做法

func IsClose(ch <-chan any) bool {
	select {
	case <-ch:
		return true
	default:
	}
	return false
}

func UglyClose() {
	c := make(chan any)
	fmt.Println(IsClose(c))
	close(c)
	fmt.Println(IsClose(c))
}

// 通过Recover进行关闭

func IsCloseByRecover(ch chan any, value any) (closed bool) {
	defer func() {
		if recover() != nil {
			closed = true
		}
	}()
	ch <- value
	return false
}

func SafeClose(ch chan any) (justClosed bool) {
	defer func() {
		if recover() != nil {
			justClosed = false
		}
	}()
	close(ch)
	return true
}

func UglyClose2() {
	c := make(chan any)
	fmt.Println(IsCloseByRecover(c, 1))
	SafeClose(c)
	fmt.Println(IsCloseByRecover(c, 1))
}

// 通过sync.Once关闭

type MyChannel struct {
	C    chan any
	once sync.Once
}

func (mc *MyChannel) SafeCloseByOnce() {
	mc.once.Do(func() {
		close(mc.C)
	})
}

func UglyClose3() {
	mc := &MyChannel{C: make(chan any)}
	mc.SafeCloseByOnce()
}

/**
1) one sender, one receiver
2) one sender, many receiver
3) many sender, one receiver
4) many sender, many receiver

1、2直接在发送端关闭， 以下是3、4两种情况
*/

// 引入信号通道，在接收端关闭信号通道发送信号给发送端。注意dataCh没有显式关闭，由于没有任何goroutine引用dataCh，最后GC会自动回收

func ElegantCloseForStage3() {
	rand.Seed(time.Now().UnixNano())
	const Max = 100
	const NumberSenders = 10

	dataCh := make(chan int, 100)
	stopCh := make(chan struct{})

	// senders
	for i := 0; i < NumberSenders; i++ {
		go func() {
			for {
				select {
				case <-stopCh:
					return
				case dataCh <- rand.Intn(Max):
				}
			}
		}()
	}

	// the receiver
	go func() {
		for i := range dataCh {
			if i == Max-1 {
				fmt.Println("send stop signal to senders")
				close(stopCh)
				return
			}
			fmt.Println(i)
		}
	}()
}

// 对于情况四，因为有多个receiver，采用stage3会多次关闭同一个channel，导致panic

func ElegantCloseForStage4() {
	rand.Seed(time.Now().UnixNano())
	const Max = 100
	const NumberSenders = 10
	const NumberReceivers = 20

	dataCh := make(chan int, 100)
	stopCh := make(chan struct{})

	toStop := make(chan string, 1)
	var stoppedBy string

	go func() {
		stoppedBy = <-toStop
		fmt.Println(stoppedBy)
		close(stopCh)
	}()

	// senders
	for i := 0; i < NumberSenders; i++ {
		go func(id string) {
			value := rand.Intn(Max)
			if value == 0 {
				select {
				case toStop <- "sender#" + id:
				default:
				}
				return
			}

			select {
			case <-stopCh:
				return
			case dataCh <- value:
			}
		}(strconv.Itoa(i))
	}

	// receivers
	for i := 0; i < NumberReceivers; i++ {
		go func(id string) {
			for {
				select {
				case <-stopCh:
					return
				case value := <-dataCh:
					if value == Max-1 {
						select {
						case toStop <- "receiver#" + id:
						default:
						}
						return
					}
					fmt.Println(value)
				}
			}
		}(strconv.Itoa(i))
	}
}

func MoreElegantCloseByContext() {

}

func main() {
	UglyClose()
	UglyClose2()
	UglyClose3()
	ElegantCloseForStage3()
	ElegantCloseForStage4()
	MoreElegantCloseByContext()
}
