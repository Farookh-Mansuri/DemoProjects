package main

import (
	"fmt"
)

type Queue struct {
	Info []string
}

// q:=queue{  Info :[]string{} , }
func New() *Queue {
	return (&Queue{
		Info: []string{},
	})
}
func (q *Queue) Push(ele string) *Queue {
	q.Info = append(q.Info, ele)
	return q
}
func (q *Queue) Pop() (string, error) {
	ele := q.Info[0]
	q.Info = q.Info[1:]
	return ele, nil
}
func (q *Queue) StackPop() (string, error) {
	ele := q.Info[len(q.Info)-1]
	q.Info = q.Info[:len(q.Info)-1]
	return ele, nil
}
func main() {
	head := New()
	head = head.Push("A")
	head = head.Push("B")
	head = head.Push("C")
	head = head.Push("D")
	head = head.Push("E")
	fmt.Println(head.Pop())
	fmt.Println(head.StackPop())
	fmt.Println(*head)
}
