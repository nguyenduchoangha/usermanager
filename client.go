package main

import (
	"flag"
	pb "github.com/nguyenduchoangha/usermanager/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	//"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	//	"io"
)

var (
	tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile     = flag.String("ca_file", "", "The file containning the CA root cert file")
	serverAddr = flag.String("server_addr", "127.0.0.1:10000", "The server address in the format of host:port")
)

func printToken(client pb.UserManagerClient, info *pb.LoginRequest) {
	grpclog.Printf("login info (%s, %s, %s)", info.Userid, info.Prodid, info.Task)
	token, err := client.GetToken(context.Background(), info)
	if err != nil {
		grpclog.Fatalf("%v.GetToken(_) = _, %v: ", client, err)
	}
	grpclog.Println(token)
}

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewUserManagerClient(conn)

	// Looking for a valid feature
	printToken(client, &pb.LoginRequest{"user46", "password", "s2t"})

}
