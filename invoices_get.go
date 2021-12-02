package netsuite

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-netsuite/utils"
)

func (c *Client) NewInvoicesGetRequest() InvoicesGetRequest {
	r := InvoicesGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type InvoicesGetRequest struct {
	client      *Client
	queryParams *InvoicesGetRequestQueryParams
	pathParams  *InvoicesGetRequestPathParams
	method      string
	headers     http.Header
	requestBody InvoicesGetRequestBody
}

func (r InvoicesGetRequest) NewQueryParams() *InvoicesGetRequestQueryParams {
	return &InvoicesGetRequestQueryParams{}
}

type InvoicesGetRequestQueryParams struct {
	Q      string `schema:"q,omitempty"`
	Limit  int    `schema:"limit,omitempty"`
	Offset int    `schema:"offset,omitempty"`
}

func (p InvoicesGetRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(DateTime{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *InvoicesGetRequest) QueryParams() *InvoicesGetRequestQueryParams {
	return r.queryParams
}

func (r InvoicesGetRequest) NewPathParams() *InvoicesGetRequestPathParams {
	return &InvoicesGetRequestPathParams{}
}

type InvoicesGetRequestPathParams struct {
}

func (p *InvoicesGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *InvoicesGetRequest) PathParams() *InvoicesGetRequestPathParams {
	return r.pathParams
}

func (r *InvoicesGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *InvoicesGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *InvoicesGetRequest) Method() string {
	return r.method
}

func (r InvoicesGetRequest) NewRequestBody() InvoicesGetRequestBody {
	return InvoicesGetRequestBody{}
}

type InvoicesGetRequestBody struct {
}

func (r *InvoicesGetRequest) RequestBody() *InvoicesGetRequestBody {
	return nil
}

func (r *InvoicesGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *InvoicesGetRequest) SetRequestBody(body InvoicesGetRequestBody) {
	r.requestBody = body
}

func (r *InvoicesGetRequest) NewResponseBody() *InvoicesGetResponseBody {
	return &InvoicesGetResponseBody{}
}

type InvoicesGetResponseBody struct {
	Links []struct {
		Rel  string `json:"rel"`
		Href string `json:"href"`
	} `json:"links"`
	Count   int  `json:"count"`
	HasMore bool `json:"hasMore"`
	Items   []struct {
		Links []struct {
			Rel  string `json:"rel"`
			Href string `json:"href"`
		} `json:"links"`
		ID string `json:"id"`
	} `json:"items"`
	Offset       int `json:"offset"`
	TotalResults int `json:"totalResults"`
}

func (r *InvoicesGetRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/record/v1/invoice", r.PathParams())
	return &u, err
}

func (r *InvoicesGetRequest) Do() (InvoicesGetResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}
