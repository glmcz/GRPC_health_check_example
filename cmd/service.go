package main

import (
	"context"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"log/slog"
)

type Service struct {
	server healthpb.UnimplementedHealthServer
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Check(context.Context, *healthpb.HealthCheckRequest) (*healthpb.HealthCheckResponse, error) {
	println("Health check req come")
	slog.Info("Health check req come")
	healthCheckResponse := &healthpb.HealthCheckResponse{
		Status: healthpb.HealthCheckResponse_SERVING,
	}
	return healthCheckResponse, nil
}

func (s *Service) Watch(_ *healthpb.HealthCheckRequest, _ healthpb.Health_WatchServer) error {
	panic("implement me")
}
