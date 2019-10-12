# GO
![go](https://gss2.bdstatic.com/-fo3dSag_xI4khGkpoWK1HF6hhy/baike/w%3D268%3Bg%3D0/sign=555bbdf0dfc451daf6f60bed8ec6355b/3b87e950352ac65c5eb643ddf9f2b21192138ae8.jpg)

* [官方文档（中文版）](http://docscn.studygolang.com/)
* [官方文档（英文版）](http://docs.studygolang.com/)
* [《Go 语言编程》](https://github.com/KeKe-Li/book/blob/master/Go/%E3%80%8AGo%E8%AF%AD%E8%A8%80%E7%BC%96%E7%A8%8B%E3%80%8B%E9%AB%98%E6%B8%85%E5%AE%8C%E6%95%B4%E7%89%88%E7%94%B5%E5%AD%90%E4%B9%A6.pdf)

## 开始GO的洗脑

* 开始总是从“hello world”起步
```
package main
import "fmt"
func main()  {
	fmt.Println("hello")
}
```
Go中的main方法的报名必须是叫main(即使真实项目中所在的包名不是main)</br>
基于由其他语言的基础开始了解GO（略过基础语法）
### 1.面向对象编程(OOP)
首先，Go语言反对函数和操作符重载（overload），
其次，Go语言支持类、类成员方法、类的组合，但反对继承，反对虚函数（virtual function） 和虚函数重载。</br>
再次，Go语言也放弃了构造函数（constructor）和析构函数（destructor）</br>
在放弃了大量的OOP特性后，Go语言送上了一份非常棒的礼物：接口（interface）<br>

#### 1.1为类型添加方法
```
package main
type Integer int
func main() {
	var a Integer = 1
	a.Max(4)
}
func (a Integer) Max(b Integer) bool {
	return a > b
}
```
上面的这个Integer例子如果不使用Go语言的面向对象特性，而使用之前我们介绍的面向</br>
过程方式实现的话，相应的实现细节将如下所示</br>
```
// 面向过程  
func Integer_Max(a Integer, b Integer) bool {     
	return a > b 
}
```

Go语言中的面向对象为直观，也无需支付额外的成本。如果要求对象必须以指针传递， 这有时会是个额外成本，</br>
因为对象有时很小（比如4字节），用指针传递并不划算。 </br>
只有在你**需要修改对象**的时候，才必须用指针。它不是Go语言的约束，而是一种自然约束。 </br>
```
//正确写法
func (a *Integer) Add(b Integer) {
*a += b 
} 
//错误写法
func (a Integer) Add(b Integer) {
a += b 
} 
```
#### 1.2 结构体
Go语言的结构体（struct）和其他语言的类（class）有同等的地位，但Go语言放弃了包括继 承在内的大量面向对象特性，</br>
只保留了组合（composition）这个基础的特性。 </br>
```
type Rect struct {
x, y float64     
width, height float64 }
```

#### 1.3 可见性
Go语言对关键字的增加非常吝啬，其中没有private、protected、public这样的关键 字。要使某个符号对其他包（package）可见</br>
（即可以访问），需要将该符号定义为以大写字母 开头:
```
type User struct {
	Name    string
	Age     int32
	Address string
}
```
#### 1.4 接口
* 其他语言接口例子
在C++、Java和C#中，为了实 现一个接口，你需要从该接口继承
```
interface IFoo {  
void Bar(); 
} 
class Foo implements IFoo { // Java文法  
	// ... 
} 
class Foo : public IFoo { // C++文法     
// ... 
} 
IFoo* foo = new Foo; 
```
* 非侵入式接口 
在Go语言中，一个类只需要实现了接口要求的所有函数，我们就说这个类实现了该接口</br>
```
type Person interface {
	GerName() string
}
func NewMan() *Man {
	return &Man{Name: "aaa"}
}
type Man struct {
	Name string
	Age  int
}
func (mam *Man) GerName() string {
	return mam.Name
}
```

#### 1.5. 泛型（Any 类型）
由于Go语言中任何对象实例都满足空接口 interface{} ,所以 interface{} 看起来像是可</br>
以指向任何对象的 Any 类型。</br>
```
fmt.XXX是个很好例子
func Println(a ...interface{}) (n int, err error) {
	return Fprintln(os.Stdout, a...)
}
```

### 2.并发编程
* **协程**。协程（Coroutine）本质上是一种用户态线程，不需要操作系统来进行抢占式调度，</br>
且在真正的实现中寄存于线程中，因此，系统开销极小，可以有效提高线程的任务并发</br>
性，而避免多线程的缺点。使用协程的优点是编程简单，结构清晰；缺点是需要语言的</br>
支持，如果不支持，则需要用户在程序中自行实现调度器。目前，原生支持协程的语言还很少。</br>

#### 2.1.goroutine使用？
```
func Add(x, y int) {
 z := x + y
 fmt.Println(z)
} 

开启一个协程：
go Add(1, 1) 
```

`真的这么简单么？` </br>
```
package main

import "fmt"

func sayHello()  {
	fmt.Println("hello")
}

func main()  {
	go sayHello()
	fmt.Println("main")
}
```
上面这个程序会输出什么？</br>
```
main
```
问题分析：</br>
Go程序从初始化main package并执行main()函数开始，当main()函数返回时，程序退出，</br>
且程序并不等待其他goroutine（非主goroutine）结束。</br>
* 方法一：
```
package main
import (
	"fmt"
	"sync"
)
var lock sync.WaitGroup
func sayHello()  {
	fmt.Println("hello")
	lock.Done()
}

func main()  {
	lock.Add(1)
	go sayHello()
	lock.Wait()
	fmt.Println("main")
}
```
下一章节做答案讲解

#### 2.2.并发通信

##### 2.2.1 sync
* Mutex: 互斥锁
* RWMutex：读写锁
* Once：执行一次</br>
```
Once 是一个可以被多次调用但是只执行一次，若每次调用Do时传入参数f不同，但是只有第一个才会被执行。</br>
用法：
func (o *Once) Do(f func())

sync.Once的使用场景例如单例模式、系统初始化。
例如并发情况下多次调用channel的close会导致panic，解决这个问题我们可以使用sync.Once来保证close只会被执行一次。
```
[Once用例](https://github.com/linyuanbin/Go/blob/master/demo/syncOnce.go)

* Cond：信号量
* Pool：临时对象池
* Map：自带锁的map
* WaitGroup 并发等待组
```
用法：
func (wg *WaitGroup) Add(delta int)
func (wg *WaitGroup) Done()
func (wg *WaitGroup) Wait()

案例：
var synWait sync.WaitGroup
func TestSync_WaitGroup(t *testing.T) {
	num:=100
	for i:=0;i<num;i++{
		synWait.Add(1)
		go func() {
            ...
			synWait.Done()
		}()
	}
	synWait.Wait()
    ...
}
```

##### 2.2.2 channel
channel是Go语言在语言级别提供的goroutine间的通信方式。我们可以使用channel在两个或</br>
多个goroutine之间传递消息。</br>
```
func Count(ch chan int) {
 ch <- 1
 fmt.Println("Counting")
}
func main() {
 chs := make([]chan int， 10)
 for i := 0; i < 10; i++ {
 chs[i] = make(chan int)
 go Count(chs[i])
 }
 for _, ch := range(chs) {
 <-ch
 }
} 
```
channel是类型相关的。也就是说，一个channel只能传递一种类型的值，这个类型需要在声</br>
明channel时指定。</br>
```
声明一个不带缓存的通道
ch := make(chan int) 
声明一个带缓存的通道
c := make(chan int, 1024) 
```

### 3.错误处理

#### 3.1.error接口
Go语言引入了一个关于错误处理的标准模式，即error接口，该接口的定义如下：
```
type error interface {
 Error() string
} 
```
在Go编程中异常情况都是用error表示，go函数可以多个返回值的特点，一般error作为最后的一个</br>
返回值，用于说明返回值是否可靠或者程序执行是否成功。</br>
```
func Stat(name string) (fi FileInfo, err error) {
 var stat syscall.Stat_t
 err = syscall.Stat(name, &stat)
 if err != nil {
 return nil, &PathError{"stat", name, err}
 }
 return fileInfoFromStat(&stat, name), nil
} 

创建新error:
errors.New("download error")

自定义error:
type myError struct {
	msg string
}

func (err myError) Error() string {
	return err.msg
}
```

* defer关键字
defer一般用于进行连接或者流的关闭工作，相当于Java中finally的作用，

* panic()和recover()
Go语言引入了两个内置函数panic()和recover()以报告和处理运行时错误和程序中的错误场景:
```
panic():
//正常的函数执行流程将立即终止，并报告错误；这个过程称为：错误处理流程
panic(errors.New("download error"))
```

recover()和defer一起使用，相对于catch异常操作</br>
[用例](https://github.com/linyuanbin/Go/tree/master/demo/panicDemo.go)


