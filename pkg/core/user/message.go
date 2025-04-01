package player

// 消息类型
type (
	MessageType = uint32
	TMessage    = Message
)

// 消息类型
type Message struct {
	Seq   int64       `msgpack:"s"`           // Sequence number
	Ts    int64       `msgpack:"t"`           // Timestamp
	Event uint32      `msgpack:"e"`           // Event ID
	Body  interface{} `msgpack:"b,omitempty"` // Raw message body
}

type LoginReqBody struct {
	Token  string `msgpack:"token"`
	Device string `msgpack:"d"`
}
