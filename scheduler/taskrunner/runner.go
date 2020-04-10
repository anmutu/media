/*
  author='du'
  date='2020/2/18 2:08'
*/
package taskrunner

//runner的结构体
type Runner struct {
	Controller controlChan
	Error      controlChan //这里实际上就是返回的close这个东西
	Data       dataChan
	dataSize   int
	longLived  bool //是否是长期存活的一个runner,如果是则不回收资源
	Dispatcher fn
	Executor   fn
}

//runner的构造函数
func NewRunner(size int, longlived bool, d fn, e fn) *Runner {
	return &Runner{
		Controller: make(chan string, 1), //这里要使用非阻塞带buffer的channel
		Error:      make(chan string, 1),
		Data:       make(chan interface{}, size),
		longLived:  longlived,
		dataSize:   size,
		Dispatcher: d,
		Executor:   e,
	}
}

//runner开始调度的函数
func (r *Runner) startDispatch() {
	//如果不是长驻则关闭
	defer func() {
		if !r.longLived {
			close(r.Controller)
			close(r.Data)
			close(r.Error)
		}
	}() //有这里的括号就表明是即时调用。

	for {
		select {
		//当这个controller里信息写进来的时候就走到这个case里去
		case c := <-r.Controller:
			if c == READY2DISPATCH {
				err := r.Dispatcher(r.Data)
				if err != nil {
					r.Error <- CLOSE
				} else {
					r.Controller <- READY2EXECUTE
				}
			}
			if c == READY2EXECUTE {
				err := r.Executor(r.Data)
				if err != nil {
					r.Error <- CLOSE
				} else {
					r.Controller <- READY2DISPATCH
				}
			}
		case e := <-r.Error:
			if e == CLOSE {
				return
			}
		}
	}

}

//开始
func (r *Runner) startAll() {
	//这里要先给个READY2DISPATCH，不然会一直卡在这里。
	r.Controller <- READY2DISPATCH
	r.startDispatch()
}
