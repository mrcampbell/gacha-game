package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mrcampbell/gacha-game/battle-service/internal/http"
	"github.com/mrcampbell/gacha-game/battle-service/proto/greeter"
	"google.golang.org/grpc"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial("[::1]:5000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	client := greeter.NewGreeterClient(conn)

	response, err := client.SayHello(context.Background(), &greeter.HelloRequest{Name: "Mike"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Message)

	r := gin.Default()

	http.RegisterHandlers(r)

	if err := r.Run(); err != nil { // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
		println(err)
	}
}
