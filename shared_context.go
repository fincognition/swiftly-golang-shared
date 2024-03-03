package shared

// Context is shared context for all microservice requests
type Context struct {
	RequestID string `json:"request_id"`
	ClientID  string `json:"client_id"`
	CaseID    string `json:"case_id"`
}

type FileContext struct {
	Context
	FileID string `json:"file_id"`
}
