package ml1changliangs

import (
	"log"
	"reflect"
	"testing"
)

func TestJg1(t *testing.T){
	rt:=reflect.TypeOf(Jg2Biaos{})
	log.Println(rt.Name(),rt.Method(0).Name)

}
