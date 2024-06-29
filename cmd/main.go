package main

import "log/slog"

func main() {
	slog.Info("Simple GRPC healt example for ArgoCD started")

	app, err := NewApp().App("localhost:8088")
	if err != nil {
		return
	}

	err = app.Run()
	if err != nil {
		return
	}
}
