package message

var globalQueue Queue

func init() {
	globalQueue = createQueue(100)
}

func createQueue(size int) Queue {
	return Queue{
		Ch:       make(chan interface{}, size),
		Size:     size,
		WorkNums: 10,
	}
}

func SendMessage(msg interface{}) {
	globalQueue.Ch <- msg //阻塞问题 超时
}
