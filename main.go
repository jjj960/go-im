package main

import (
	"errors"
	"log"
	"net"
	"net/rpc"
)

type UserServer struct {
}

type (
	GetUserReq struct {
		Id string `json:"id"`
	}

	GetUserResp struct {
		Id    string `json:"id"`
		Name  string `json:"name"`
		Phone string `json:"phone"`
	}
)

func (*UserServer) GetUser(req GetUserReq, resp *GetUserResp) error {
	if u, ok := users[req.Id]; ok {
		*resp = GetUserResp{
			Id:    u.Id,
			Name:  u.Name,
			Phone: u.Phone,
		}
		return nil
	}
	return errors.New("没找到相关用户")
}

func main() {
	userServer := new(UserServer)
	rpc.Register(userServer)
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("监听失败")
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("接受客户端连接信息失败")
			continue
		}
		go rpc.ServeConn(conn)
	}

}
