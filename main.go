// 声明包  effective go中的代码
package main

//引入包声明
import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"sync"
	"time"
)

var (
	countLock   sync.Mutex
	inputCount  uint32
	outputCount uint32
	errorCount  uint32
)

// 函数声明
func printInConsole(s string) {
	fmt.Println(s)
}

// 全局变量声明
var str string = "hello world"

type T struct {
	name  string // name of the object
	value int    // its value
}

func returnNumber(z int) int {
	var x int
	var y int
	x = 1
	y = 2
	if x >= 2 {
		return y
	}
	return z
}

func returnErr() error {
	var file *os.File
	if err := file.Chmod(0664); err != nil {
		log.Print(err)
		return err
	}
	return nil
}

func unhex(c byte) byte {
	switch {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	}
	return 0
}

func shouldEscape(c byte) bool {
	switch c {
	case ' ', '?', '&', '=', '#', '+', '%':
		return true
	}
	return false
}

func Compare(a, b []byte) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		switch {
		case a[i] > b[i]:
			return 1
		case a[i] < b[i]:
			return -1
		}
	}
	switch {
	case len(a) > len(b):
		return 1
	case len(a) < len(b):
		return -1
	}
	return 0
}

func nextInt(b []byte, i int) (int, int) {
	for ; i < len(b); i++ {
	}
	var x int
	x = 0
	for ; i < len(b); i++ {
		x = x*10 + int(b[i]) - '0'
	}
	return x, i
}

func readFull(r io.Reader, buf []byte) (n int, err error) {
	for len(buf) > 0 && err == nil {
		var nr int
		nr, err = r.Read(buf)
		n += nr
		buf = buf[nr:]
	}
	return
}

// defer函数：Go 的 defer 语句用于预设一个函数调用（即推迟执行函数）， 该函数会在执行 defer 的函数返回之前立即执行。它显得非比寻常， 但却是处理一些事情的有效方式，例如无论以何种路径返回，都必须释放资源的函数。 典型的例子就是解锁互斥和关闭文件。

func Contents(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close() //f.Close will run when we're finished.
	var result []byte
	buf := make([]byte, 100)
	for {
		n, err := f.Read(buf[0:])
		result = append(result, buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}
	}
	return string(result), nil
}

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	defer un(trace("a"))
	fmt.Println("in a")
}

func c() {
	defer un(trace("b"))
	fmt.Println("in b")
	a()
}

//推迟诸如 Close 之类的函数调用有两点好处：第一， 它能确保你不会忘记关闭文件。如果你以后又为该函数添加了新的返回路径时， 这种情况往往就会发生。第二，它意味着 “关闭” 离 “打开” 很近， 这总比将它放在函数结尾处要清晰明了。

/*
Go 提供了两种分配原语，即内建函数 new 和 make。 它们所做的事情不同，所应用的类型也不同。 new。这是个用来分配内存的内建函数， 但与其它语言中的同名函数不同，它不会初始化内存，只会将内存置零。 也就是说，new(T) 会为类型为 T 的新项分配已置零的内存空间， 并返回它的地址，也就是一个类型为 *T 的值。用 Go 的术语来说，它返回一个指针， 该指针指向新分配的，类型为 T 的零值。
内建函数 make(T, args) 的目的不同于 new(T)。它只用于创建切片、映射和信道，并返回类型为 T（而非 *T）的一个已初始化 （而非置零）的值。 出现这种差异的原因在于，这三种类型本质上为引用数据类型，它们在使用前必须初始化。 例如，切片是一个具有三项内容的描述符，包含一个指向（数组内部）数据的指针、长度以及容量， 在这三项被初始化之前，该切片为 nil。对于切片、映射和信道，make 用于初始化其内部的数据结构并准备好将要使用的值。
请记住，make 只适用于映射、切片和信道且不返回指针。若要获得明确的指针， 请使用 new 分配内存或显式地获取一个变量的地址。
*/
type SynceBuffer struct {
	lock   sync.Mutex
	buffer bytes.Buffer
}

func Sum(a *[3]float64) (sum float64) {
	for _, v := range *a {
		sum += v
	}
	return
}

func append(slice, data []byte) []byte {
	l := len(slice)
	if l+len(data) > cap(slice) {
		newSlice := make([]byte, (len(slice)+len(data))*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : l+len(data)]
	for i, c := range data {
		slice[l+i] = c
	}
	return slice
}

var timeZone = map[string]int{
	"UTC": 0 * 60 * 60,
	"EST": -5 * 60 * 60,
	"CST": -6 * 60 * 60,
	"MST": -7 * 60 * 60,
	"PST": -8 * 60 * 60,
}

func offset(tz string) int {
	if seconds, ok := timeZone[tz]; ok {
		return seconds
	}
	log.Println("unknown time zone", tz)
	return 0
}

type ByteSize float64

const (
	// 通过赋予空白标识符来忽略第一个值
	_           = iota // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func (b ByteSize) String() string {
	switch {
	case b >= YB:
		return fmt.Sprintf("%.2fYB", b/YB)
	case b >= ZB:
		return fmt.Sprintf("%.2fZB", b/ZB)
	case b >= EB:
		return fmt.Sprintf("%.2fEB", b/EB)
	case b >= PB:
		return fmt.Sprintf("%.2fPB", b/PB)
	case b >= TB:
		return fmt.Sprintf("%.2fTB", b/TB)
	case b >= GB:
		return fmt.Sprintf("%.2fGB", b/GB)
	case b >= MB:
		return fmt.Sprintf("%.2fMB", b/MB)
	case b >= KB:
		return fmt.Sprintf("%.2fKB", b/KB)
	}
	return fmt.Sprintf("%.2fB", b)
}

var (
	home   = os.Getenv("HOME")
	user   = os.Getenv("USER")
	gopath = os.Getenv("GOPATH")
)

func init() {
	if user == "" {
		log.Fatal("$USER not set")
	}
	if home == "" {
		home = "/home/" + user
	}
	if gopath == "" {
		gopath = home + "/go"
	}
	// gopath 可通过命令行中的 --gopath 标记覆盖掉。
	flag.StringVar(&gopath, "gopath", gopath, "override default GOPATH")
}

type Sequence []int

func (s Sequence) Len() int {
	//TODO implement me
	panic("implement me")
	return len(s)
}

func (s Sequence) Less(i, j int) bool {
	//TODO implement me
	panic("implement me")
	return s[i] < s[j]

}

func (s Sequence) Swap(i, j int) {
	//TODO implement me
	panic("implement me")
	s[i], s[j] = s[j], s[i]
}

func (s Sequence) String() string {
	sort.Sort(s)
	str := "["
	for i, elem := range s {
		if i > 0 {
			str += " "
		}
		str += fmt.Sprint(elem)
	}
	return str + "]"
}
func (s Sequence) String1() string {
	sort.Sort(s)
	return fmt.Sprint([]int(s))
}

func (s Sequence) String2() string {
	sort.IntSlice(s).Sort()
	return fmt.Sprint([]int(s))
}

func announce(message string, delay time.Duration) {
	go func() {
		time.Sleep(delay)
		fmt.Println(message)
	}()
}

var maxOutStanding = 100
var sem = make(chan int, maxOutStanding)

func handle() {
	sem <- 1
	<-sem
}

// Bug出现在Go的 for 循环中，该循环变量在每次迭代时会被重用，因此 req 变量会在所有的Go程间共享，这不是我们想要的。我们需要确保 req 对于每个Go程来说都是唯一的。有一种方法能够做到，就是将 req 的值作为实参传入到该Go程的闭包中：
func Serve(queue chan *http.Request) {
	for req := range queue {
		sem <- 1
		go func(req *http.Request) {
			//process(req)
			<-sem
		}(req)
	}
}

// 另一种管理资源的好方法就是启动固定数量的 handle Go程，一起从请求信道中读取数据。Go程的数量限制了同时调用 process 的数量。Serve 同样会接收一个通知退出的信道， 在启动所有Go程后，它将阻塞并暂停从信道中接收消息。
func handle1(queue chan *http.Request) {
	for r := range queue {
		fmt.Println(r)
		//process(r)
	}
}

func Serve3(clientRequests chan *http.Request, quit chan bool) {
	// 启动处理程序
	const MaxOutstanding = 100
	for i := 0; i < MaxOutstanding; i++ {
		go handle1(clientRequests)
	}
	<-quit // 等待通知退出。
}

// 信道可以作为一个值来传递
type Request struct {
	args       []int
	f          func([]int) int
	resultChan chan int
}

func sum1(a []int) (s int) {
	for _, v := range a {
		s += v
	}
	return
}

type Vector []float64

// 并行化
// 将此操应用至 v[i], v[i+1] ... 直到 v[n-1]
func (v Vector) DoSome(i, n int, u Vector, c chan int) {
	for ; i < n; i++ {
		//v[i] += u.Op(v[i])
	}
	c <- 1 // 发信号表示这一块计算完成。
}

const NCPU = 4 // CPU核心数

func (v Vector) DoAll(u Vector) {
	c := make(chan int, NCPU) // 缓冲区是可选的，但明显用上更好
	for i := 0; i < NCPU; i++ {
		go v.DoSome(i*len(v)/NCPU, (i+1)*len(v)/NCPU, u, c)
	}
	// 排空信道。
	for i := 0; i < NCPU; i++ {
		<-c // 等待任务完成
	}
	// 一切完成。
}

/**
目前Go运行时的实现默认并不会并行执行代码，它只为用户层代码提供单一的处理核心。 任意数量的Go程都可能在系统调用中被阻塞，而在任意时刻默认只有一个会执行用户层代码。 它应当变得更智能，而且它将来肯定会变得更智能。但现在，若你希望CPU并行执行， 就必须告诉运行时你希望同时有多少Go程能执行代码。有两种途径可意识形态，要么 在运行你的工作时将 GOMAXPROCS 环境变量设为你要使用的核心数， 要么导入 runtime 包并调用 runtime.GOMAXPROCS(NCPU)。 runtime.NumCPU() 的值可能很有用，它会返回当前机器的逻辑CPU核心数。 当然，随着调度算法和运行时的改进，将来会不再需要这种方法。
注意不要混淆并发和并行的概念：并发是用可独立执行的组件构造程序的方法， 而并行则是为了效率在多CPU上平行地进行计算。尽管Go的并发特性能够让某些问题更易构造成并行计算， 但Go仍然是种并发而非并行的语言，且Go的模型并不适合所有的并行问题。
*/

// 可能泄露的缓冲区
var freeList = make(chan *bytes.Buffer, 100)
var serverChan = make(chan *bytes.Buffer)

func client() {
	for {
		var b *bytes.Buffer
		// 若缓冲区可用就用它，不可用就分配个新的。
		select {
		case b = <-freeList:
			// 获取一个，不做别的。
		default:
			// 非空闲，因此分配一个新的。
			b = new(bytes.Buffer)
		}
		//load(b)              // 从网络中读取下一条消息。
		serverChan <- b // 发送至服务器。
	}
}

func server() {
	for {
		b := <-serverChan // 等待工作。
		//process(b)
		// 若缓冲区有空间就重用它。
		select {
		case freeList <- b:
			// 将缓冲区放大空闲列表中，不做别的。
		default:
			// 空闲列表已满，保持就好。
		}
	}
}

/**
panic函数 运行错误时终止的函数
当 panic 被调用后（包括不明确的运行时错误，例如切片检索越界或类型断言失败）， 程序将立刻终止当前函数的执行，并开始回溯Go程的栈，运行任何被推迟的函数。 若回溯到达Go程栈的顶端，程序就会终止。不过我们可以用内建的 recover 函数来重新或来取回Go程的控制权限并使其恢复正常执行。

调用 recover 将停止回溯过程，并返回传入 panic 的实参。 由于在回溯时只有被推迟函数中的代码在运行，因此 recover 只能在被推迟的函数中才有效。

recover 的一个应用就是在服务器中终止失败的Go程而无需杀死其它正在执行的Go程。
相当于吃异常
*/
// Error 是解析错误的类型，它满足 error 接口。
type Error string

func (e Error) Error() string {
	return string(e)
}

func main() {
	var x int
	var y int
	var z = x<<8 + y<<16
	fmt.Println(z)

	var sum int
	sum = 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	var m = make(map[string]int)
	var n = make(map[string]int)
	for key, value := range m {
		n[key] = value
	}
	sum = 0
	for _, value := range m {
		sum += value
	}
	//printInConsole(str)
	var a [9]int
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	var b []byte
	for i := 0; i < len(b); {
		x, i = nextInt(b, i)
		fmt.Println(x)
	}
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
	//被推迟的函数按照后进先出（LIFO）的顺序执行，因此以上代码在函数返回时会打印 4 3 2 1 0。一个更具实际意义的例子是通过一种简单的方法， 用程序来跟踪函数的执行。
	c()

	var _ = make([]int, 10, 100)
	/**
	会分配一个具有 100 个 int 的数组空间，接着创建一个长度为 10， 容量为 100 并指向该数组中前 10 个元素的切片结构。（生成切片时，其容量可以省略，更多信息见切片一节。） 与此相反，new([]int) 会返回一个指向新分配的，已置零的切片结构， 即一个指向 nil 切片值的指针。
	*/

	/**
	go语言和c语言中数组的区别：
	数组是值。将一个数组赋予另一个数组会复制其所有元素。
	特别地，若将某个数组传入某个函数，它将接收到该数组的一份副本而非指针。
	数组的大小是其类型的一部分。类型 [10]int 和 [20]int 是不同的。
	*/
	array := [...]float64{7.0, 8.5, 9.1}
	var xx float64 = Sum(&array) // Note the explicit address-of operator
	fmt.Println(xx)

	/**
	切片
	切片通过对数组进行封装，为数据序列提供了更通用、强大而方便的接口。 除了矩阵变换这类需要明确维度的情况外，Go 中的大部分数组编程都是通过切片来完成的。
	切片保存了对底层数组的引用，若你将某个切片赋予另一个切片，它们会引用同一个数组。 若某个函数将一个切片作为参数传入，则它对该切片元素的修改对调用者而言同样可见， 这可以理解为传递了底层数组的指针。因此，Read 函数可接受一个切片实参 而非一个指针和一个计数；切片的长度决定了可读取数据的上限。
	只要切片不超出底层数组的限制，它的长度就是可变的，只需将它赋予其自身的切片即可。 切片的容量可通过内建函数 cap 获得，它将给出该切片可取得的最大长度。 以下是将数据追加到切片的函数。若数据超出其容量，则会重新分配该切片。返回值即为所得的切片。 该函数中所使用的 len 和 cap 在应用于 nil 切片时是合法的，它会返回 0.
	最终我们必须返回切片，因为尽管 Append 可修改 slice 的元素，但切片自身（其运行时数据结构包含指针、长度和容量）是通过值传递的。
	向切片追加东西的想法非常有用，因此有专门的内建函数 append。
	有时必须分配一个二维切片，例如在处理像素的扫描行时，这种情况就会发生。 我们有两种方式来达到这个目的。一种就是独立地分配每一个切片；而另一种就是只分配一个数组， 将各个切片都指向它。采用哪种方式取决于你的应用。若切片会增长或收缩， 就应该通过独立分配来避免覆盖下一行；若不会，用单次分配来构造对象会更加高效。
	*/

	type Transform [3][3]float64
	type LinesOfText [][]byte

	text := LinesOfText{
		[]byte("hello world"),
		[]byte("hello world"),
		[]byte("hello world"),
	}
	for _, line := range text {
		fmt.Println(line)
	}

	YSize := 8
	Xsize := 8
	var picture = make([][]uint8, YSize) // 每 y 个单元一行。

	for i := range picture {
		picture[i] = make([]uint8, Xsize)
	}

	//Map

	for name, offset := range timeZone {
		fmt.Printf("%s\t%d\n", name, offset)
	}

	var offset = timeZone["EST"]
	fmt.Println(offset)

	delete(timeZone, "PDT")

	/**
	常量
	Go 中的常量就是不变量。它们在编译时创建，即便它们可能是函数中定义的局部变量。 常量只能是数字、字符（符文）、字符串或布尔值。由于编译时的限制， 定义它们的表达式必须也是可被编译器求值的常量表达式。
	*/

	/**
	init函数
	每个源文件都可以通过定义自己的无参数 init 函数来设置一些必要的状态。 （其实每个文件都可以拥有多个 init 函数。）而它的结束就意味着初始化结束： 只有该包中的所有变量声明都通过它们的初始化器求值后 init 才会被调用， 而那些 init 只有在所有已导入的包都被初始化后才会被求值。除了那些不能被表示成声明的初始化外，init 函数还常被用在程序真正开始执行前，检验或校正程序的状态。
	*/

	/**
	指针 vs 值
	指针或值为接收者的区别在于：值方法可通过指针和值调用， 而指针方法只能通过指针来调用。之所以会有这条规则是因为指针方法可以修改接收者；通过值调用它们会导致方法接收到该值的副本， 因此任何修改都将被丢弃，因此该语言不允许这种错误。不过有个方便的例外：若该值是可寻址的， 那么该语言就会自动插入取址操作符来对付一般的通过值调用的指针方法。
	*/

	/**
	并发
	在并发编程中，为实现对共享变量的正确访问需要精确的控制，这在多数环境下都很困难。 Go语言另辟蹊径，它将共享的值通过信道传递，实际上，多个独立执行的线程从不会主动共享。 在任意给定的时间点，只有一个Go程能够访问该值。数据竞争从设计上就被杜绝了。
	不要通过共享内存来通信，而应通过通信来共享内存。
	Goroutines（go程、协程）
	 Go程具有简单的模型：它是与其它Go程并发运行在同一地址空间的函数。
	在Go中，函数字面都是闭包：其实现在保证了函数内引用变量的生命周期与函数的活动时间相同。
	信道
	信道（channel）是 Go 程与 Go 程之间进行通信的机制。
	在Go中，函数字面都是闭包：其实现在保证了函数内引用变量的生命周期与函数的活动时间相同。
	接收者在收到数据前会一直阻塞。若信道是不带缓冲的，那么在接收者收到值前， 发送者会一直阻塞；若信道是带缓冲的，则发送者仅在值被复制到缓冲区前阻塞； 若缓冲区已满，发送者会一直等待直到某个接收者取出一个值为止。
	带缓冲的信道可被用作信号量，例如限制吞吐量。在此例中，进入的请求会被传递给 handle，它从信道中接收值，处理请求后将值发回该信道中，以便让该 “信号量”准备迎接下一次请求。信道缓冲区的容量决定了同时调用 process 的数量上限，因此我们在初始化时首先要填充至它的容量上限。
	*/

	// 创建请求并发送
	clientRequests := make(chan *Request, 100) // 假设信道已定义
	request := &Request{[]int{3, 4, 5}, sum1, make(chan int)}
	clientRequests <- request
	fmt.Printf("answer: %d\n", <-request.resultChan)

}
