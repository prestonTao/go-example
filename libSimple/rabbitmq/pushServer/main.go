package main

import (
	"./queue"
)

func main() {
	msgQueue := new(queue.MsgQueue)
	msgQueue.StartUP()
	msgQueue.Send("你好啊")

	worker := new(queue.Worker)
	worker.Connect()
}
