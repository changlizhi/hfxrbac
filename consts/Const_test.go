package consts

import (
	"fmt"
	"log"
	"reflect"
	"testing"
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

type Test struct {
	value string
}
func (test *Test)SetValue(value string) {
	test.value = value
}
func (test *Test)Print()string{
	log.Println(test.value)
	return test.value
}
func TestName(t *testing.T){
	rsm := registerStructMaps{}
	//注册test
	rsm.Register("test",&Test{})
	//获取新的test的interface
	test11,_ := rsm.New("test")
	test22,_ := rsm.New("test")
	//因为 test11 和 test22都是interface{},必须转换为*Test
	test1 := test11.(*Test)
	test2 := test22.(*Test)
	test1.SetValue("aaa")
	test2.SetValue("bbb")
	test1.Print()
	test2.Print()
	r11:=reflect.ValueOf(test11)
	r11ret:=r11.Method(1).Call([]reflect.Value{
		reflect.ValueOf("123"),
	})
	r111:=reflect.ValueOf(test11)
	r111ret:=r111.Method(0).Call([]reflect.Value{
	})
	log.Println("r11ret---",r11ret,r111ret,r11)

	r11t:=reflect.TypeOf(test11)

	log.Println("r11t.Method(0).Name---",r11t.NumMethod(),r11t.String(),r11t.Name(),r11t.Method(0).Type)
}