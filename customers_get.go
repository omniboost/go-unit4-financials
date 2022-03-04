package netsuite

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-netsuite-rest/utils"
)

func (c *Client) NewCustomersGetRequest() CustomersGetRequest {
	r := CustomersGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type CustomersGetRequest struct {
	client      *Client
	queryParams *CustomersGetRequestQueryParams
	pathParams  *CustomersGetRequestPathParams
	method      string
	headers     http.Header
	requestBody CustomersGetRequestBody
}

func (r CustomersGetRequest) NewQueryParams() *CustomersGetRequestQueryParams {
	return &CustomersGetRequestQueryParams{}
}

type CustomersGetRequestQueryParams struct {
	Q      string `schema:"q,omitempty"`
	Limit  int    `schema:"limit,omitempty"`
	Offset int    `schema:"offset,omitempty"`
}

func (p CustomersGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *CustomersGetRequest) QueryParams() *CustomersGetRequestQueryParams {
	return r.queryParams
}

func (r CustomersGetRequest) NewPathParams() *CustomersGetRequestPathParams {
	return &CustomersGetRequestPathParams{}
}

type CustomersGetRequestPathParams struct {
}

func (p *CustomersGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *CustomersGetRequest) PathParams() *CustomersGetRequestPathParams {
	return r.pathParams
}

func (r *CustomersGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *CustomersGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *CustomersGetRequest) Method() string {
	return r.method
}

func (r CustomersGetRequest) NewRequestBody() CustomersGetRequestBody {
	return CustomersGetRequestBody{}
}

type CustomersGetRequestBody struct {
}

func (r *CustomersGetRequest) RequestBody() *CustomersGetRequestBody {
	return nil
}

func (r *CustomersGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *CustomersGetRequest) SetRequestBody(body CustomersGetRequestBody) {
	r.requestBody = body
}

func (r *CustomersGetRequest) NewResponseBody() *CustomersGetResponseBody {
	return &CustomersGetResponseBody{}
}

type CustomersGetResponseBody struct {
	Links   Links `json:"links"`
	Count   int   `json:"count"`
	HasMore bool  `json:"hasMore"`
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

func (r *CustomersGetRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/record/v1/customer", r.PathParams())
	return &u, err
}

func (r *CustomersGetRequest) Do() (CustomersGetResponseBody, error) {
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
