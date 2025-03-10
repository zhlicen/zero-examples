// Code generated by goctl. DO NOT EDIT!
// Source: transform.proto

package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"

	"shorturl/rpc/transform/internal/config"
	"shorturl/rpc/transform/internal/server"
	"shorturl/rpc/transform/internal/svc"
	transform "shorturl/rpc/transform/pb"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var configFile = flag.String("f", "etc/transform.yaml", "the config file")

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load server's certificate and private key
	serverCert, err := tls.LoadX509KeyPair("cert/server-cert.pem", "cert/server-key.pem")
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	c := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
	}

	return credentials.NewTLS(c), nil
}

func main() {
	flag.Parse()

	logx.DisableStat()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	transformerSrv := server.NewTransformerServer(ctx)

	s, err := zrpc.NewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		transform.RegisterTransformerServer(grpcServer, transformerSrv)
	})
	if err != nil {
		log.Fatal(err)
	}

	tlsCfg, err := loadTLSCredentials()
	if err != nil {
		log.Fatal(err)
	}

	s.AddOptions(grpc.Creds(tlsCfg))

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
