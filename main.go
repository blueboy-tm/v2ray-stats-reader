package main

import (
	"context"
	"strings"
	"time"

	"github.com/jessevdk/go-flags"
	"github.com/v2fly/v2ray-core/v4/app/stats/command"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var opts struct {
	V2RayEndpoint          string `short:"e" long:"v2ray-endpoint" description:"V2Ray API endpoint" value-name:"HOST:PORT" default:"127.0.0.1:8080"`
	ScrapeTimeoutInSeconds int64  `short:"t" long:"timeout" description:"The timeout in seconds" value-name:"N" default:"3"`
	Version                bool   `short:"v" long:"version" description:"Display the version and exit"`
}

func main() {
	var err error
	if _, err = flags.Parse(&opts); err != nil {
		return
	}

	if opts.Version {
		println("v2ray-stats-reader v1.0.0")
		return
	}
	timeout := time.Duration(opts.ScrapeTimeoutInSeconds) * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	conn, err := grpc.DialContext(ctx, opts.V2RayEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return
	}
	client := command.NewStatsServiceClient(conn)
	resp, err := client.QueryStats(ctx, &command.QueryStatsRequest{Reset_: false})
	if err != nil {
		return
	}

	uplink := 0
	downlink := 0
	for _, s := range resp.GetStat() {
		p := strings.Split(s.GetName(), ">>>")
		metric := p[2] + "_" + p[3]
		if metric == "traffic_uplink" {
			uplink += int(s.GetValue())
		} else if metric == "traffic_downlink" {
			downlink += int(s.GetValue())
		}
	}
	println("uplink", uplink)
	println("downlink", downlink)

	resp2, err := client.GetSysStats(ctx, &command.SysStatsRequest{})
	if err != nil {
		return
	}
	println("duration", resp2.GetUptime())
}
