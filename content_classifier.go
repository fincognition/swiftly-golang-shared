package shared

import (
	"context"
	"time"

	"github.com/nats-io/nats.go"
)

type ClassifyPageRequest struct {
	Context
	Path string `json:"path"`
}

type ClassifyPageResponse struct {
	Context
	Path     string    `json:"path"`
	DocTypes []DocType `json:"doc_types,omitempty"`
}

func ClassifyPage(ctx context.Context, nc *nats.Conn, req ClassifyPageRequest, timeout time.Duration) (*ClassifyPageResponse, error) {
	resp, err := RunRequest[ClassifyPageResponse](ctx, nc, string(ClassifyPageSubject), req, timeout)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
