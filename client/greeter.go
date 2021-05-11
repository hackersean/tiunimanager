package client

import (
	greeterPb "tcp/proto/greeter"

	_ "github.com/asim/go-micro/plugins/registry/etcd/v3"
	"github.com/asim/go-micro/v3"
)

// Make request
/*
	rsp, err := GreeterClient.Hello(context.Background(), &pb.Request{
		Name: "Foo",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp.Greeting)
*/
var GreeterClient greeterPb.GreeterService

func init() {
	// create a new service
	service := micro.NewService()

	// parse command line flags
	service.Init()

	// Use the generated client stub
	GreeterClient = greeterPb.NewGreeterService("greeter", service.Client())
}
