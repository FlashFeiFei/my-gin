package ws

import (
	"bytes"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

const (
	// Time allowed to write a message to the peer.
	// 时间允许写一条消息给同伴。？？？？
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	//允许时间读取来自对等方的下一条pong消息。？？？？
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	//发送pings到peer。一定比彭维特小。？？？
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	//字节大小设置
	// 每次从缓冲读取的数据大小，UTF-8编码中，一个英文字符等于一个字节，一个中文（含繁体）等于三个字节。
	maxMessageSize = 512
)

//websocket升级组件,网络的io读写缓冲都为1k
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

//连接的客户定义

type Client struct {
	//客户在哪个房间
	hub *Hub

	//连接客户的句柄，也就是给这个用户或者读取这个用户信息的输入输出流
	conn *websocket.Conn

	//房间中的聊天内容
	send chan []byte
}

//ws中读信息,然后把信息给房间
func (c *Client) readPump() {

	//用户断开连接的时候，关闭这个用户的连接，和释放一些资源
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()

	//设置从io读中取出的消息大小,取出最大的字节，超过字节会断开连接
	c.conn.SetReadLimit(maxMessageSize)
	//设置读取的超时时间为1分钟
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	// WebSocket为了保持客户端、服务端的实时双向通信，需要确保客户端、服务端之间的TCP通道保持连接没有断开。
	// 然而，对于长时间没有数据往来的连接，如果依旧长时间保持着，可能会浪费包括的连接资源。
	// 但不排除有些场景，客户端、服务端虽然长时间没有数据往来，但仍需要保持连接。这个时候，可以采用心跳来实现。
	//  发送方->接收方：ping；
	//  接收方->发送方：pong；

	//心跳检测还活着回调，设置下一次的读取时间
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	for {
		//从网络中读取0.5k的字节数据
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			//获取不到信息,原因是断开了连接
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				//判断是否是异常断开，什么情况算异常断开？
				log.Printf("error: %v", err)
			}
			//结束循环，用户莫名其妙的断开的时候,关闭这个用户，和释放一些资源
			break
		}

		//消息做一下转化，将空格转换行
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))

		//信息从ws网络中获取，然后将信息给房间，让房间把信息专递给所以用户
		c.hub.broadcast <- message
	}
}

//房间收到信息之后，讲信息传递给用户
func (c *Client) writePump() {
	//设置一个定时时间，这个的作用是用来执行ping的心跳检测
	ticker := time.NewTicker(pingPeriod)

	defer func() {
		//发现用户断开连接了，死了
		//完美释放一些资源
		//倒计时停止
		ticker.Stop()
		//关闭用户的连接
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			//设置一下写入的超时时间
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				//通道关闭了，发送关闭信息
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				//关闭这个用户
				return
			}

			//获取一下这个连接的写io
			w, err := c.conn.NextWriter(websocket.TextMessage)

			if err != nil {
				//获取失败，这个用户断开连接了
				return
			}

			//写入这条消息
			w.Write(message)
			//判断一下chan是否有缓冲信息，如果有，消费完它
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}

			//关闭对用户的写
			if err := w.Close(); err != nil {
				//关闭失败，用户断开了连接
				return
			}

		case <-ticker.C:
			//心跳检测一些用户，保证他还活着
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				//连接不存活，关闭这个连接，释放一些资源
				return
			}
		}
	}
}

// 新用户加入房间
func serveWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	//讲协议升级为websocket，长连接，做个转化
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
		return
	}
	//创建一个用户，是一个有缓冲chan哦，直接缓冲
	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256)}

	//通知房间，有新用户加入，能让房间转发信息给用户，和管理用户
	client.hub.register <- client

	//执行用户在房间的读取信息和发送信息的逻辑
	go client.writePump()
	go client.readPump()
}
