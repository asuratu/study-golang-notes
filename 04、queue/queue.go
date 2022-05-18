package queue

// Queue
// 该slice仅支持int
// type Queue []int

// Queue
// 该slice支持所有类型
type Queue []interface{}

//func (q *Queue) Push(v ...interface{}) {
//	*q = append(*q, v...)
//}

func (q *Queue) Push(v interface{}) {
	// v.(int) 将interface{}强制转换为int，运行时会出错
	*q = append(*q, v.(int))
}

func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	// head.(int) 将interface{}强制转换为int
	return head.(int)
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
