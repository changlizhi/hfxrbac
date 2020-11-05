package testreflect

import (
	"log"
	"reflect"
	"testing"
)

type Obj struct{}

func (o Obj) OnlyObj(str string, str2 string) (int, string) {
	log.Println(str + str2)
	return 123, "1234"
}

func (o Obj) OnlyObj2(str string, str2 string) (int, string) {
	log.Println(str + str2)
	return 123, "1234"
}

func TestReflect(t *testing.T) {

	//retType:=reflect.TypeOf(Obj{})
	retValue := reflect.ValueOf(Obj{})
	//mt:=retType.Method(0)
	mv := retValue.Method(0)
	ret := mv.Call([]reflect.Value{
		reflect.ValueOf("123"),
		reflect.ValueOf("123"),
	})
	log.Println("mt--", ret[0], ret[1], retValue.NumMethod())

}
