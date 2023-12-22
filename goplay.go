package goplay

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"time"
)

func NewClient(option Options, opts ...Option) (*Client, error) {
	for _, opt := range opts {
		opt(&option)
	}

	if option.Address == "" {
		option.Address = DefaultPlayground
	}

	if option.Timeout == 0 {
		option.Timeout = time.Second * 10
	}

	client := resty.New()
	client.BaseURL = option.Address
	client.SetTimeout(option.Timeout)

	return &Client{
		cnn: client,
		opt: option,
	}, nil
}

// Client is the go playground http client
// see https://github.com/golang/playground to learn more about playground server.
type Client struct {
	cnn *resty.Client
	opt Options
}

type Version struct {
	Version, Release, Name string
}

// Version returns the version of playground server
func (client *Client) Version() (Version, error) {
	resp, err := client.cnn.R().Get(VersionUrl)
	if err != nil {
		return Version{}, err
	}
	var v Version
	if err := json.Unmarshal(resp.Body(), &v); err != nil {
		return v, err
	}
	return v, nil
}

// HealCheck checks whether playground server is available
func (client *Client) HealCheck() (bool, error) {
	resp, err := client.cnn.R().Get(HealthUrl)
	if err != nil {
		return false, err
	}
	respStr := b2str(resp.Body())
	if len(respStr) == 0 {
		return false, errors.New("health check failed")
	} else if respStr != "ok" {
		return false, errors.New(respStr)
	}

	return true, nil
}

// View returns the playground code snippet specified by the given id
func (client *Client) View(id string) ([]byte, error) {
	url := fmt.Sprintf(ViewUrl, id)
	resp, err := client.cnn.R().Get(url)
	if err != nil {
		return nil, err
	}
	return resp.Body(), nil
}

// Share shares the code snippet to go playground and returns snippet id.
func (client *Client) Share(raw []byte) (string, error) {
	req := client.cnn.R()
	req.SetBody(raw)
	req.SetHeader("Content-Type", "text/plain; charset=utf-8")
	resp, err := req.Post(ShareUrl)
	if err != nil {
		return "", nil
	}
	return string(resp.Body()), nil
}

type FmtResult struct {
	Body  string
	Error string
}

// FmtRaw fmt the give code snippet and return FmtResult
func (client *Client) FmtRaw(raw []byte, fixImport bool) ([]byte, error) {
	req := client.cnn.R()
	req.SetFormData(map[string]string{
		"body":    b2str(raw),
		"imports": fmt.Sprintf("%+v", fixImport),
	})

	resp, err := req.Post(FmtUrl)
	if err != nil {
		return nil, err
	}

	var res FmtResult
	if err := json.Unmarshal(resp.Body(), &res); err != nil {
		return nil, err
	}

	if res.Error != "" {
		return nil, errors.New(res.Error)
	}

	return []byte(res.Body), nil
}

// Fmt fmt the specified code snippet and return FmtResult
func (client *Client) Fmt(id string, fixImport bool) ([]byte, error) {
	bytes, err := client.View(id)
	if err != nil {
		return nil, err
	}
	return client.FmtRaw(bytes, fixImport)
}

type Event struct {
	Message string
	Kind    string        // "stdout" or "stderr"
	Delay   time.Duration // time to wait before printing Message
}

type CompileResult struct {
	Errors      string
	Events      []Event
	Status      int
	IsTest      bool
	TestsFailed int

	// VetErrors, if non-empty, contains any vet errors. It is
	// only populated if request.WithVet was true.
	VetErrors string `json:",omitempty"`
	// VetOK reports whether vet ran & passed. It is only
	// populated if request.WithVet was true. Only one of
	// VetErrors or VetOK can be non-zero.
	VetOK bool `json:",omitempty"`
}

// CompileRaw compiles the give code snippet, and returns the result
func (client *Client) CompileRaw(raw []byte, vet bool) (CompileResult, error) {
	req := client.cnn.R()
	req.SetFormData(map[string]string{
		"body":    b2str(raw),
		"withVet": fmt.Sprintf("%v", vet),
	})

	var result CompileResult

	resp, err := req.Post(CompileUrl)
	if err != nil {
		return result, err
	}

	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return result, err
	}
	return result, nil
}

// CompileSnippet compiles the specified code snippet, and returns the result
func (client *Client) CompileSnippet(id string, vet bool) (CompileResult, error) {
	bytes, err := client.View(id)
	if err != nil {
		return CompileResult{}, err
	}

	return client.CompileRaw(bytes, vet)
}
