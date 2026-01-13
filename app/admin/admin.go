// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/lxjianbai/swift/app/admin/internal/config"
	"github.com/lxjianbai/swift/app/admin/internal/handler"
	"github.com/lxjianbai/swift/app/admin/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/admin.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	if c.Mode == service.DevMode || c.Mode == service.TestMode {
		if c.Log.Mode == "file" {
			logx.AddWriter(logx.NewWriter(os.Stdout))
		}
	}

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
