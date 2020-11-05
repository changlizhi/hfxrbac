package ml1changliangs

import (
	"log"
	"runtime"
)

type XiTongs struct{}
type YeWus struct{}
type Biaos struct{}

func (xts XiTongs) YeWus()string{
	p := make([]uintptr,1)
	runtime.Callers(1, p)
	f := runtime.FuncForPC(p[0])
	log.Println("ZiJi---",f.Name(),)

	runtime.Callers(2, p)
	f2 := runtime.FuncForPC(p[0])
	log.Println("DiaoYongZhe---",f2.Name())
	return f.Name()
}

//系统和表怎么关联起来呢？总不能再加前缀吧，表名已经缩写，系统再加就太累赘了。
//按照创建表的思路就是做一个中间表，也就是说系统表也是一个表，表名表也是一个表。仍然是常量表里生出来的一个数据
//只是会约定系统表必须拥有多个表，只有这样才能生成数据，所以一切的生成其实都围绕着常量先生成。

func (xts XiTongs) Biaos()string{
	p := make([]uintptr,1)
	runtime.Callers(1, p)
	f := runtime.FuncForPC(p[0])
	log.Println("ZiJi---",f.Name(),)

	runtime.Callers(2, p)
	f2 := runtime.FuncForPC(p[0])
	log.Println("DiaoYongZhe---",f2.Name())
	return f.Name()
}

