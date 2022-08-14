package nnc

import (
	"crypto/tls"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"tiny-dfs/gen-go/tdfs"
)

func NewNameNodeClient(addr string) (*tdfs.NameNodeClient, error) {
	fmt.Println("running client")

	var protocolFactory thrift.TProtocolFactory
	protocolFactory = thrift.NewTBinaryProtocolFactoryConf(nil)

	var transportFactory thrift.TTransportFactory
	transportFactory = thrift.NewTBufferedTransportFactory(8192)

	cfg := &thrift.TConfiguration{
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	// 建立连接
	var transport thrift.TTransport
	transport = thrift.NewTSocketConf(addr, cfg)
	transport, err := transportFactory.GetTransport(transport)
	if err != nil {
		return nil, err
	}
	if err := transport.Open(); err != nil {
		return nil, err
	}
	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)
	return tdfs.NewNameNodeClient(thrift.NewTStandardClient(iprot, oprot)), nil
}
