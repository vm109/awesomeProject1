package my_utils

type Queue struct{
	Queue []interface{}
}

func (q *Queue) Enqueue(value interface{}){
	q.Queue = append(q.Queue, value)
}

func(q *Queue) Dequeue() interface{}{
	value := q.Queue[0]
	q.Queue = q.Queue[1:]
	return value
}
