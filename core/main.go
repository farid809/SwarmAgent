package main

import (
	"context"
	"log"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "github.com/farid809/SwarmAgent/dataservice"
)

// DataStore represents our in-memory data store
type DataStore struct {
	sync.RWMutex
	data map[string]string
}

// Server implements the gRPC server
type Server struct {
	pb.UnimplementedDataServiceServer
	store *DataStore
}

// CreateOrUpdate implements the CreateOrUpdate RPC method
func (s *Server) CreateOrUpdate(ctx context.Context, in *pb.KeyValuePair) (*pb.OperationResult, error) {
	s.store.Lock()
	defer s.store.Unlock()
	s.store.data[in.Key] = in.Value
	return &pb.OperationResult{Success: true, Message: "Operation successful"}, nil
}

// Read implements the Read RPC method
func (s *Server) Read(ctx context.Context, in *pb.Key) (*pb.KeyValuePair, error) {
	s.store.RLock()
	defer s.store.RUnlock()
	if value, exists := s.store.data[in.Key]; exists {
		return &pb.KeyValuePair{Key: in.Key, Value: value}, nil
	}
	return nil, status.Errorf(codes.NotFound, "Key not found")
}

// Delete implements the Delete RPC method
func (s *Server) Delete(ctx context.Context, in *pb.Key) (*pb.OperationResult, error) {
	s.store.Lock()
	defer s.store.Unlock()
	if _, exists := s.store.data[in.Key]; exists {
		delete(s.store.data, in.Key)
		return &pb.OperationResult{Success: true, Message: "Key deleted successfully"}, nil
	}
	return &pb.OperationResult{Success: false, Message: "Key not found"}, nil
}

// eventProcessor simulates processing events and updating the data store
func eventProcessor(store *DataStore, done chan bool) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			store.Lock()
			store.data["lastProcessed"] = time.Now().String()
			store.Unlock()
			log.Println("Event processed")
		case <-done:
			return
		}
	}
}

func main() {
	// Initialize the data store
	store := &DataStore{
		data: make(map[string]string),
	}

	// Start the event processing goroutine
	done := make(chan bool)
	go eventProcessor(store, done)

	// Set up the gRPC server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterDataServiceServer(s, &Server{store: store})

	// Start the gRPC server
	log.Println("Starting gRPC server on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	// Signal the event processor to stop (this won't be reached in this simple example)
	done <- true
}