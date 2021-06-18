package main

import (
	"crypto/tls"
	"fmt"
	"github.com/pingcap/ticp/micro-api/route"
	"github.com/pingcap/ticp/micro-cluster/infrastructure/adapt"
	proto2 "github.com/pingcap/ticp/micro-cluster/proto"
	"github.com/pingcap/ticp/temp"
	"log"
	"net/http"

	"github.com/asim/go-micro/plugins/wrapper/monitoring/prometheus/v3"
	mylogger "github.com/pingcap/ticp/addon/logger"
	"github.com/pingcap/ticp/addon/tracer"
	"github.com/pingcap/ticp/config"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	mlogrus "github.com/asim/go-micro/plugins/logger/logrus/v3"
	_ "github.com/asim/go-micro/plugins/registry/etcd/v3"
	"github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v3"
	"github.com/asim/go-micro/v3"
	mlog "github.com/asim/go-micro/v3/logger"
	"github.com/asim/go-micro/v3/transport"
)

func init() {
	// log
	mlog.DefaultLogger = mlogrus.NewLogger(mlogrus.WithLogger(mylogger.WithContext(nil)))
}

func main() {
	{
		// only use to init the config
		srv := micro.NewService(
			config.GetMicroCliArgsOption(),
		)
		srv.Init()
		config.Init()
		srv = nil
	}
	{
		// tls
		cert, err := tls.LoadX509KeyPair(config.GetCertificateCrtFilePath(), config.GetCertificateKeyFilePath())
		if err != nil {
			log.Fatal(err)
			return
		}
		tlsConfigPtr := &tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}
		//
		srv := micro.NewService(
			micro.Name(adapt.TICP_COMMON_SERVICE_NAME),
			micro.WrapHandler(prometheus.NewHandlerWrapper()),
			micro.WrapClient(opentracing.NewClientWrapper(tracer.GlobalTracer)),
			micro.WrapHandler(opentracing.NewHandlerWrapper(tracer.GlobalTracer)),
			micro.Transport(transport.NewHTTPTransport(transport.Secure(true), transport.TLSConfig(tlsConfigPtr))),
		)
		srv.Init()
		temp.InitClient(srv)
		{
			proto2.RegisterCommonHandler(srv.Server(), new(adapt.Common))
		}
		go func() {
			if err := srv.Run(); err != nil {
				log.Fatal(err)
			}
		}()
		// start promhttp
		http.Handle("/metrics", promhttp.Handler())
		go func() {
			addr := fmt.Sprintf(":%d", config.GetPrometheusPort())
			err := http.ListenAndServe(addr, nil)
			if err != nil {
				log.Fatal("promhttp ListenAndServe err:", err)
			}
		}()
	}
}
