/*
  author='du'
  date='2020/2/18 2:08'
*/
package taskrunner

//runner的结构体
type Runner struct {
	Controller controlChan
	Error      controlChan
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

	defer func() {
		if !r.longLived {
			close(r.Controller)
			close(r.Data)
			close(r.Error)
		}
	}()

	for {
		select {
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
					r.Controller <- READY2EXECUTE
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
	r.Controller <- READY2DISPATCH
	r.startDispatch()
}
