package main

import (
	"flag"
	"fmt"
	//	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"

	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"

	pb "github.com/nguyenduchoangha/usermanager/proto"
)

var (
	tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile   = flag.String("cert_file", "testdata/server1.pem", "The TLS cert file")
	keyFile    = flag.String("key_file", "testdata/server1.key", "The TLS key file")
	jsonDBFile = flag.String("json_db_file", "testdata/route_guide_db.json", "A json file containing a list of features")
	port       = flag.Int("port", 10000, "The server port")
)

type userManagerServer struct{}

func (s *userManagerServer) GetToken(ctx context.Context, info *pb.LoginRequest) (*pb.LoginReply, error) {
	fmt.Println(info.Userid)
	fmt.Println(info.Prodid)
	fmt.Println(info.Task)
	return &pb.LoginReply{"xxx.yyy.zzz", ""}, nil
}

func newServer() *userManagerServer {
	s := new(userManagerServer)
	return s
}

func main() {
	fmt.Println("hello world")
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
	pb.RegisterUserManagerServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
