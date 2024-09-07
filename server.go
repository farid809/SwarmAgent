package dataservice

import (
    "context"
    "sync"

    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

type Server struct {
    UnimplementedDataServiceServer
    mu   sync.RWMutex
    data map[string]string
}

func NewServer() *Server {
    return &Server{
        data: make(map[string]string),
    }
}

func (s *Server) CreateOrUpdate(ctx context.Context, in *KeyValuePair) (*OperationResult, error) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.data[in.Key] = in.Value
    return &OperationResult{Success: true, Message: "Operation successful"}, nil
}

func (s *Server) Read(ctx context.Context, in *Key) (*KeyValuePair, error) {
    s.mu.RLock()
    defer s.mu.RUnlock()
    if value, exists := s.data[in.Key]; exists {
        return &KeyValuePair{Key: in.Key, Value: value}, nil
    }
    return nil, status.Errorf(codes.NotFound, "Key not found")
}

func (s *Server) Delete(ctx context.Context, in *Key) (*OperationResult, error) {
    s.mu.Lock()
    defer s.mu.Unlock()
    if _, exists := s.data[in.Key]; exists {
        delete(s.data, in.Key)
        return &OperationResult{Success: true, Message: "Key deleted successfully"}, nil
    }
    return &OperationResult{Success: false, Message: "Key not found"}, nil
}
