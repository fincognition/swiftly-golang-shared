package shared

import (
	"context"
	"time"

	"github.com/nats-io/nats.go"
)

type CaseFile struct {
	ID    string   `json:"id"`
	Name  string   `json:"name"`
	Hash  uint64   `json:"hash"`
	Pages []string `json:"pages"`
}

type ProcessFilesRequest struct {
	Context
}

type ProcessFilesResponse struct {
	Context
	Files []CaseFile `json:"files,omitempty"`
}

func ProcessFiles(ctx context.Context, nc *nats.Conn, req ProcessFilesRequest, timeout time.Duration) (*ProcessFilesResponse, error) {
	resp, err := RunRequest[ProcessFilesResponse](ctx, nc, string(ProcessFilesSubject), req, timeout)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type CleanupFilesRequest struct {
	Context
}

type CleanupFilesResponse struct {
	Context
}

func CleanupFiles(ctx context.Context, nc *nats.Conn, req CleanupFilesRequest, timeout time.Duration) (*CleanupFilesResponse, error) {
	resp, err := RunRequest[CleanupFilesResponse](ctx, nc, string(CleanupFilesSubject), req, timeout)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
