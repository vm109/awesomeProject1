package my_utils

type Queue struct {
	Queue []interface{}
}

func NewQueue() (q Queue) {
	q = Queue{
		Queue: make([]interface{}, 0),
	}
	return q
}
func (q *Queue) Enqueue(value interface{}) {
	q.Queue = append(q.Queue, value)
}

func (q *Queue) Dequeue() interface{} {
	value := q.Queue[0]
	q.Queue = q.Queue[1:]
	return value
}

func (q *Queue) Peek() interface{} {
	return q.Queue[0]
}
