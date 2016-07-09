package main

import (
       "fmt"
       "reflect"
       "unsafe"
       )

type Entry struct {
}

type testObj struct {
    *Entry
    name string
}

func (obj *testObj) GetName() string{
	return obj.name
}

func (obj *testObj) SetName(newname string) {
	obj.name = newname
}

var dataManager map[uint32] interface{}

func Init() {
	dataManager = make(map[uint32] interface{})
	obj := &testObj {Entry :&Entry{},name:"111111111",}
	kind := reflect.ValueOf(obj).Kind()
	fmt.Println("Init================",kind)
	if kind != reflect.Ptr {
		panic("must be pointer")
	}
	
	//dataManager[1] = obj	
}

func GetEntryByID(id uint32) *interface {} {
	proc := dataManager[id]
	if proc != nil {
		ptrproc := (*interface{})(unsafe.Pointer((reflect.ValueOf(proc).Pointer())))
		return ptrproc
	}
	return nil
}

func main() {
     	/*
	dataManager := make(map[uint32] interface{})
	obj := &testObj {name :"1111111",}
	dataManager[1] = obj
	obj2 := &testObj {name:"2222222",}
	dataManager[2] = obj2
	proc := dataManager[2]
	fmt.Println(reflect.ValueOf(dataManager[2]))
	ptrproc := (*testObj)(unsafe.Pointer((reflect.ValueOf(proc).Pointer())))
	fmt.Println(ptrproc.GetName())
	*/
	/*
	ret := reflect.ValueOf(proc).MethodByName("GetName").Call([]reflect.Value{})
	fmt.Println(reflect.TypeOf(ret),ret)
	*/
	
	Init()
	ret := (*testObj)(unsafe.Pointer(GetEntryByID(uint32(1))))
	if ret != nil {
		ret.SetName("new name")
		fmt.Println(ret.GetName())
	
		ret2 := (*testObj)(unsafe.Pointer(GetEntryByID(uint32(1))))
		fmt.Println(ret2.GetName())
	}
	
}

