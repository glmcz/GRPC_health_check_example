package main

import (
	"fmt"
	"github.com/oklog/run"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"log/slog"
	"net"
)

type Factory struct{}

func NewApp() *Factory {
	return &Factory{}
}

func (f *Factory) App(AddressGRPCServer string) (*run.Group, error) {
	r := run.Group{}

	svr := NewService()
	s := grpc.NewServer()

	grpc_health_v1.RegisterHealthServer(s, svr)
	reflection.Register(s)

	r.Add(func() error {
		slog.Info("start grpc server", "on address", AddressGRPCServer)

		lis, err := net.Listen("tcp", AddressGRPCServer)
		if err != nil {
			slog.Error(fmt.Errorf("grpc server listening on: %s: %w", AddressGRPCServer, err).Error())
		}

		return s.Serve(lis)
	}, func(error) {
		s.GracefulStop()
		s.Stop()
	})

	return &r, nil
}
