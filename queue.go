package message

type Queue struct {
	Ch       chan interface{} `json:"ch"`   //存储消息
	Size     int              `json:"size"` //队列大小
	WorkNums int              `json:"work_nums"`
}
