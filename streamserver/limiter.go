/*
  author='du'
  date='2020/2/17 15:59'
*/
package main

import "log"

/*
  1.这里是流控逻辑的实现。
  2.可以看到用非常简短的代码实现了流控的逻辑。
  3.其得益于golang的channel的机制了。
*/

func main() {

}

//结构体
type ConnLimiter struct {
	conncurrentConn int //当前连接数
	bucket          chan int
}

//构造函数
func NewConnLimiter(cc int) *ConnLimiter {
	return &ConnLimiter{
		conncurrentConn: cc,
		bucket:          make(chan int, cc),
	}
}

//ConnLimiter的函数，当前连接数小于bucket这个channel的数量才return true
func (cl *ConnLimiter) GetConn() bool {
	if len(cl.bucket) >= cl.conncurrentConn {
		log.Printf("达到限制")
		return false
	}
	cl.bucket <- 1
	return true
}

//释放连接
func (cl *ConnLimiter) RealseConn() {
	c := <-cl.bucket
	log.Printf("新连接:%d", c)
}
