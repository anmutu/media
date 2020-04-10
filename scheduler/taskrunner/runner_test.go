/*
  author='du'
  date='2020/2/18 3:34'
*/
package taskrunner

import (
	"errors"
	"log"
	"testing"
	"time"
)

func TestRunner(t *testing.T) {
	dispatcher := func(dc dataChan) error {
		for i := 0; i < 50; i++ {
			dc <- i
			log.Printf("dispatcher发送了数据:%v", i)
		}
		return nil
	}

	executor := func(dc dataChan) error {
		// for d:=range dc 这种不是异步的，是顺序的，这里不用。
	forloop:
		for {
			select {
			case d := <-dc:
				log.Printf("exector接收到数据:%v", d)
			default:
				break forloop
			}
		}

		return errors.New("executor执行完成")
	}

	runner := NewRunner(50, false, dispatcher, executor)
	go runner.startAll()
	time.Sleep(5 * time.Second)
}
