package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	example "github.com/shukubota/grpc-playground/gen/go/proto"
	"github.com/shukubota/grpc-playground/handler"
	"golang.org/x/sync/errgroup"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
)

const grpcServerAddress = "localhost:5001"

func main() {
	// grpc pure
	srv := grpc.NewServer()
	api := handler.NewExampleAPIServer()
	example.RegisterExampleServer(srv, api)

	lis, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatalf("failed to listen(tcp, :5001)")
	}

	// grpc gateway
	grpcGateway := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	withCors := cors.New(cors.Options{
		//AllowOriginFunc: func(origin string) bool { return true },
		AllowedOrigins: []string{
			"http://127.0.0.1:5174",
		},
		AllowedMethods: []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		//AllowedHeaders: []string{"ACCEPT", "Authorization", "Content-Type", "X-CSRF-Token"},
		//ExposedHeaders:   []string{"Link"},
		//AllowCredentials: true,
		MaxAge: 300,
	}).Handler(grpcGateway)

	if err := example.RegisterExampleHandlerFromEndpoint(context.Background(), grpcGateway, grpcServerAddress, opts); err != nil {
		log.Fatal("failed to register grpc-server")
	}

	ctx := context.Background()

	eg, ctx := errgroup.WithContext(ctx)

	fmt.Println(lis)
	eg.Go(func() error {
		log.Printf("grpc server started at port: 5001")
		if err := srv.Serve(lis); err != nil {
			log.Fatalf("err has occured while serving: %v", err)
			return err
		}
		return nil
	})
	eg.Go(func() error {
		log.Printf("grpc gateway server started at port: 8085")
		if err := http.ListenAndServe(":8085", withCors); err != nil {
			log.Fatal("err")
			return err
		}
		return nil
	})
	eg.Wait()
}
