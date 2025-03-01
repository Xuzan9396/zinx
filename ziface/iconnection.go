// Package ziface 主要提供zinx全部抽象层接口定义.
// 包括:
//
//			IServer 服务mod接口
//			IRouter 路由mod接口
//			IConnection 连接mod层接口
//	     IMessage 消息mod接口
//			IDataPack 消息拆解接口
//	     IMsgHandler 消息处理及协程池接口
//
// 当前文件描述:
// @Title  iconnection.go
// @Description  全部连接相关方法声明
// @Author  Aceld - Thu Mar 11 10:32:29 CST 2019
package ziface

import (
	"context"
	"github.com/gorilla/websocket"
	"net"
)

// 定义连接接口
type IConnection interface {
	Start()                   //启动连接，让当前连接开始工作
	Stop()                    //停止连接，结束当前连接状态
	Context() context.Context //返回ctx，用于用户自定义的go程获取连接退出状态

	GetConnection() net.Conn    //从当前连接获取原始的socket
	GetWsConn() *websocket.Conn // 从当前连接中获取原始的websocket连接
	// Deprecated: use GetConnection instead
	GetTCPConnection() net.Conn //从当前连接获取原始的socket TCPConn
	GetConnID() uint64          //获取当前连接ID
	RemoteAddr() net.Addr       //获取链接远程地址信息
	LocalAddr() net.Addr        //获取链接本地地址信息

	Send(data []byte) error
	SendToQueue(data []byte) error
	SendMsg(msgID uint32, data []byte) error     //直接将Message数据发送数据给远程的TCP客户端(无缓冲)
	SendBuffMsg(msgID uint32, data []byte) error //直接将Message数据发送给远程的TCP客户端(有缓冲)

	SetProperty(key string, value interface{})   //设置链接属性
	GetProperty(key string) (interface{}, error) //获取链接属性
	RemoveProperty(key string)                   //移除链接属性
	IsAlive() bool                               //判断当前连接是否存活
	SetHeartBeat(checker IHeartbeatChecker)      //设置心跳检测器
}
