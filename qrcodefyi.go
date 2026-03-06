// Package qrcodefyi provides a Go client for the QRCodeFYI API.
//
// QRCodeFYI is a comprehensive QR code reference covering types, versions,
// components, encodings, and standards. This client requires no authentication
// and has zero external dependencies.
//
// Usage:
//
//	client := qrcodefyi.NewClient()
//	result, err := client.Search("micro")
package qrcodefyi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// DefaultBaseURL is the default base URL for the QRCodeFYI API.
const DefaultBaseURL = "https://qrcodefyi.com/api"

// Client is a QRCodeFYI API client.
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

// NewClient creates a new QRCodeFYI API client with default settings.
func NewClient() *Client {
	return &Client{
		BaseURL:    DefaultBaseURL,
		HTTPClient: &http.Client{},
	}
}

func (c *Client) get(path string, result interface{}) error {
	resp, err := c.HTTPClient.Get(c.BaseURL + path)
	if err != nil {
		return fmt.Errorf("qrcodefyi: request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("qrcodefyi: HTTP %d: %s", resp.StatusCode, string(body))
	}

	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return fmt.Errorf("qrcodefyi: decode failed: %w", err)
	}
	return nil
}

// Search searches across QR code types, standards, and glossary terms.
func (c *Client) Search(query string) (*SearchResult, error) {
	var result SearchResult
	path := fmt.Sprintf("/search/?q=%s", url.QueryEscape(query))
	if err := c.get(path, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// QRType returns details for a QR code type by slug.
func (c *Client) QRType(slug string) (*QRTypeDetail, error) {
	var result QRTypeDetail
	if err := c.get("/type/"+slug+"/", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Version returns details for a QR code version by number.
func (c *Client) Version(version int) (*VersionDetail, error) {
	var result VersionDetail
	path := fmt.Sprintf("/version/%d/", version)
	if err := c.get(path, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Component returns details for a QR code component by slug.
func (c *Client) Component(slug string) (*ComponentDetail, error) {
	var result ComponentDetail
	if err := c.get("/component/"+slug+"/", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Encoding returns details for a QR code encoding mode by slug.
func (c *Client) Encoding(slug string) (*EncodingDetail, error) {
	var result EncodingDetail
	if err := c.get("/encoding/"+slug+"/", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Standard returns details for a QR code standard by slug.
func (c *Client) Standard(slug string) (*StandardDetail, error) {
	var result StandardDetail
	if err := c.get("/standard/"+slug+"/", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UseCase returns details for a QR code use case by slug.
func (c *Client) UseCase(slug string) (*UseCaseDetail, error) {
	var result UseCaseDetail
	if err := c.get("/use-case/"+slug+"/", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GlossaryTerm returns a glossary term by slug.
func (c *Client) GlossaryTerm(slug string) (*GlossaryTerm, error) {
	var result GlossaryTerm
	if err := c.get("/glossary/"+slug+"/", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Compare compares two QR code types.
func (c *Client) Compare(slugA, slugB string) (*CompareResult, error) {
	var result CompareResult
	path := fmt.Sprintf("/compare/?a=%s&b=%s", url.QueryEscape(slugA), url.QueryEscape(slugB))
	if err := c.get(path, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Random returns a random QR code type.
func (c *Client) Random() (*QRTypeDetail, error) {
	var result QRTypeDetail
	if err := c.get("/random/", &result); err != nil {
		return nil, err
	}
	return &result, nil
}
