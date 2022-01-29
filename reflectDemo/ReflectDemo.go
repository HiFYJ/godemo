package reflectDemo

import (
	"fmt"
	"reflect"
)

func reflectType(x interface{}) {
	v:=reflect.TypeOf(x)
	fmt.Printf("type is %s\n",v)
	fmt.Printf("typename is %v\n",v.Name())
	fmt.Printf("typekind is %v\n",v.Kind())
}

/**
通过反射获取原值
**/
func reflectValue(x interface{}) {
	v:=reflect.ValueOf(x)
	fmt.Printf("type is %s\n",v.Type())
	fmt.Printf("typekind is %v\n",v.Kind())
	switch v.Kind() {
	case reflect.Int8:
		fmt.Printf("原值：%v\n",int64(v.Int()))
		v.Elem().SetInt(101)
		fmt.Printf("修改后值：%v\n",int64(v.Int()))
	case reflect.Float32:
		fmt.Printf("原值：%v\n",float32(v.Float()))
		v.Elem().SetFloat(3.1415926)
		fmt.Printf("修改后值：%v\n",float32(v.Float()))
	}
}

/**
通过反射设置值
**/
func reflectSetValue(x interface{}) {
	v:=reflect.ValueOf(x)
	fmt.Printf("typekind is %v\n",v.Elem().Kind())
	// 反射中使用 Elem()方法获取指针对应的值
	switch v.Elem().Kind() {
	case reflect.Int8:
		fmt.Printf("原值：%v\n",int64(v.Elem().Int()))
		v.Elem().SetInt(101)
		fmt.Printf("修改后值：%v\n",int64(v.Elem().Int()))
	case reflect.Float32:
		fmt.Printf("原值：%v\n",float32(v.Elem().Float()))
		v.Elem().SetFloat(3.1415926)
		fmt.Printf("修改后值：%v\n",float32(v.Elem().Float()))
	}
}

func Input() {
	var a float32 = 3.14
	//reflectType(a)
	//reflectValue(a)
	reflectSetValue(&a)

	var b int8= 100
	//reflectType(b)
	//reflectValue(b)
	reflectSetValue(&b)
	var d = struct {

	}{}
	var e = reflect.TypeOf(d).Kind()
	fmt.Printf("struct type:%v\n",e)

	/*数组、切片、Map、指针等类型的变量，它们的.Name()都是返回空*/
	//var c = make([]int, 3, 5)
	//reflectType(c)
}