package sdk

import (
	"fmt"

	"github.com/abmpio/abmp/pkg/log"
	pb "github.com/abmpio/hi/sdk/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	option *Options

	conn *grpc.ClientConn

	pb.HiClient
	pb.CodeValueServiceClient
	pb.SettingsServiceClient
	pb.AppServiceClient
}

func NewClient(opts ...Option) *Client {
	client := &Client{
		option: newDefaultOptions(),
	}
	for _, eachOpt := range opts {
		eachOpt(client.option)
	}
	return client
}

func (c *Client) GetOption() *Options {
	return c.option
}

// 初始化client
func (c *Client) InitConnnection(opts ...grpc.DialOption) error {
	mergedOpts := make([]grpc.DialOption, 0)
	mergedOpts = append(mergedOpts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	mergedOpts = append(mergedOpts, grpc.WithKeepaliveParams(*c.option.keepaliveParam))
	mergedOpts = append(mergedOpts, opts...)
	conn, err := grpc.NewClient(c.option.getHostTarget(), mergedOpts...)
	if err != nil {
		log.Logger.Error(fmt.Sprintf("occur error when create grpc server connection , host:%s,error: %s",
			c.option.getHostTarget(), err.Error()))
		return err
	}
	log.Logger.Info(fmt.Sprintf("initialize grpc connection finished,host:%s", c.option.getHostTarget()))
	c.conn = conn
	//保存客户端
	c.HiClient = pb.NewHiClient(conn)
	c.CodeValueServiceClient = pb.NewCodeValueServiceClient(conn)
	c.SettingsServiceClient = pb.NewSettingsServiceClient(conn)
	c.AppServiceClient = pb.NewAppServiceClient(conn)

	return nil
}
