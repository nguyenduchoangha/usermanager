package main

import (
	"flag"
	"fmt"
	pb "github.com/nguyenduchoangha/usermanager/proto/speechdata"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"io"
	"net"
)

var (
	tls      = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile = flag.String("cert_file", "testdata/server1.pem", "The TLS cert file")
	keyFile  = flag.String("key_file", "testdata/server1.key", "The TLS key file")
	port     = flag.Int("port", 10001, "The server port")
)

type speechDataServer struct{}

func (s *speechDataServer) RecordSpeech(stream pb.SpeechData_RecordSpeechServer) error {
	msg0, _ := stream.Recv()
	conf := msg0.GetStreamingConfig()
	fmt.Println("got a config, token: %s", conf.Tok)

	for {
		msg_i, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.RecordSummary{"hello", ""})
		}
		if err != nil {
			return err
		}
		audio := msg_i.GetAudioContent()
		fmt.Println("got  a chunk; len ", len(audio), "; data: ", audio)
	}

}

func newServer() *speechDataServer {
	s := new(speechDataServer)
	return s
}

func main() {
	fmt.Println("Hello")
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	if *tls {
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			grpclog.Fatalf("Failed to generate credentials %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterSpeechDataServer(grpcServer, newServer())
	grpcServer.Serve(lis)

}
