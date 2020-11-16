package ml1changliangs

import (
	"fmt"
	"reflect"
)

//定义注册结构map
type registerStructMaps map[string]reflect.Type

//根据name初始化结构
//在这里根据结构的成员注解进行DI注入，这里没有实现，只是简单都初始化
func (rsm registerStructMaps)New(name string) (c interface{},err error){
	if v,ok := rsm[name];ok{
		c = reflect.New(v).Interface()
	} else{
		err = fmt.Errorf("not found %s struct",name)
	}
	return
}
//根据名字注册实例
func (rsm registerStructMaps)Register(name string,c interface{})  {
	rsm[name] = reflect.TypeOf(c).Elem()
}

