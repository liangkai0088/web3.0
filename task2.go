package main

import (
	"fmt"
	"math"
	"sync"
	"sync/atomic"
	"time"
)

// increaseByTen 接收一个整数指针作为参数，将该指针指向的值增加10
func increaseByTen(num *int) {
	*num += 10
}

func doubleSlice(slice *[]int) {
	// 解引用指针获取切片
	s := *slice

	// 遍历切片，将每个元素乘以2
	for i := 0; i < len(s); i++ {
		s[i] *= 2
	}
}

// 交替打印奇数和偶数的程序
func mainTask13() {
	// 创建两个通道用于协程间通信
	oddCh := make(chan int)
	evenCh := make(chan int)

	// 创建WaitGroup等待协程完成
	var wg sync.WaitGroup
	wg.Add(2)

	// 启动打印奇数的协程
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i += 2 {
			<-oddCh // 等待奇数通道信号
			fmt.Printf("奇数: %d\n", i)
			evenCh <- i // 发送信号给偶数通道
		}
	}()

	// 启动打印偶数的协程
	go func() {
		defer wg.Done()
		for i := 2; i <= 10; i += 2 {
			num := <-evenCh // 等待偶数通道信号
			fmt.Printf("偶数: %d\n", i)
			if i < 10 {
				oddCh <- num
			}
		}
	}()

	// 启动打印序列，先打印奇数
	oddCh <- 0

	// 等待所有协程完成
	wg.Wait()
}

func mainTask12() {
	// 创建一个整数切片
	numbers := []int{1, 2, 3, 4, 5}

	fmt.Println("修改前的切片:", numbers)

	// 调用函数，传入切片的指针
	doubleSlice(&numbers)

	fmt.Println("修改后的切片:", numbers)
}

func mainTask() {
	// 定义一个整数变量
	value := 5
	fmt.Printf("修改前的值: %d\n", value)

	// 调用函数，传入value的地址
	increaseByTen(&value)

	// 输出修改后的值
	fmt.Printf("修改后的值: %d\n", value)
}

// Shape 接口定义，包含Area和Perimeter两个方法
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle 结构体表示矩形
type Rectangle struct {
	Width  float64
	Height float64
}

// Rectangle实现Shape接口的Area方法
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Rectangle实现Shape接口的Perimeter方法
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Circle 结构体表示圆形
type Circle struct {
	Radius float64
}

// Circle实现Shape接口的Area方法
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Circle实现Shape接口的Perimeter方法
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func mainTask14() {
	// 创建Rectangle实例
	rectangle := Rectangle{
		Width:  5.0,
		Height: 3.0,
	}

	// 创建Circle实例
	circle := Circle{
		Radius: 4.0,
	}

	// 通过接口调用方法
	printShapeInfo(rectangle)
	printShapeInfo(circle)

	// 直接调用方法
	fmt.Println("\n直接调用:")
	fmt.Printf("矩形 面积: %.2f, 周长: %.2f\n", rectangle.Area(), rectangle.Perimeter())
	fmt.Printf("圆形 面积: %.2f, 周长: %.2f\n", circle.Area(), circle.Perimeter())
}

// printShapeInfo 使用Shape接口打印形状信息
func printShapeInfo(s Shape) {
	fmt.Printf("形状信息 - 面积: %.2f, 周长: %.2f\n", s.Area(), s.Perimeter())
}

// Person 结构体包含基本信息
type Person struct {
	Name string
	Age  int
}

// Employee 结构体通过组合Person结构体并添加额外字段
type Employee struct {
	Person     // 匿名字段组合Person结构体
	EmployeeID int
}

// PrintInfo 为Employee结构体实现的方法，输出员工信息
func (e Employee) PrintInfo() {
	fmt.Printf("员工ID: %d, 姓名: %s, 年龄: %d\n",
		e.EmployeeID, e.Name, e.Age)
}

// PrintDetailedInfo 带有更多详细信息的打印方法
func (e Employee) PrintDetailedInfo() {
	fmt.Printf("=== 员工详细信息 ===\n")
	fmt.Printf("员工ID: %d\n", e.EmployeeID)
	fmt.Printf("姓名: %s\n", e.Name)
	fmt.Printf("年龄: %d\n", e.Age)
	fmt.Printf("==================\n")
}

func mainTask15() {
	// 创建Employee实例
	employee1 := Employee{
		Person: Person{
			Name: "张三",
			Age:  30,
		},
		EmployeeID: 1001,
	}

	// 另一种创建方式
	employee2 := Employee{}
	employee2.Name = "李四"
	employee2.Age = 25
	employee2.EmployeeID = 1002

	// 第三种创建方式
	employee3 := Employee{
		Person:     Person{"王五", 28},
		EmployeeID: 1003,
	}

	// 调用PrintInfo方法输出员工信息
	employee1.PrintInfo()
	employee2.PrintInfo()
	employee3.PrintInfo()

	// 调用详细信息打印方法
	fmt.Println()
	employee1.PrintDetailedInfo()

	// 直接访问组合字段
	fmt.Printf("直接访问: %s, %d岁, 员工ID: %d\n",
		employee1.Name, employee1.Age, employee1.EmployeeID)
}

// 生产者函数，只写入通道
func producer(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 10; i++ {
		ch <- i
		fmt.Printf("生产: %d\n", i)
	}

	close(ch)
}

// 消费者函数，只读取通道
func consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for num := range ch {
		fmt.Printf("消费: %d\n", num)
	}

	fmt.Println("消费完成")
}

func mainWithUnidirectional() {
	// 创建通道
	ch := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	// 启动生产者和消费者协程
	go producer(ch, &wg)
	go consumer(ch, &wg)

	// 等待所有协程完成
	wg.Wait()

	fmt.Println("程序执行完毕")
}

// producer 生产者协程，向缓冲通道发送100个整数
func producer1(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done() // 协程结束时通知WaitGroup

	fmt.Println("生产者开始工作...")

	// 生产100个整数
	for i := 1; i <= 100; i++ {
		ch <- i        // 发送整数到缓冲通道
		if i%10 == 0 { // 每10个数字显示一次进度
			fmt.Printf("生产者已生产 %d 个数字\n", i)
		}
	}

	close(ch) // 生产完成后关闭通道
	fmt.Println("生产者工作完成")
}

// consumer 消费者协程，从缓冲通道接收整数并打印
func consumer2(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done() // 协程结束时通知WaitGroup

	fmt.Println("消费者开始工作...")
	count := 0

	// 从通道接收数据直到通道关闭
	for num := range ch {
		fmt.Printf("消费者接收到: %d\n", num)
		count++

		if count%10 == 0 { // 每接收10个数字显示一次进度
			fmt.Printf("消费者已消费 %d 个数字\n", count)
		}

		// 模拟处理时间
		time.Sleep(10 * time.Millisecond)
	}

	fmt.Printf("消费者工作完成，总共消费 %d 个数字\n", count)
}

func mainTask17() {
	// 创建一个带缓冲的通道，缓冲区大小为20
	bufferSize := 20
	ch := make(chan int, bufferSize)

	fmt.Printf("创建了缓冲大小为 %d 的通道\n", bufferSize)

	// 创建WaitGroup用于等待协程完成
	var wg sync.WaitGroup
	wg.Add(2) // 添加两个协程任务

	// 启动生产者和消费者协程
	go producer1(ch, &wg)
	go consumer2(ch, &wg)

	// 等待所有协程完成
	wg.Wait()

	fmt.Println("程序执行完毕")
}

func mainTask18() {
	// 创建互斥锁和共享计数器
	var mu sync.Mutex
	counter := 0

	// 创建WaitGroup等待所有协程完成
	var wg sync.WaitGroup

	// 设置协程数量和每个协程的递增次数
	numWorkers := 10
	incrementsPerWorker := 1000
	totalExpected := numWorkers * incrementsPerWorker

	fmt.Printf("启动 %d 个协程，每个协程递增 %d 次\n", numWorkers, incrementsPerWorker)
	fmt.Printf("期望的最终计数器值: %d\n", totalExpected)

	// 记录开始时间
	startTime := time.Now()

	// workerFunc 工作协程函数
	workerFunc := func(id int, increments int) {
		defer wg.Done()

		fmt.Printf("工作协程 #%d 开始执行 %d 次递增操作\n", id, increments)

		for i := 0; i < increments; i++ {
			mu.Lock()   // 加锁
			counter++   // 递增计数器
			mu.Unlock() // 解锁
		}

		fmt.Printf("工作协程 #%d 完成\n", id)
	}

	// 启动10个协程
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1) // 添加协程任务
		go workerFunc(i, incrementsPerWorker)
	}

	// 等待所有协程完成
	wg.Wait()

	// 计算执行时间
	duration := time.Since(startTime)

	// 输出最终结果
	fmt.Printf("\n执行完成!\n")
	fmt.Printf("实际计数器值: %d\n", counter)
	fmt.Printf("期望计数器值: %d\n", totalExpected)
	fmt.Printf("结果正确: %t\n", counter == totalExpected)
	fmt.Printf("执行时间: %v\n", duration)
}

func main1111() {
	// 使用int64类型的计数器
	var counter int64

	// 创建WaitGroup等待所有协程完成
	var wg sync.WaitGroup

	// 设置协程数量和每个协程的递增次数
	numWorkers := 10
	incrementsPerWorker := 1000
	totalExpected := int64(numWorkers * incrementsPerWorker)

	fmt.Printf("启动 %d 个协程，每个协程递增 %d 次\n", numWorkers, incrementsPerWorker)
	fmt.Printf("期望的最终计数器值: %d\n", totalExpected)

	// 记录开始时间
	startTime := time.Now()

	// workerFunc 工作协程函数，使用原子操作递增计数器
	workerFunc := func(id int, increments int) {
		defer wg.Done()

		fmt.Printf("工作协程 #%d 开始执行 %d 次递增操作\n", id, increments)

		for i := 0; i < increments; i++ {
			// 使用原子操作递增计数器
			atomic.AddInt64(&counter, 1)
		}

		fmt.Printf("工作协程 #%d 完成\n", id)
	}

	// 启动10个协程
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1) // 添加协程任务
		go workerFunc(i, incrementsPerWorker)
	}

	// 等待所有协程完成
	wg.Wait()

	// 计算执行时间
	duration := time.Since(startTime)

	// 使用原子操作获取最终值
	actualValue := atomic.LoadInt64(&counter)

	// 输出最终结果
	fmt.Printf("\n执行完成!\n")
	fmt.Printf("实际计数器值: %d\n", actualValue)
	fmt.Printf("期望计数器值: %d\n", totalExpected)
	fmt.Printf("结果正确: %t\n", actualValue == totalExpected)
	fmt.Printf("执行时间: %v\n", duration)
}
