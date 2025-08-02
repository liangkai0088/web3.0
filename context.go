package main

import (
	"context"
	"fmt"
	"reflect"
	"sync"
	"time"
)

func main2() {

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {

		defer wg.Done()

		watchDog1("【监控狗1】")

	}()

	wg.Wait()

}

func watchDog1(name string) {

	//开启for select循环，一直后台监控

	for {

		select {

		default:

			fmt.Println(name, "正在监控……")

		}

		time.Sleep(1 * time.Second)

	}

}

// 使用select + channel来中断监控
func main3() {

	var wg sync.WaitGroup

	wg.Add(1)

	stopCh := make(chan bool) //用来停止监控狗

	go func() {

		defer wg.Done()

		watchDog2(stopCh, "【监控狗1】")

	}()

	time.Sleep(5 * time.Second) //先让监控狗监控5秒

	stopCh <- true //发停止指令

	wg.Wait()

}

func watchDog2(stopCh chan bool, name string) {

	//开启for select循环，一直后台监控

	for {

		select {

		case <-stopCh:

			fmt.Println(name, "停止指令已收到，马上停止")

			return

		default:

			fmt.Println(name, "正在监控……")

		}

		time.Sleep(1 * time.Second)

	}

}

/**
这个示例是使用 select+channel 的方式改造的 watchDog 函数，实现了通过 channel 发送指令让监控狗停止，进而达到协程退出的目的。以上示例主要有两处修改，具体如下：

为 watchDog 函数增加 stopCh 参数，用于接收停止指令；
在 main 函数中，声明用于停止的 stopCh，传递给 watchDog 函数，然后通过 stopCh<-true 发送停止指令让协程退出。
*/

func main4() {

	var wg sync.WaitGroup

	wg.Add(1)

	ctx, stop := context.WithCancel(context.Background())

	go func() {

		defer wg.Done()

		watchDog(ctx, "【监控狗1】")

	}()

	time.Sleep(5 * time.Second) //先让监控狗监控5秒

	stop() //发停止指令

	wg.Wait()

}

func watchDog(ctx context.Context, name string) {

	//开启for select循环，一直后台监控

	for {

		select {

		case <-ctx.Done():

			fmt.Println(name, "停止指令已收到，马上停止")

			return

		default:

			fmt.Println(name, "正在监控……")

		}

		time.Sleep(1 * time.Second)

	}

}

/**
相比 select+channel 的方案，Context 方案主要有 4 个改动点。

watchDog 的 stopCh 参数换成了 ctx，类型为 context.Context。
原来的 case <-stopCh 改为 case <-ctx.Done()，用于判断是否停止。
使用 context.WithCancel(context.Background()) 函数生成一个可以取消的 Context，用于发送停止指令。这里的 context.Background() 用于生成一个空 Context，一般作为整个 Context 树的根节点。
原来的 stopCh <- true 停止指令，改为 context.WithCancel 函数返回的取消函数 stop()。
可以看到，这和修改前的整体代码结构一样，只不过从 channel 换成了 Context。以上示例只是 Context 的一种使用场景，它的能力不止于此，现在我来介绍什么是 Context。

*/

func main5() {
	ch := make(chan string)
	go func() {
		fmt.Println("飞雪无情")
		ch <- "goroutine 完成"
	}()
	fmt.Println("我是 main goroutine")
	v := <-ch
	fmt.Println("接收到的chan中的值为：", v)
}

type IntStrMap struct {
	m sync.Map
}

func (iMap *IntStrMap) Delete(key int) {
	iMap.m.Delete(key)
}

func (iMap *IntStrMap) Load(key int) (value string, ok bool) {
	v, ok := iMap.m.Load(key)
	if v != nil {
		value = v.(string)
	}
	return
}

func (iMap *IntStrMap) LoadOrStore(key int, value string) (actual string, loaded bool) {
	a, loaded := iMap.m.LoadOrStore(key, value)
	actual = a.(string)
	return
}

func (iMap *IntStrMap) Range(f func(key int, value string) bool) {
	f1 := func(key, value interface{}) bool {
		return f(key.(int), value.(string))
	}
	iMap.m.Range(f1)
}

func (iMap *IntStrMap) Store(key int, value string) {
	iMap.m.Store(key, value)
}

/**
新启动的 goroutine 中向 chan 类型的变量 ch 发送值；在 main goroutine 中，从变量 ch 接收值；如果 ch 中没有值，则阻塞等待到 ch 中有值可以接收为止。
*/

type ConcurrentMap struct {
	m         sync.Map
	keyType   reflect.Type
	valueType reflect.Type
}

func (cMap *ConcurrentMap) Load(key interface{}) (value interface{}, ok bool) {
	if reflect.TypeOf(key) != cMap.keyType {
		return
	}
	return cMap.m.Load(key)
}

func (cMap *ConcurrentMap) Store(key, value interface{}) {
	if reflect.TypeOf(key) != cMap.keyType {
		panic(fmt.Errorf("wrong key type: %v", reflect.TypeOf(key)))
	}
	if reflect.TypeOf(value) != cMap.valueType {
		panic(fmt.Errorf("wrong value type: %v", reflect.TypeOf(value)))
	}
	cMap.m.Store(key, value)
}
