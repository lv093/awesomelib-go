package base_algo

/**
队列：先进先出
最基本的操作：
	Enqueue-入队
	Dequeue-出队
 */
type Queue interface {
	Enqueue(data interface{}) bool
	Dequeue() interface{}
}

/**
顺序队列:用数组实现的队列
同样是个循环队列
 */
type ArrayQueue struct {

	items []interface{}
	head  int
	tail  int
}


func (this *ArrayQueue) Enqueue(data interface{}) bool {
	if (this.tail+1)%(len(this.items)) == this.head {
		return false
	}
	next := (this.tail + 1) % (len(this.items))
	this.items[next] = data
	this.tail = next
	return true
}

func (this *ArrayQueue) Dequeue() interface{} {
	if this.head == this.tail {
		return nil
	}
	value := this.items[this.head]
	this.head++
	return value
}

/**
链表队列
 */
type LinkedQueue struct {

}


/**
阻塞队列
数组实现方式
利用双向cahnnel
TODO 使用CAS保证出队和入队的线程安全性，即是实现一种线程安全队列
 */
type BlockArrayQueue struct {
	n int
	items []interface{}
	lockers chan int
	head int
	tail int
}

/**
先写入items，然后写入lockers
 */
func (this *BlockArrayQueue)Enqueue(data interface{}) bool {
	this.lockers <- 1
	next := (this.tail + 1) % (len(this.items))
	this.items[next] = data
	this.tail = next
	return true
}

/**
先获取lockers，然后才能获取items
 */
func (this *BlockArrayQueue)Dequeue() interface{} {
	locker := 0
	ok := false
	for {
		locker, ok = <- this.lockers
		if ok {
			break
		}
	}
	if locker > 0 {
		value := this.items[this.head]
		this.head++
		return value
	}
	return nil
}
