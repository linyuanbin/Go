package main

type Integer int

func (a Integer) Less(b Integer) bool {
	return a < b
}
func (a *Integer) Add(b Integer) {
	*a += b
}

type LessAdder interface {
	Less(b Integer) bool
	Add(b Integer)
}

func main() {
	var a Integer = 1
	var b LessAdder = &a //  1 正确
	var c LessAdder = a  //  2 错误
}

/**
Go语言可以根据下面的函数:
func (a Integer) Less(b Integer) bool 自动生成一个新的 Less() 方法:
func (a *Integer) Less(b Integer) bool {
return (*a).Less(b)
}

这样,类型 *Integer 就既存在 Less() 方法,也存在 Add() 方法,满足 LessAdder 接口。而
从另一方面来说,根据
func (a *Integer) Add(b Integer)
这个函数无法自动生成以下这个成员方法:
func (a Integer) Add(b Integer) {
(&a).Add(b)
}

因为 (&a).Add() 改变的只是函数参数 a ,对外部实际要操作的对象并无影响,这不符合用
户的预期。所以,Go语言不会自动为其生成该函数。因此,类型 Integer 只存在 Less() 方法,
缺少 Add() 方法,不满足 LessAdder 接口,故此上面的语句(2)不能赋值。
 */
