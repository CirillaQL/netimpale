package pipeline

import "sync"

/*
	因为Channel在满了的情况下会阻塞,这在单进程的情况下会造成严重的堵塞,因此封装一个独立的Pipeline
	的方法,会动态创建多个Channel顺序Handler连接信息。
*/

type PiPeLine struct {
	mu sync.Mutex
}
