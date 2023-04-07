package gozero

import (
	"fmt"
	"github.com/go-packagist/example/app/framework/foundation"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Kernel struct {
	*foundation.Application
}

func NewKernel(app *foundation.Application) *Kernel {
	return &Kernel{
		Application: app,
	}
}

func (k *Kernel) Boot() {
	k.Application.Boot()

	k.start()
}

func (k *Kernel) start() {
	conf := k.Config.Gozero.RpcServerConf

	s := zrpc.MustNewServer(*conf, func(grpcServer *grpc.Server) {

		if conf.Mode == service.DevMode || conf.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", conf.ListenOn)
	s.Start()
}
