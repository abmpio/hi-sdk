package starter

import (
	"context"
	"fmt"
	"time"

	"github.com/abmpio/abmp/pkg/log"
	"github.com/abmpio/app"
	"github.com/abmpio/app/cli"
	"github.com/abmpio/hi-sdk/options"
	pb "github.com/abmpio/hi-sdk/proto"
	"github.com/abmpio/hi-sdk/sdk"
)

func init() {
	fmt.Println("abmpio.hi_sdk.starter init")

	cli.ConfigureService(serviceConfigurator)
}

func serviceConfigurator(wa cli.CliApplication) {
	if app.HostApplication.SystemConfig().App.IsRunInCli {
		return
	}
	var _client sdk.IClient

	opt := options.GetOptions()
	if !opt.Disabled {
		hiClient := sdk.NewClient(sdk.WithHost(opt.Host), sdk.WithPort(opt.Port))
		//测试ping
		for {
			err := hiClient.InitConnnection()
			if err != nil {
				log.Logger.Warn(fmt.Sprintf("初始化hi grpc连接时出现异常,option:%s, err:%s",
					opt.String(),
					err.Error()))
			} else {
				res, err := hiClient.HealthCheck(context.TODO(), &pb.HealthCheckRequest{})
				if err != nil {
					log.Logger.Warn(fmt.Sprintf("检测hi grpc 服务健康是否正常时出现异常,option:%s, err:%s",
						opt.String(),
						err.Error()))
				} else {
					if res != nil {
						log.Logger.Info(fmt.Sprintf("hi grpc status:%s", res.Status.String()))
					}
					_client = hiClient
					break
				}
			}

			log.Logger.Warn("2s后重新测试...")
			time.Sleep(2 * time.Second)
		}
	} else {
		log.Logger.Warn("hi grpc client disabled")
		_client = &sdk.NullableClient{}
	}
	sdk.SetGlobalClient(_client)
	app.Context.RegistInstanceAs(_client, new(pb.HiClient))
	app.Context.RegistInstanceAs(_client, new(pb.CodeValueServiceClient))
	app.Context.RegistInstanceAs(_client, new(pb.SettingsServiceClient))
	app.Context.RegistInstanceAs(_client, new(sdk.Client))
	app.Context.RegistInstanceAs(_client, new(pb.SmsServiceClient))
}
