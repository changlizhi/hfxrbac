package ml1jiegous

import (
	"log"
	"runtime"
)

type Jg1YongHu struct{}
type Jg1YhBiaos struct{}

func Yh()string{
	p := make([]uintptr,1)
	runtime.Callers(1, p)
	f := runtime.FuncForPC(p[0])
	log.Println("ZiJi---",f.Name(),)

	runtime.Callers(2, p)
	f2 := runtime.FuncForPC(p[0])
	log.Println("DiaoYongZhe---",f2.Name())
	return f.Name()
}

func (jg Jg1YongHu)JianXie()string{
	p := make([]uintptr,1)
	runtime.Callers(1, p)
	f := runtime.FuncForPC(p[0])
	log.Println("ZiJi---",f.Name(),)

	runtime.Callers(2, p)
	f2 := runtime.FuncForPC(p[0])
	log.Println("DiaoYongZhe---",f2.Name())
	return Yh()
}
