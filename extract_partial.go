package shared

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
