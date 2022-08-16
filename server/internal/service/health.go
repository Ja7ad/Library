package service

import (
	"context"
	"google.golang.org/grpc/health/grpc_health_v1"
)

type HealthyServer struct {
	grpc_health_v1.UnimplementedHealthServer
}

func (h *HealthyServer) Check(ctx context.Context, in *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	return &grpc_health_v1.HealthCheckResponse{Status: grpc_health_v1.HealthCheckResponse_SERVING}, nil
}
