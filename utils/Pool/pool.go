package Pool

import (
	"container/list"
	"sync"
)

type Pool struct {
	idle    *list.List  // 空闲队列(双向)链表
	actives int         // 总连接数
	mtx     *sync.Mutex // 同步锁
	cond    *sync.Cond  // 用于阻塞/唤醒
}

// NewPool 创建连接池,初始化连接队列,更新队列大小
