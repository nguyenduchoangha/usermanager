package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
	//	"github.com/golang/protobuf/proto"
	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
	//"github.com/SermoDigital/jose/jwt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"net"

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

func createJWT() string {
	claims := jws.Claims{}
	claims.SetSubject("s2t")
	claims.SetExpiration(time.Now().AddDate(0, 0, 1))
	//signMethod := jws.GetSigningMethod("HS256")
	token := jws.NewJWT(claims, crypto.SigningMethodHS256)
	signKey := []byte("hello world")
	byteToken, err := token.Serialize(signKey)
	if err != nil {
		log.Fatal("Error signing the key. ", err)
		os.Exit(1)
	}

	return string(byteToken)
}

func (s *userManagerServer) GetToken(ctx context.Context, info *pb.LoginRequest) (*pb.LoginReply, error) {
	fmt.Println(info.Userid)
	fmt.Println(info.Prodid)
	fmt.Println(info.Task)
	tok := createJWT()
	return &pb.LoginReply{tok, ""}, nil
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
