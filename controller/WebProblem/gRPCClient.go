package WebProblem

import (
	"CTF/controller/WebProblem/model"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func InitGrpcClient(port string) (*grpc.ClientConn, model.DockerServerClient) {
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		fmt.Println("err:", err)
	}
	productServer := model.NewDockerServerClient(conn)
	return conn, productServer
}

func Client(port string, operate <-chan int, dockerID <-chan string) {
	fmt.Println("进入客户端成功")
	conn, product := InitGrpcClient(port)
	defer conn.Close()
	for {
		for i := range operate {
			switch i {
			case 1:
				fmt.Println("进入分支成功")
				docId := <-dockerID
				fmt.Println("docId:", docId)
				res, err := product.GetStartContainUrl(context.Background(), &model.StartContainRequest{ContainerID: docId})
				if err != nil {
					fmt.Println("调用失败", err)
				} else {
					fmt.Println("结果：", res.ContainUrl)
				}
			}
		}
	}
}
