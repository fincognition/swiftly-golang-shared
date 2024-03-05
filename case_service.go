package shared

import (
	"context"
	"time"

	"github.com/nats-io/nats.go"
)

type Partial struct {
	Content string `json:"content"`
}

type Document struct {
	FileID   string    `json:"file_id"`
	DocType  DocType   `json:"doc_type"`
	Content  string    `json:"content"`
	Partials []Partial `json:"partials"`
}

type ValidateCaseRequest struct {
	Context
}

type ValidateCaseResponse struct {
	Context
	Valid bool `json:"valid,omitempty"`
}

func ValidateCase(ctx context.Context, nc *nats.Conn, req ValidateCaseRequest, timeout time.Duration) (*ValidateCaseResponse, error) {
	resp, err := RunRequest[ValidateCaseResponse](ctx, nc, string(ValidateCaseSubject), req, timeout)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type WriteCaseFileRequest struct {
	Context
	Filename string `json:"filename"`
	Filehash string `json:"filehash"`
	Mimetype string `json:"mimetype"`
}

type WriteCaseFileResponse struct {
	Context
	FileID string `json:"file_id"`
}

func WriteCaseFile(ctx context.Context, nc *nats.Conn, req WriteCaseFileRequest, timeout time.Duration) (*WriteCaseFileResponse, error) {
	resp, err := RunRequest[WriteCaseFileResponse](ctx, nc, string(WriteCaseFileSubject), req, timeout)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type WriteDocumentsRequest struct {
	Context
	Documents []Document `json:"documents"`
}

type WriteDocumentsResponse struct {
	Context
	Documents []string `json:"documents"`
	Partials  []string `json:"partials"`
}

func WriteDocuments(ctx context.Context, nc *nats.Conn, req WriteDocumentsRequest, timeout time.Duration) (*WriteDocumentsResponse, error) {
	resp, err := RunRequest[WriteDocumentsResponse](ctx, nc, string(WriteDocumentsSubject), req, timeout)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
