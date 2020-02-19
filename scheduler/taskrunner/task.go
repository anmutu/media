/*
  author='du'
  date='2020/2/18 2:07'
*/
package taskrunner

import (
	"errors"
	"log"
	"media/scheduler/s_dbops"
	"os"
	"sync"
)

func VideoClearDispatcher(dc dataChan) error {
	res, err := s_dbops.ReadVideoDelRec(3)
	if err != nil {
		return err
	}

	if len(res) == 0 {
		return errors.New("所有任务已完成，当前数据为0")
	}

	//把查到的数据给到dataChan
	for _, id := range res {
		dc <- id
	}
	return nil
}

func VideoClearExecutor(dc dataChan) error {
	errMap := &sync.Map{}
	var err error

forloop:
	for {
		select {
		case vid := <-dc:
			go func(id interface{}) {
				//id.(string)的写法
				if err = deleteVideo(id.(string)); err != nil {
					errMap.Store(id, err)
					return
				}
				if err = s_dbops.DeleteVideoDelRec(id.(string)); err != nil {
					errMap.Store(id, err)
					return
				}
			}(vid)
		default:
			break forloop
		}
	}

	return nil
}

func deleteVideo(vid string) error {
	err := os.Remove(VIDEO_PATH + vid)
	if err != nil && os.IsNotExist(err) {
		log.Printf("删除视频失败:%v", err)
		return err
	}
	return nil
}
