package Collections

import (
	"fmt"
	"reflect"
)

var (
	TypeNotSupported      = fmt.Errorf(" type is not supported")
	GetValueError         = fmt.Errorf("get value error")
	WhereTypeNotSameError = fmt.Errorf("where type not same")
)

type Collection struct {
	InnerSlice          any
	ReflectValueOfInner reflect.Value
	Error               error
}

func NewCollection(array any) *Collection {
	var err error = nil
	//todo 检查入参是否是array
	return &Collection{
		InnerSlice:          array,
		ReflectValueOfInner: reflect.ValueOf(array),
		Error:               err,
	}
}

func (c *Collection) Get(dest any) {
	if c.Error != nil || c.ReflectValueOfInner.Len() <= 0 {
		return
	}
	reflectValueOfDest := reflect.ValueOf(dest)
	if reflectValueOfDest.Kind() != reflect.Ptr || reflectValueOfDest.Elem().Kind() != reflect.Slice {
		c.Error = GetValueError
		return
	}
	reflectValueOfDest = reflectValueOfDest.Elem()
	//if c.ReflectValueOfInner.Type() != reflect.ValueOf(reflectValueOfDest).Type() {
	//	c.Error = GetValueError
	//	return
	//}

	// 确保b的容量足够
	if reflectValueOfDest.Cap() < c.ReflectValueOfInner.Len() {
		reflectValueOfDest.Set(reflect.MakeSlice(reflectValueOfDest.Type(), 0, c.ReflectValueOfInner.Len()))
	}
	reflectValueOfDest.SetLen(0)

	// 逐个复制元素
	for i := 0; i < c.ReflectValueOfInner.Len(); i++ {
		reflectValueOfDest = reflect.Append(reflectValueOfDest, c.ReflectValueOfInner.Index(i))
	}

	//设置回去：因为append过程中reflectValueOfDest有更新，不会自动映射到原来的切片
	reflect.ValueOf(dest).Elem().Set(reflectValueOfDest)
}

// transfer
//
//	@Description: 对保存的元素逐个进行transFunc
//	@receiver c
//	@param transFunc
//	@return *Collection
func (c *Collection) transfer(transFunc func(in any) (out any)) *Collection {
	if c.Error != nil || c.ReflectValueOfInner.Len() <= 0 {
		return c
	}
	oldReflectValueOfInner := c.ReflectValueOfInner
	//oldInnerSlice := c.InnerSlice
	//outType, err := GetReturnType(transFunc) 失败的方案1:因为这个func就是返回的any，所以虽然用了反射，但是依然获取不到类型，什么都打印不出来
	//
	//if err != nil {
	//	c.Error = err
	//	return c
	//}
	// todo 如果出现传入参数和inner参数不一致的情况，会发生runtime panic？或许还是需要尝试下怎么在真正transfer之前揭开in到底是一个什么type

	outType := reflect.TypeOf(transFunc(oldReflectValueOfInner.Index(0).Interface()))
	fmt.Printf("transfer Output Type:%s\n", outType.Name())

	c.ReflectValueOfInner = reflect.MakeSlice(reflect.SliceOf(outType), oldReflectValueOfInner.Len(), oldReflectValueOfInner.Len())
	//c.InnerSlice = make([]any, reflect.ValueOf(c.InnerSlice).Len(), reflect.ValueOf(c.InnerSlice).Len()) //这里是不是应该将any转成out 的类型
	//newReflectValue := reflect.ValueOf(c.InnerSlice)
	for i := 0; i < c.ReflectValueOfInner.Len(); i++ {
		newElem := transFunc(oldReflectValueOfInner.Index(i).Interface())
		valueOf := reflect.ValueOf(newElem)
		//fmt.Println(reflect.TypeOf(newElem).Name())
		c.ReflectValueOfInner.Index(i).Set(valueOf)
	}
	return c
}

// where
//
//	@Description: 只保留符合where判断条件的元素
//	@receiver c
//	@param transFunc
//	@return *Collection
func (c *Collection) where(whereFunc func(in any) bool) *Collection {
	if c.Error != nil || c.ReflectValueOfInner.Len() <= 0 {
		return c
	}
	// todo ： add 输入类型检查以避免运行时错误
	//InnerType := c.ReflectValueOfInner.Index(0).Type()
	//fmt.Printf("where InnerType :%s\n", InnerType.Name())
	//if InnerType != reflect.TypeOf(in) {
	//	c.Error = TypeNotSupported
	//	return c
	//}
	oldReflectValueOfInner := c.ReflectValueOfInner
	c.ReflectValueOfInner = reflect.MakeSlice(c.ReflectValueOfInner.Type(), 0, oldReflectValueOfInner.Len())

	for i := 0; i < oldReflectValueOfInner.Len(); i++ {
		if whereFunc(oldReflectValueOfInner.Index(i).Interface()) {
			c.ReflectValueOfInner = reflect.Append(c.ReflectValueOfInner, oldReflectValueOfInner.Index(i))
		}
	}
	return c
}

// GetReturnType 获取函数返回值类型的函数
func GetReturnType(fn func(in any) (out any)) (reflect.Type, error) {
	// 获取函数的反射类型
	funcType := reflect.TypeOf(fn)

	// 检查是否是函数类型
	if funcType.Kind() != reflect.Func {
		fmt.Println("Input is not a function")
		return nil, fmt.Errorf("input is not a function")
	}

	// 获取函数的第一个返回值类型
	// 函数可以有多个返回值，所以我们用 Out(0) 获取第一个返回值
	return0Name := funcType.Out(0).Name()
	fmt.Println(return0Name)
	return funcType.Out(0), nil
}
