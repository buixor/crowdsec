package apiclient

import (
	"context"
	"fmt"

	qs "github.com/google/go-querystring/query"
)

type DecisionsService service

type DecisionsListOpts struct {
	scope_equals *string `"json:scope"`
	value_equals *string `"json:value"`
	type_equals  *string `"json:type"`
	ListOpts
}

type DecisionsDeleteOpts struct {
	scope_equals *string `"json:scope"`
	value_equals *string `"json:value"`
	type_equals  *string `"json:type"`
	ListOpts
}

/*should be define by swagger or smth ?*/
type Decision struct {
	ID            *int64
	decisionType  *string
	decisionScope *string
	decisionValue *string
}

type Decisions struct {
	TotalCount int         `"json:count"`
	Decisions  []*Decision `"json:decisions"`
}

//to demo query arguments
func (s *DecisionsService) List(ctx context.Context, opts DecisionsListOpts) (*Decisions, *Response, error) {
	var decisions Decisions
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
