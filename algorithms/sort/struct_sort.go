package main

import "fmt"

type student struct{
	id int
	name string
	score int
}

func main(){
	//结构体数组存储多为学员信息
	var  arr [3]student = [3]student{
		student{1,"李白",100},
		student{2,"赵四",99},
		student{3,"小白",88}}
	fmt.Println(arr)
	//{1 李白 100}
	//{2 曹操 100}
	//{3 小白 100}


	//作为函数参数使用
	Sort(arr)
	fmt.Println(arr)
}

func Sort(array [3]student){
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-1-i; j++ {

			//根据成绩排序
			if array[j].score > array[j+1].score{   // >升序  <降序
				array[j],array[j+1] = array[j+1],array[j]
			}
		}
	}
}