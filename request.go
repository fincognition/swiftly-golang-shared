package shared

import (
	"context"
	"encoding/json"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog"
)

func RunRequest[Response any](ctx context.Context, nc *nats.Conn, method string, req interface{}, timeout time.Duration) (*Response, error) {
	l := zerolog.Ctx(ctx)
	payload, err := json.Marshal(req)
	if err != nil {
		l.Error().Err(err).Msgf("error serializing %T request", req)
		return nil, err
	}
	resp, err := nc.Request(method, payload, timeout)
	if err != nil {
		l.Error().Err(err).Msgf("error processing %s request", method)
		return nil, err
	}
	var result Response
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		l.Error().Err(err).Msgf("error deserializing %T response", result)
		return nil, err
	}
	return &result, nil
}
