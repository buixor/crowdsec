package apiclient

import (
	"context"
	"fmt"

	"github.com/crowdsecurity/crowdsec/pkg/models"
	qs "github.com/google/go-querystring/query"
)

type DecisionsService service

type DecisionsListOpts struct {
	Scope_equals *string `url:"scope,omitempty"`
	Value_equals *string `url:"value,omitempty"`
	Type_equals  *string `url:"type,omitempty"`
	ListOpts
}

type DecisionsDeleteOpts struct {
	Scope_equals *string `url:"scope,omitempty"`
	Value_equals *string `url:"value,omitempty"`
	Type_equals  *string `url:"type,omitempty"`
	ListOpts
}

//to demo query arguments
func (s *DecisionsService) List(ctx context.Context, opts DecisionsListOpts) (*models.GetDecisionsResponse, *Response, error) {
	var decisions models.GetDecisionsResponse
	params, err := qs.Values(opts)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("decisions/?%s", params.Encode())

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err := s.client.Do(ctx, req, &decisions)
	if err != nil {
		return nil, resp, err
	}
	return &decisions, resp, nil
}

func (s *DecisionsService) StartStream(ctx context.Context) (*Response, error) {
	u := "decisions/stream?startup=true"
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (s *DecisionsService) StopStream(ctx context.Context) (*Response, error) {
	return nil, nil

}

// func (s *DecisionsService) StreamPoll(ctx context.Context) ([]Decision, []Decision, *Response, error) {
// 	return nil, nil

// }

// func (s *DecisionsService) List(ctx context.Context, Opts DecisionsListOpts) ([]Decision, *Response, error) {
// 	return nil, nil

// }

// func (s *DecisionsService) Delete(ctx context.Context, Opts DecisionsDeleteOpts) ([]Decision, *Response, error) {
// 	return nil, nil

// }
