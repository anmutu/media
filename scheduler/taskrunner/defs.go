/*
  author='du'
  date='2020/2/18 2:07'
*/
package taskrunner

const (
	//三个controlChan里的消息
	READY2DISPATCH = "d"
	READY2EXECUTE  = "e"
	CLOSE          = "c"
)

//控制传输的channel
type controlChan chan string

//传输具体数据的channel
type dataChan chan interface{}

//这里就是调度者dispatcher和执行者excutor了
type fn func(dc dataChan) error
