/*
	Pool是客户端与服务端之间的连接池，用来维护客户端与服务端之间的信息
	交流，并能够确保客户端与服务端之间的连接，在连接断开时能够及时重连，
	并能够对其进行监控。
*/
package pool

import (
	"context"
	"fmt"
	"net"
	Connect "netimpale/pkg/connection"
	"netimpale/utils/log"
	"sync"
	"time"

	uuid "github.com/satori/go.uuid"
)

var LOG = log.LOG

// Pool 连接池
type Pool struct {
	connectPool []Connect.Conn //连接池，数据结构为Slice
	size        uint8          //连接池指定大小
	actives     uint8          //当前的连接数
	mutex       *sync.Mutex    //同步锁
	cond        *sync.Cond     //阻塞唤醒
	msgChan     chan []byte    //消息队列（初步）
}

// NewClientPool 在客户端侧初始化连接，主动向服务端发送连接请求
func NewClientPool(size uint8) (p *Pool, err error) {
	var failedTime uint8
	poolSlice := make([]Connect.Conn, size)
	for i := 0; i < int(size); i++ {
		conn, err := Connect.NewConn("127.0.0.1:8080")
		if err == nil {
			if failedTime != 0 {
				failedTime = 0
			}
			poolSlice = append(poolSlice, *conn)
		} else {
			if failedTime == 5 {
				LOG.Errorf("Connect to Server failed.")
				break
			}
			failedTime++
			i--
			LOG.Infof("Can't connect to server. Error: %v,  try to reconnect time: %d", err, failedTime)
			time.Sleep(1 * time.Second)
		}
	}
	mutex := new(sync.Mutex)
	cond := sync.NewCond(mutex)
	msgChan := make(chan []byte)
	p = &Pool{
		poolSlice,
		size,
		uint8(len(poolSlice)),
		mutex,
		cond,
		msgChan,
	}
	return p, nil
}

// NewServerPool 在服务端侧初始化连接，等待从客户端发来的连接请求
func NewServerPool(size uint8) (p *Pool, err error) {
	poolSlice := make([]Connect.Conn, size)
	_size := 0
	listener, err := net.Listen("tcp", "localhost:50000")
	if err != nil {
		fmt.Println("Error listening", err.Error())
		return //终止程序
	}
	mutex := new(sync.Mutex)
	cond := sync.NewCond(mutex)
	p = &Pool{
		connectPool: poolSlice,
		size:        size,
		actives:     0,
		mutex:       mutex,
		cond:        cond,
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting", err.Error())
			return nil, err // 终止程序
		}
		if _size == int(size) || p.actives == size {
			LOG.Infof("Now Server get %d connect in pool.", _size)
			return p, nil
		}
		//生成对应的Conn
		connect := &Connect.Conn{
			ID:      uuid.Must(uuid.NewV4(), nil).String(),
			TCPConn: conn.(*net.TCPConn),
		}
		connect.Ctx, connect.CtxCancel = context.WithCancel(context.Background())
		//将其保存到链表(连接池)中
		poolSlice = append(poolSlice, *connect)
		_size++
		p.actives++
	}
}

// Get 从连接池中获取连接
func (p *Pool) Get() (conn *Connect.Conn, err error) {
	//加锁，保证原子性
	p.mutex.Lock()
	defer p.mutex.Unlock()

	//判断当前连接池中连接个数
	if p.actives == 0 || len(p.connectPool) == 0 {
		LOG.Error("Now ConnectionPool is Empty. Can't get Connection.")
		return nil, nil
	}

	// conn = p.connectPool.Remove(p.connectPool.Front()).(*Connect.Conn)
	conn = &p.connectPool[0]
	p.actives--
	p.connectPool = p.connectPool[1:]
	return conn, nil
}

// Put 将连接放回连接池
func (p *Pool) Put(conn *Connect.Conn) {
	//加锁，保证原子性
	p.mutex.Lock()
	defer p.mutex.Unlock()

	//判断当前连接池个数
	if p.actives == 5 {
		//此时连接池已经满了
		LOG.Infof("Pool is Full, Can't put conn {%s} in Pool", conn.ID)
	} else {
		p.connectPool = append(p.connectPool, *conn)
		p.actives++
	}
}

// Listen 开始监听连接池中所有连接，将消息存入队列中
func (p *Pool) Listen() {
	for i := 0; i < int(p.size); i++ {

	}
}
