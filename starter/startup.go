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
	opt := options.GetOptions()
	_client := sdk.NewClient(sdk.WithHost(opt.Host), sdk.WithPort(opt.Port))
	//测试ping
	for {
		err := _client.InitConnnection()
		if err == nil {
			var res *pb.HealthCheckResponse
			res, err = _client.HealthCheck(context.TODO(), &pb.HealthCheckRequest{})
			if err == nil {
				break
			}
			if res != nil {
				log.Logger.Info(res.Status.String())
			}
		}

		log.Logger.Warn(err.Error())
		log.Logger.Warn("2s后重新测试...")
		time.Sleep(2 * time.Second)
	}
	sdk.SetGlobalClient(_client)
	app.Context.RegistInstanceAs(_client, new(pb.HiClient))
	app.Context.RegistInstanceAs(_client, new(pb.CodeValueServiceClient))
	app.Context.RegistInstanceAs(_client, new(pb.SettingsServiceClient))
	app.Context.RegistInstanceAs(_client, new(sdk.Client))
	app.Context.RegistInstanceAs(_client, new(pb.SmsServiceClient))
}
