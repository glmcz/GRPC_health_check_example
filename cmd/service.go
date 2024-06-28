package main

import (
	"context"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

type Service struct {
	server healthpb.UnimplementedHealthServer
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Check(context.Context, *healthpb.HealthCheckRequest) (*healthpb.HealthCheckResponse, error) {
	println("Health check req come")
	return nil, nil
}

func (s *Service) Watch(_ *healthpb.HealthCheckRequest, _ healthpb.Health_WatchServer) error {
	panic("implement me")
}
