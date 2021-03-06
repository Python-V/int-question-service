package server

import (
	"context"
	"encoding/json"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"questionService/config"
	gw "questionService/libs"
	question "questionService/libs"
	"questionService/libs/auth"
	"questionService/methods"
	"strings"
)

var AuthClient auth.AuthClient

func GrpcAuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("--> unary interceptor: %v", info.FullMethod)
	return handler(ctx, req)
}

func StartGrpcServer() {
	lis, err := net.Listen("tcp", config.GrpcPort)

	if err != nil {
		log.Fatalf("Error in starting server %v", err)
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(GrpcAuthInterceptor),
	)
	question.RegisterQuestionServiceServer(s, &methods.Server{})
	reflection.Register(s)

	log.Printf("Listening gRPC at : %v", lis.Addr())

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to Serve %v", err)
		}
	}()

	// wait for ctrl c to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// Block until signal is received
	<-ch
	log.Println("Stopping Server...")
	s.Stop()
	log.Println("Stopping Listener...")
	lis.Close()
	log.Println("Server Stopped.")
}

func GatewayAuthenticate(token string) error {
	if !strings.Contains(token, "Bearer") {
		return status.Errorf(codes.InvalidArgument, "Invalid auth token format")
	}
	token = strings.ReplaceAll(token, "Bearer ", "")
	req := &auth.TokenValidatorRequest{Bearer: token}
	res, err := AuthClient.ValidateToken(context.Background(), req)
	if err != nil {
		return err
	}
	if valid := res.Success; valid {
		return nil
	}
	return status.Errorf(codes.Unauthenticated, res.Msg)
}

func StartGatewayServer() {
	// mux
	mux := runtime.NewServeMux()
	// register
	err := gw.RegisterQuestionServiceHandlerServer(context.Background(), mux, &methods.Server{})
	if err != nil {
		panic(err.Error())
	}
	gwServer := &http.Server{
		Addr: config.GwRestPort,
		// Handle authentication through auth interceptor
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			bearer := r.Header.Get("Authorization")
			// Call grpc auth server
			err := GatewayAuthenticate(bearer)
			if err == nil {
				mux.ServeHTTP(w, r)
				return
			}
			log.Printf("Error in Authentication: %v", err)
			// Case: Invalid auth token, write message to response writer object
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			err = json.NewEncoder(w).Encode(map[string]string{"code": "16", "message": "Unauthorized User"})
			if err != nil {
				log.Printf("Error writing to response writer: %v", err)
				return
			}
		}),
	}
	log.Printf("Listening gateway REST at : %v", config.GwRestPort)
	// http server
	log.Fatalln(gwServer.ListenAndServe())
}
