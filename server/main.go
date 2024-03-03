package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/m3talsmith/packets/server/api"
	"google.golang.org/grpc"
)

var (
	host = flag.String("host", "localhost", "The host address")
	port = flag.Int("port", 3000, "The host port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *host, *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	api.RegisterApiServer(server, &api.Service{})
	log.Printf("server listening: %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
