package utils

var ErrorQueue chan error

func NewQueue()  {
	ErrorQueue = make(chan error) //异常信号
}

