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


### 2.





