package queue

// 定义队列接口
type Queue interface {
	EnQueue(v interface{}) (ok bool) //入队列
	DeQueue() (v interface{})        //出队列
	IsFull() (ok bool)               //判断队列是否已满
	IsEmpty() (ok bool)              //判断队列是否为空
	Peek() (v interface{})           //获取队头元素
	Reset()                          //重置队列
	Size() int                       // 获取队列大小
	Print()                          //打印队列
}
