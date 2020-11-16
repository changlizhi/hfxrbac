package ml1changliangs

import (
	"log"
	"runtime"
)

func (b Jg2Biaos)Jg1ChangLiangs()string{
	p := make([]uintptr,1)
	runtime.Callers(1, p)
	f := runtime.FuncForPC(p[0])
	log.Println("ZiJi---",f.Name(),)

	runtime.Callers(2, p)
	f2 := runtime.FuncForPC(p[0])
	log.Println("DiaoYongZhe---",f2.Name())
	return f.Name()
}
