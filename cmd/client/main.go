package main

import (
	"context"
	"fmt"
	pb "github.com/BitofferHub/proto_center/api/xtimer/v1"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	transgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
)

func main() {
	callGRPC()
	//callHTTP()
	// callGRPCDiscover()
}

// callGRPC
//
//	@Author <a href="https://bitoffer.cn">狂飙训练营</a>
//	@Description:  just a demo for rpc call without discover
func callGRPC() {
	conn, err := transgrpc.DialInsecure(
		context.Background(),
		transgrpc.WithEndpoint("127.0.0.1:6001"),
		transgrpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := pb.NewXTimerClient(conn)
	reply, err := client.CreateTimer(context.Background(), &pb.CreateTimerRequest{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("[grpc] CreateUser reply %+v\n", reply)

}

// callGRPCDiscover
//
//	@Author <a href="https://bitoffer.cn">狂飙训练营</a>
//	@Description: just a demo for rpc call with discovery
func callGRPCDiscover() {
	// new etcd client
	client, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"127.0.0.1:2379"},
	})
	if err != nil {
		panic(err)
	}
	// new dis with etcd client
	dis := etcd.New(client)

	endpoint := "discovery:///user-svr"
	conn, err := transgrpc.DialInsecure(context.Background(), transgrpc.WithEndpoint(endpoint), transgrpc.WithDiscovery(dis))
	if err != nil {
		panic(err)
	}

	defer conn.Close()
	cli := pb.NewXTimerClient(conn)
	reply, err := cli.EnableTimer(context.Background(), &pb.EnableTimerRequest{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("[grpc] CreateUser reply %+v\n", reply)

}
