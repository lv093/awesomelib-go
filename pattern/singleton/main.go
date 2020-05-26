/*
 Singleton 单件：
        保证一个类仅有一个实例，并提供一个访问它的全局访问点

1.单纯使用is nil判断存在并发问题
2.使用mutex 锁机制，非原子操作
2.创建实例时，用同步库sync.Once来保证全局只有一个对象实例。
*/
package main

import (
	"fmt"
	"sync"
)

type singleton struct {
	data int
}

// 定义一个包级别的private实例变量
var instance *singleton

//1.使用mutex锁
var lock sync.Mutex
// 2.同步Once,保证每次调用时，只有第一次生效
var once sync.Once

//1.使用mutex锁实现单例
func GetInstance() *singleton {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		instance = &singleton{12}
	}
	return instance
}

// 2.获取实例对象函数
func GetSingleton() *singleton {
	once.Do(func() {
		instance = &singleton{12}
	})
	fmt.Println("instance info and pointer:", instance, &instance)
	return instance
}


func main() {
	GetSingleton()
}