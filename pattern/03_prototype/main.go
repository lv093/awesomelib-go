package main

import (
	"fmt"
)

/**
prototype 原型模式
创建型模式
用原型实例指定创建对象的种类，并且通过拷贝这些原型创建新的对象
将一个对象作为原型，对其进行复制、克隆，产生一个和原对象类似的新对象

原型模式适用场景：
	1.结构体初始化消耗大
	2.初始化对象时需要非常繁琐的过程（数据准备、访问权限等）
注意：注意克隆过程中深拷贝、浅拷贝问题。
 */
type Resume struct {
	id int
}

func (r *Resume) clone() *Resume {
	if r == nil {
		return nil
	}
	new_obj := (*r)
	return &new_obj
}

func NewResume() *Resume {
	return &Resume{}
}

func main() {
	old1 := NewResume()
	fmt.Printf("old pointer:%p\n", old1)

	new1 := old1.clone()
	fmt.Printf("new pointer:%p\n", new1)

	old2 := NewResume()
	fmt.Printf("old2 pointer:%p\n", old2)
}