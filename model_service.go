package shared

import (
	"context"
	"time"

	"github.com/nats-io/nats.go"
)

type ClassifyImageRequest struct {
	Context
	Path string `json:"path"`
}

type ClassifyImageResponse struct {
	Context
	Path     string    `json:"path"`
	DocTypes []DocType `json:"doc_types,omitempty"`
}

func ClassifyImage(ctx context.Context, nc *nats.Conn, req ClassifyImageRequest, timeout time.Duration) (*ClassifyImageResponse, error) {
	resp, err := RunRequest[ClassifyImageResponse](ctx, nc, string(ClassifyImageSubject), req, timeout)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ExtractPageRequest struct {
	FileContext
	DocType DocType `json:"doc_type"`
	Path    string  `json:"path"`
}

type ExtractPageResponse struct {
	FileContext
	DocType DocType     `json:"doc_type"`
	Partial interface{} `json:"partial"`
}

func ExtractPage(ctx context.Context, nc *nats.Conn, req ExtractPageRequest, timeout time.Duration) (*ExtractPageResponse, error) {
	resp, err := RunRequest[ExtractPageResponse](ctx, nc, string(ExtractPageSubject), req, timeout)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
