/*
  author='du'
  date='2020/2/18 2:06'
*/
package taskrunner

import "time"

//结构体
type Worker struct {
	ticker *time.Ticker
	runner *Runner
}

//构造函数
func NewWorker(internal time.Duration, r *Runner) *Worker {
	return &Worker{
		ticker: time.NewTicker(internal * time.Second),
		runner: r,
	}
}

func (w *Worker) startWork() {
	for {
		select {
		//这里的"C"就是ticker传过来的channel
		case <-w.ticker.C:
			go w.runner.startAll()
		}
	}
}
