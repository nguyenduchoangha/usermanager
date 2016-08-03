package main

import (
	"flag"
	"fmt"
	sd "github.com/nguyenduchoangha/usermanager/proto/speechdata"
	um "github.com/nguyenduchoangha/usermanager/proto/usermanager"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"strconv"
	//"google.golang.org/grpc/credentials"
	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
	"google.golang.org/grpc/grpclog"
	//	"io"
)

var (
	tls              = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile           = flag.String("ca_file", "", "The file containning the CA root cert file")
	serverAddr       = flag.String("user_server_addr", "127.0.0.1:10000", "The server address in the format of host:port")
	speechServerAddr = flag.String("speech_server_addr", "127.0.0.1:10001", "The server address in the format of host:port")
)

func printToken(client um.UserManagerClient, info *um.LoginRequest) string {
	grpclog.Printf("login info (%s, %s, %s)", info.Userid, info.Prodid, info.Task)
	token, err := client.GetToken(context.Background(), info)
	if err != nil {
		grpclog.Fatalf("%v.GetToken(_) = _, %v: ", client, err)
	}
	grpclog.Println(token)
	jwt_data, _ := jws.ParseJWT([]byte(token.Token))
	err = jwt_data.Validate([]byte("hello world"), crypto.SigningMethodHS256)
	if err != nil {
		fmt.Println("Error signature")
	} else {
		fmt.Println("Ok")
	}
	return token.Token
}

func testStreaming(client sd.SpeechDataClient, tok string) {
	fmt.Println("tok:", tok)
	stream, err := client.RecordSpeech(context.Background())
	if err != nil {
		grpclog.Fatalf("%v.RecordRoute(_) = _, %v", client, err)
	}
	msg := sd.RecordRequest{&sd.RecordRequest_StreamingConfig{&sd.StreamingConfig{tok, "12345"}}}
	if err := stream.Send(&msg); err != nil {
		grpclog.Fatalf("%v.Send(_) = _, %v", client, err)
	}
	for i := 0; i < 1000; i += 100 {
		msg1 := sd.RecordRequest{&sd.RecordRequest_AudioContent{[]byte(strconv.Itoa(i))}}
		if err := stream.Send(&msg1); err != nil {
			grpclog.Fatalf("%v.Send(_) = _, %v", client, err)
		}
	}

}

func main() {
	flag.Parse()

	// Get Token

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := um.NewUserManagerClient(conn)

	tok := printToken(client, &um.LoginRequest{"user46", "password", "s2t"})

	// Send audio

	conn1, err := grpc.Dial(*speechServerAddr, opts...)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	defer conn1.Close()
	client1 := sd.NewSpeechDataClient(conn1)

	testStreaming(client1, tok)

}
