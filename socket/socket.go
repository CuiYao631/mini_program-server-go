package socket

import (
	"bufio"
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"net"
	"time"
)

//客户端的连接信息
type client_info struct {
	conn      net.Conn
	conn_time time.Time
}

var MapClient = make(map[string]client_info)

type socket struct {
	listen net.Listener
	opAi   *openai.Client
}

func MakeSocket(listen net.Listener, opAi *openai.Client) *socket {
	return &socket{listen: listen, opAi: opAi}
}

func (s *socket) RunSocket() {
	fmt.Println("服务端已经启动等待客户端连接......")
	for {
		conn, err := s.listen.Accept() // 监听客户端建立连接
		client_ip := conn.RemoteAddr()
		MapClient[client_ip.String()] = client_info{conn, time.Now()}
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		} else {
			fmt.Println(fmt.Sprintf("客户端%s建立连接,当前在线数量:%d", client_ip.String(), len(MapClient)))
		}
		go s.ReceiveMessage(conn) // 启动一个goroutine处理连接
	}
}

// 处理函数
func (s *socket) ReceiveMessage(conn net.Conn) {
	defer conn.Close() // 关闭连接
	for {
		reader := bufio.NewReader(conn)
		msg, err := Decode(reader)
		client_ip := conn.RemoteAddr()
		if err != nil {
			delete(MapClient, client_ip.String())
			fmt.Println(fmt.Sprintf("客户端%s断开连接, err:%s,当前在线客户端数量:%d", client_ip.String(), err, len(MapClient))) //客户端断开连接
			break
		}
		fmt.Println("收到客户端"+client_ip.String()+"端发来的数据：", msg)
		if msg != "" {
			//SendMessage(conn, msg)
			s.Chat(context.Background(), conn, msg)
		}
	}
}

//将发送的消息进行编码

func Encode(message string) ([]byte, error) {
	//先读取消息的长度，转成int32（占4个字节）
	var length = int32(len(message))
	//开辟一个字节缓冲区
	var pkg = new(bytes.Buffer)
	//先将消息长度写入缓冲区
	var err error
	err = binary.Write(pkg, binary.LittleEndian, length) //LittleEndian按大端的顺序写
	//开始写入消息实体
	err = binary.Write(pkg, binary.LittleEndian, []byte(message))
	if err != nil {
		return nil, err
	}
	return pkg.Bytes(), err
}

//消息解码

func Decode(reader *bufio.Reader) (string, error) {
	//先读取消息的长度
	var err error
	var length int32
	var lengthByte []byte
	//读取前4个字节的数据
	lengthByte, err = reader.Peek(4) //返回前面4个字节的字节数据
	// NewBuffer的目的是准备一个Buffer来读取现有数据
	lengthBuff := bytes.NewBuffer(lengthByte)
	//读取出数据包的长度
	err = binary.Read(lengthBuff, binary.LittleEndian, &length)
	if err != nil {
		return "", err
	}
	//小于4个字节不读取数据
	if int32(reader.Buffered()) < length+4 {
		return "", err
	}
	//读取真正的数据消息
	pkg := make([]byte, int(4+length))
	_, err = reader.Read(pkg)
	if err != nil {
		return "", err
	}
	return string(pkg[4:]), nil
}
