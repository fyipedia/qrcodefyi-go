package qrcodefyi

// SearchResult represents the API search response.
type SearchResult struct {
	Query   string       `json:"query"`
	Results []SearchItem `json:"results"`
	Total   int          `json:"total"`
}

// SearchItem represents a single search result item.
type SearchItem struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
	Type string `json:"type"`
	URL  string `json:"url"`
}

// QRTypeDetail represents a QR code type.
type QRTypeDetail struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Category    string `json:"category"`
	URL         string `json:"url"`
}

// VersionDetail represents a QR code version.
type VersionDetail struct {
	Version     int    `json:"version"`
	Modules     int    `json:"modules"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

// ComponentDetail represents a QR code component.
type ComponentDetail struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

// EncodingDetail represents a QR code encoding mode.
type EncodingDetail struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

// StandardDetail represents a QR code standard.
type StandardDetail struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

// UseCaseDetail represents a QR code use case.
type UseCaseDetail struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

// GlossaryTerm represents a glossary term.
type GlossaryTerm struct {
	Term       string `json:"term"`
	Slug       string `json:"slug"`
	Definition string `json:"definition"`
	URL        string `json:"url"`
}

// CompareResult represents a comparison between two QR code types.
type CompareResult struct {
	TypeA interface{} `json:"type_a"`
	TypeB interface{} `json:"type_b"`
	URL   string      `json:"url"`
}
