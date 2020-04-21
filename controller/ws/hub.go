package ws

//聊天的房间定义
//聊天室的作用，是把一个用户发送的信息通知给其他在这个房间的用户


type Hub struct {
	//房间中的客户
	clients map[*Client]bool

	//通知给用户新信息
	broadcast chan []byte

	//新加入房间的用户
	register chan *Client

	//从房间离开的用户
	unregister chan *Client
}

//创建一个房间
func newHub() *Hub {
	return &Hub{
		clients: make(map[*Client]bool),
		//为什么定义为无缓冲chan,作为广播呢?，保证接收端和发送端都同时准备好？不要一直接收消息，确保不会出现并发，数据库出现脏数据?确保数据的原子性
		broadcast: make(chan []byte),
		//为什么定义为无缓冲chan，作为注册呢？保证接收端和发送端都同时准备好？不要一直接收消息，确保不会出现并发，数据库出现脏数据?确保数据的原子性
		register: make(chan *Client),
		//为什么定义为无缓冲chan呢？,保证接收端和发送端都同时准备好？不要一直接收消息，确保不会出现并发，数据库出现脏数据?确保数据的原子性
		unregister: make(chan *Client),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			//房间加入了一个用户
			h.clients[client] = true
		case client := <-h.unregister:
			//房间离开了一个用户
			//删除这个用户，并且释放用户的一写内存
			if _, ok := h.clients[client]; ok {
				//删除用户
				delete(h.clients, client)
				//释放用户接收信息
				close(client.send)
			}

		case message := <-h.broadcast:
			//房间收到信息，把消息传递给用户
			for client, _ := range h.clients {
				select {
				case client.send <- message:
					//消息传递给用户
				default:
					//case是无序的，当任何一个case都不起作用的时候，会走到default
					//这里的意思，也就是消息传递不了给这个用户，也就是用户断开连接了
					//释放用户的一些资源
					delete(h.clients, client)
					close(client.send)
				}
			}
		}
	}
}