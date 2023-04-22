package beacon

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"

	api "github.com/troydai/grpcbeacon/gen/api/protos"
	"github.com/troydai/grpcbeacon/internal/settings"
)

type service struct {
	api.UnimplementedBeaconServer

	details map[string]string
	logger  *zap.Logger
}

var _ api.BeaconServer = (*service)(nil)

func newService(env settings.Environment, logger *zap.Logger) *service {
	s := &service{
		details: make(map[string]string),
		logger:  logger,
	}

	s.details["Hostname"] = env.HostName
	s.details["BeaconName"] = env.BeaconName

	return s
}

func (s *service) Signal(ctx context.Context, req *api.SignalRequest) (*api.SignalResponse, error) {
	logger := s.logger
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		logger = logger.With(zap.Any("metadata", md))
	}

	logger.Info("Signal received")
	resp := &api.SignalResponse{
		Reply: fmt.Sprintf("Beacon signal at %s", time.Now().Format(time.RFC1123)),
	}

	if len(s.details) > 0 {
		resp.Details = s.details
	}

	return resp, nil
}
