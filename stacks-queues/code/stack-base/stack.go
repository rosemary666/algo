package stack

// 定义栈接口
type Stack interface {
	Push(v interface{})   //入栈
	Pop() (v interface{}) //出栈
	IsEmpty() (ok bool)   //判断是否为空
	Top() (v interface{}) //获取栈顶元素
	Reset()               //重置栈
	Size() int            // 获取栈大小
	Print()               //打印栈
}
