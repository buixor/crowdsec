package apiclient

import (
	"fmt"
	"net/http"
	"net/url"
)

type ApiClient struct {
	/*The http client used to make requests*/
	client *http.Client
	/*Reuse a single struct instead of allocating one for each service on the heap.*/
	common service
	/*config stuff*/
	BaseURL   *url.URL
	UserAgent string
	/*exposed Services*/
	Decisions *DecisionsService
	Alerts    *AlertsService
	// Auth      *ApiAuth
	// Consensus *ApiConsensus
}

type service struct {
	client *ApiClient
}

func NewClient(httpClient *http.Client) *ApiClient {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	baseURL, _ := url.Parse("http://127.0.0.1:4242/")

	c := &ApiClient{client: httpClient, BaseURL: baseURL}
	c.common.client = c
	c.Decisions = (*DecisionsService)(&c.common)
	return c
}

type Response struct {
	*http.Response
	//add our pagination stuff
	//NextPage int
	//...
}

func newResponse(r *http.Response) *Response {
	response := &Response{Response: r}
	//response.populatePageValues()
	return response
}

func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}
	return fmt.Errorf("error from api : %+v", r)
}

type ListOpts struct {
	//Page    int
	//PerPage int
}

type DeleteOpts struct {
	//??
}

type AddOpts struct {
	//??
}
