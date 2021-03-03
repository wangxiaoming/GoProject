/** thrift_example.go */
package main

import (
	"GoProject/src/thrift/gen-go/example"
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"log"
	"net"
	"os"
)

const (
	HOST = "localhost"
	PORT = "19090"
)

type TransdataImpl struct {
}

func (trandata *TransdataImpl) SendMsg(ctx context.Context, msgJson string) (r bool, err error) {
	fmt.Println("-->SendMsg Call:", msgJson)
	return true, nil
}

func Server() {
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	//protocolFactory := thrift.NewTCompactProtocolFactory()

	serverTransport, err := thrift.NewTServerSocket(net.JoinHostPort(HOST, PORT))
	if err != nil {
		fmt.Println("Error!", err)
		os.Exit(1)
	}

	handler := &TransdataImpl{}
	processor := example.NewTransdataProcessor(handler)

	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	fmt.Println("thrift server in", net.JoinHostPort(HOST, PORT))
	server.Serve()
}

func Client() {
	tSocket, err := thrift.NewTSocket(net.JoinHostPort(HOST, PORT))
	if err != nil {
		log.Fatalln("tSocket error:", err)
	}
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	transport, _ := transportFactory.GetTransport(tSocket)
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	client := example.NewTransdataClientFactory(transport, protocolFactory)

	if err := transport.Open(); err != nil {
		log.Fatalln("Error opening:", HOST+":"+PORT)
	}
	defer transport.Close()

	d, err := client.SendMsg(nil, "test string")
	fmt.Println(d)
}
