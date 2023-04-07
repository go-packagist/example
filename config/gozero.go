package config

import "github.com/zeromicro/go-zero/zrpc"

type Gozero struct {
	*zrpc.RpcServerConf
}

var gozero = &Gozero{
	RpcServerConf: &zrpc.RpcServerConf{},
}
