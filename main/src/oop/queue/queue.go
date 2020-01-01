package queue

// 使Queue支持任何类型时
// type Queue []interface{}
type Queue []int

// func (q *Queue) Push(v interface{}) {
func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

// func (q *Queue) Pop() interface{}) {
func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q Queue) IsEmpty() bool {
	return len(q) == 0
}
